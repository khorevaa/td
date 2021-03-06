package crypto

import (
	"crypto/sha1" // #nosec
	"crypto/sha256"
	"hash"
	"sync"

	"github.com/gotd/td/bin"
)

// See https://core.telegram.org/mtproto/description#defining-aes-key-and-initialization-vector

// AuthKey represents 2048-bit authorization key.
type AuthKey [256]byte

// Zero reports whether AuthKey is zero value.
func (k AuthKey) Zero() bool {
	return k == AuthKey{}
}

// ID returns auth_key_id.
func (k AuthKey) ID() [8]byte {
	raw := sha1.Sum(k[:]) // #nosec
	var id [8]byte
	copy(id[:], raw[12:])
	return id
}

// AuxHash returns aux_hash value of key.
func (k AuthKey) AuxHash() [8]byte {
	raw := sha1.Sum(k[:]) // #nosec
	var id [8]byte
	copy(id[:], raw[0:8])
	return id
}

// WithID creates new AuthKeyWithID from AuthKey.
func (k AuthKey) WithID() AuthKeyWithID {
	return AuthKeyWithID{
		AuthKey:   k,
		AuthKeyID: k.ID(),
	}
}

// Side on which encryption is performed.
type Side byte

const (
	// Client side of encryption (e.g. messages from client).
	Client Side = 0
	// Server side of encryption (e.g. RPC responses).
	Server Side = 1
)

// DecryptSide returns Side for decryption.
func (s Side) DecryptSide() Side {
	return s ^ 1 // flips bit, so 0 becomes 1, 1 becomes 0
}

func getX(mode Side) int {
	switch mode {
	case Client:
		return 0
	case Server:
		return 8
	default:
		return 0
	}
}

// nolint:gochecknoglobals // optimization for sha256-hash reuse
var sha256Pool = &sync.Pool{
	New: func() interface{} {
		return sha256.New()
	},
}

func getSHA256() hash.Hash {
	h := sha256Pool.Get().(hash.Hash)
	h.Reset()
	return h
}

// Message keys are defined here:
// * https://core.telegram.org/mtproto/description#defining-aes-key-and-initialization-vector

// msgKeyLarge returns msg_key_large value.
func msgKeyLarge(authKey AuthKey, plaintextPadded []byte, mode Side) []byte {
	h := getSHA256()
	defer sha256Pool.Put(h)

	x := getX(mode)
	_, _ = h.Write(authKey[88+x : 32+88+x])
	_, _ = h.Write(plaintextPadded)
	return h.Sum(nil)
}

// messageKey returns msg_key = substr (msg_key_large, 8, 16).
func messageKey(messageKeyLarge []byte) bin.Int128 {
	var v bin.Int128
	b := messageKeyLarge[8 : 16+8]
	copy(v[:len(b)], b)
	return v
}

// SHA256 returns SHA256 hash.
func SHA256(from []byte) []byte {
	h := getSHA256()
	defer sha256Pool.Put(h)
	_, _ = h.Write(from)
	return h.Sum(nil)
}

// sha256a returns sha256_a value.
//
// sha256_a = SHA256 (msg_key + substr (auth_key, x, 36));
func sha256a(authKey AuthKey, msgKey bin.Int128, x int) []byte {
	h := getSHA256()
	defer sha256Pool.Put(h)

	_, _ = h.Write(msgKey[:])
	_, _ = h.Write(authKey[x : x+36])

	return h.Sum(nil)
}

// sha256b returns sha256_b value.
//
// sha256_b = SHA256 (substr (auth_key, 40+x, 36) + msg_key);
func sha256b(authKey AuthKey, msgKey bin.Int128, x int) []byte {
	h := getSHA256()
	defer sha256Pool.Put(h)

	_, _ = h.Write(authKey[40+x : 40+x+36])
	_, _ = h.Write(msgKey[:])

	return h.Sum(nil)
}

// aesKey returns aes_key value.
//
// aes_key = substr (sha256_a, 0, 8) + substr (sha256_b, 8, 16) + substr (sha256_a, 24, 8);
func aesKey(sha256a, sha256b []byte) bin.Int256 {
	var v bin.Int256
	copy(v[:8], sha256a[:8])
	copy(v[8:], sha256b[8:16+8])
	copy(v[24:], sha256a[24:24+8])
	return v
}

// aesIV returns aes_iv value.
//
// aes_iv = substr (sha256_b, 0, 8) + substr (sha256_a, 8, 16) + substr (sha256_b, 24, 8);
func aesIV(sha256a, sha256b []byte) bin.Int256 {
	// Same as aes_key, but with swapped params.
	return aesKey(sha256b, sha256a)
}

// Keys returns (aes_key, aes_iv) pair for AES-IGE.
//
// Reference:
// * https://core.telegram.org/mtproto/description#defining-aes-key-and-initialization-vector
//
// Example:
//	key, iv := crypto.Keys(authKey, messageKey, crypto.Client)
//	cipher, err := aes.NewCipher(key[:])
//	if err != nil {
//		return nil, err
//	}
//	encryptor := ige.NewIGEEncrypter(cipher, iv[:])
func Keys(authKey AuthKey, msgKey bin.Int128, mode Side) (key, iv bin.Int256) {
	x := getX(mode)

	// `sha256_a = SHA256 (msg_key + substr (auth_key, x, 36));`
	a := sha256a(authKey, msgKey, x)
	// `sha256_b = SHA256 (substr (auth_key, 40+x, 36) + msg_key);`
	b := sha256b(authKey, msgKey, x)

	return aesKey(a, b), aesIV(a, b)
}

// MessageKey computes message key for provided auth_key and padded payload.
func MessageKey(authKey AuthKey, plaintextPadded []byte, mode Side) bin.Int128 {
	// `msg_key_large = SHA256 (substr (auth_key, 88+x, 32) + plaintext + random_padding);`
	msgKeyLarge := msgKeyLarge(authKey, plaintextPadded, mode)
	// `msg_key = substr (msg_key_large, 8, 16);`
	return messageKey(msgKeyLarge)
}
