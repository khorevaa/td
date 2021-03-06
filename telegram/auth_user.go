package telegram

import (
	"context"
	"errors"

	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/crypto/srp"
	"github.com/gotd/td/tg"
)

// AuthPassword performs login via secure remote password (aka 2FA).
//
// Method can be called after AuthSignIn to provide password if requested.
func (c *Client) AuthPassword(ctx context.Context, password string) error {
	p, err := c.tg.AccountGetPassword(ctx)
	if err != nil {
		return xerrors.Errorf("get SRP parameters: %w", err)
	}

	algo, ok := p.CurrentAlgo.(*tg.PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow)
	if !ok {
		return xerrors.Errorf("unsupported algo: %T", p.CurrentAlgo)
	}

	s := srp.NewSRP(c.rand)
	a, err := s.Hash([]byte(password), p.SrpB, p.SecureRandom, srp.Input(*algo))
	if err != nil {
		return xerrors.Errorf("create SRP answer: %w", err)
	}

	auth, err := c.tg.AuthCheckPassword(ctx, &tg.InputCheckPasswordSRP{
		SrpID: p.SrpID,
		A:     a.A,
		M1:    a.M1,
	})
	if err != nil {
		return xerrors.Errorf("check password: %w", err)
	}
	if err := c.checkAuthResult(auth); err != nil {
		return xerrors.Errorf("check: %w", err)
	}

	return nil
}

// SendCodeOptions defines how to send auth code to user.
type SendCodeOptions struct {
	// AllowFlashCall allows phone verification via phone calls.
	AllowFlashCall bool
	// Pass true if the phone number is used on the current device.
	// Ignored if AllowFlashCall is not set.
	CurrentNumber bool
	// If a token that will be included in eventually sent SMSs is required:
	// required in newer versions of android, to use the android SMS receiver APIs.
	AllowAppHash bool
}

// AuthSendCode requests code for provided phone number, returning code hash
// and error if any. Use tgflow.Auth to reduce boilerplate.
//
// This method should be called first in user authentication flow.
func (c *Client) AuthSendCode(ctx context.Context, phone string, options SendCodeOptions) (codeHash string, err error) {
	var settings tg.CodeSettings
	if options.AllowAppHash {
		settings.SetAllowAppHash(true)
	}
	if options.AllowFlashCall {
		settings.SetAllowFlashcall(true)
	}
	if options.CurrentNumber {
		settings.SetCurrentNumber(true)
	}

	sentCode, err := c.tg.AuthSendCode(ctx, &tg.AuthSendCodeRequest{
		PhoneNumber: phone,
		APIID:       c.appID,
		APIHash:     c.appHash,
		Settings:    settings,
	})
	if err != nil {
		return "", xerrors.Errorf("send code: %w", err)
	}
	return sentCode.PhoneCodeHash, nil
}

// ErrPasswordAuthNeeded means that 2FA auth is required.
//
// Call Client.AuthPassword to provide 2FA password.
var ErrPasswordAuthNeeded = errors.New("2FA required")

// AuthSignIn performs sign in with provided user phone, code and code hash.
//
// If ErrPasswordAuthNeeded is returned, call AuthPassword to provide 2FA
// password.
//
// To obtain codeHash, use AuthSendCode.
func (c *Client) AuthSignIn(ctx context.Context, phone, code, codeHash string) error {
	a, err := c.tg.AuthSignIn(ctx, &tg.AuthSignInRequest{
		PhoneNumber:   phone,
		PhoneCodeHash: codeHash,
		PhoneCode:     code,
	})
	var rpcErr *Error
	if errors.As(err, &rpcErr) && rpcErr.Message == "SESSION_PASSWORD_NEEDED" {
		return ErrPasswordAuthNeeded
	}
	if err != nil {
		return xerrors.Errorf("sign in: %w", err)
	}
	if err := c.checkAuthResult(a); err != nil {
		return xerrors.Errorf("check: %w", err)
	}

	return nil
}
