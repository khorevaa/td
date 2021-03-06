package telegram

import (
	"context"
	"crypto/rsa"
	"io"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/internal/tmap"
	"github.com/gotd/td/telegram/internal/rpc"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/transport"
)

// UpdateHandler will be called on received updates from Telegram.
type UpdateHandler func(ctx context.Context, u *tg.Updates) error

// Available MTProto default server addresses.
//
// See https://my.telegram.org/apps.
const (
	AddrProduction = "149.154.167.50:443"
	AddrTest       = "149.154.167.40:443"
)

// Test-only credentials. Can be used with AddrTest and tgflow.TestAuth to
// test authentication.
//
// Reference:
//	* https://github.com/telegramdesktop/tdesktop/blob/5f665b8ecb48802cd13cfb48ec834b946459274a/docs/api_credentials.md
const (
	TestAppID   = 17349
	TestAppHash = "344583e45741c457fe1862106095a5eb"
)

// Client represents a MTProto client to Telegram.
type Client struct {
	// tg provides RPC calls via Client.
	tg *tg.Client

	// conn is owned by Client and not exposed.
	transport Transport
	conn      transport.Conn
	connMux   sync.RWMutex
	addr      string

	// Wrappers for external world, like current time, logs or PRNG.
	// Should be immutable.
	clock  func() time.Time
	rand   io.Reader
	cipher crypto.Cipher
	log    *zap.Logger

	sessionCreated *condOnce

	// Access to authKey and authKeyID is not synchronized because
	// serial access ensured in Dial (i.e. no concurrent access possible).
	authKey crypto.AuthKeyWithID

	salt    int64 // atomic access only
	session int64 // atomic access only

	// sentContentMessages is count of created content messages, used to
	// compute sequence number within session.
	//
	// protected by sentContentMessagesMux.
	sentContentMessages    int32
	sentContentMessagesMux sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	appID   int    // immutable
	appHash string // immutable

	updateHandler  UpdateHandler  // immutable
	sessionStorage SessionStorage // immutable

	rpc *rpc.Engine

	// ackSendChan is queue for outgoing message id's that require waiting for
	// ack from server.
	ackSendChan  chan int64
	ackBatchSize int
	ackInterval  time.Duration

	// callbacks for ping results protected by pingMux.
	// Key is ping id.
	ping    map[int64]func()
	pingMux sync.Mutex

	// immutable
	rsaPublicKeys []*rsa.PublicKey

	types *tmap.Map
}

func (c *Client) newUnencryptedMessage(payload bin.Encoder, b *bin.Buffer) error {
	b.Reset()
	if err := payload.Encode(b); err != nil {
		return err
	}
	msg := proto.UnencryptedMessage{
		MessageID:   c.newMessageID(),
		MessageData: b.Copy(),
	}
	b.Reset()
	return msg.Encode(b)
}

// NewClient creates new unstarted client.
func NewClient(appID int, appHash string, opt Options) *Client {
	// Set default values, if user does not set.
	opt.setDefaults()

	clientCtx, clientCancel := context.WithCancel(context.Background())
	client := &Client{
		addr:      opt.Addr,
		transport: opt.Transport,

		clock:  time.Now,
		rand:   opt.Random,
		cipher: crypto.NewClientCipher(opt.Random),
		log:    opt.Logger,
		ping:   map[int64]func(){},

		sessionCreated: createCondOnce(),

		ackSendChan:  make(chan int64),
		ackInterval:  opt.AckInterval,
		ackBatchSize: opt.AckBatchSize,

		ctx:    clientCtx,
		cancel: clientCancel,

		appID:   appID,
		appHash: appHash,

		sessionStorage: opt.SessionStorage,
		rsaPublicKeys:  opt.PublicKeys,
		updateHandler:  opt.UpdateHandler,

		types: tmap.New(
			mt.TypesMap(),
			tg.TypesMap(),
			proto.TypesMap(),
		),
	}

	client.rpc = rpc.New(client.write, rpc.Config{
		Logger:        opt.Logger.Named("rpc"),
		RetryInterval: opt.RetryInterval,
		MaxRetries:    opt.MaxRetries,
	})

	// Initializing internal RPC caller.
	client.tg = tg.NewClient(client)

	return client
}

// Connect initializes connection to Telegram server and starts internal
// read loop.
func (c *Client) Connect(ctx context.Context) (err error) {
	// Loading session from storage if provided.
	if err := c.loadSession(ctx); err != nil {
		// TODO: Add opt-in config to ignore session load failures.
		return xerrors.Errorf("load session: %w", err)
	}

	// Starting connection.
	//
	// This will send initial packet to telegram and perform key exchange
	// if needed.
	if err := c.connect(ctx); err != nil {
		return xerrors.Errorf("start: %w", err)
	}

	// Spawning goroutines.
	go c.readLoop(c.ctx)
	go c.ackLoop(c.ctx)
	go c.pingLoop(c.ctx)

	if err := c.initConnection(ctx); err != nil {
		return xerrors.Errorf("init: %w", err)
	}

	return nil
}

// connect establishes connection in intermediate mode, creating new auth key
// if needed.
func (c *Client) connect(ctx context.Context) error {
	conn, err := c.transport.DialContext(ctx, "tcp", c.addr)
	if err != nil {
		return xerrors.Errorf("dial failed: %w", err)
	}

	c.connMux.Lock()
	defer c.connMux.Unlock()
	c.conn = conn

	if c.authKey.Zero() {
		c.log.Info("Generating new auth key")
		start := c.clock()
		if err := c.createAuthKey(ctx); err != nil {
			return xerrors.Errorf("create auth key: %w", err)
		}

		if err := c.saveSession(ctx); err != nil {
			return xerrors.Errorf("failed to save session: %w", err)
		}

		c.log.With(zap.Duration("duration", c.clock().Sub(start))).Info("Auth key generated")
	}
	return nil
}
