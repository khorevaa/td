// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// AuthBindTempAuthKeyRequest represents TL type `auth.bindTempAuthKey#cdd42a05`.
// Binds a temporary authorization key temp_auth_key_id to the permanent authorization key perm_auth_key_id. Each permanent key may only be bound to one temporary key at a time, binding a new temporary key overwrites the previous one.
// For more information, see Perfect Forward Secrecy¹.
//
// Links:
//  1) https://core.telegram.org/api/pfs
//
// See https://core.telegram.org/method/auth.bindTempAuthKey for reference.
type AuthBindTempAuthKeyRequest struct {
	// Permanent auth_key_id to bind to
	PermAuthKeyID int64
	// Random long from Binding message contents¹
	//
	// Links:
	//  1) https://core.telegram.org#binding-message-contents
	Nonce int64
	// Unix timestamp to invalidate temporary key, see Binding message contents¹
	//
	// Links:
	//  1) https://core.telegram.org#binding-message-contents
	ExpiresAt int
	// See Generating encrypted_message¹
	//
	// Links:
	//  1) https://core.telegram.org#generating-encrypted-message
	EncryptedMessage []byte
}

// AuthBindTempAuthKeyRequestTypeID is TL type id of AuthBindTempAuthKeyRequest.
const AuthBindTempAuthKeyRequestTypeID = 0xcdd42a05

// String implements fmt.Stringer.
func (b *AuthBindTempAuthKeyRequest) String() string {
	if b == nil {
		return "AuthBindTempAuthKeyRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthBindTempAuthKeyRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPermAuthKeyID: ")
	sb.WriteString(fmt.Sprint(b.PermAuthKeyID))
	sb.WriteString(",\n")
	sb.WriteString("\tNonce: ")
	sb.WriteString(fmt.Sprint(b.Nonce))
	sb.WriteString(",\n")
	sb.WriteString("\tExpiresAt: ")
	sb.WriteString(fmt.Sprint(b.ExpiresAt))
	sb.WriteString(",\n")
	sb.WriteString("\tEncryptedMessage: ")
	sb.WriteString(fmt.Sprint(b.EncryptedMessage))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *AuthBindTempAuthKeyRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode auth.bindTempAuthKey#cdd42a05 as nil")
	}
	buf.PutID(AuthBindTempAuthKeyRequestTypeID)
	buf.PutLong(b.PermAuthKeyID)
	buf.PutLong(b.Nonce)
	buf.PutInt(b.ExpiresAt)
	buf.PutBytes(b.EncryptedMessage)
	return nil
}

// Decode implements bin.Decoder.
func (b *AuthBindTempAuthKeyRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode auth.bindTempAuthKey#cdd42a05 to nil")
	}
	if err := buf.ConsumeID(AuthBindTempAuthKeyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: %w", err)
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field perm_auth_key_id: %w", err)
		}
		b.PermAuthKeyID = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field nonce: %w", err)
		}
		b.Nonce = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field expires_at: %w", err)
		}
		b.ExpiresAt = value
	}
	{
		value, err := buf.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field encrypted_message: %w", err)
		}
		b.EncryptedMessage = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthBindTempAuthKeyRequest.
var (
	_ bin.Encoder = &AuthBindTempAuthKeyRequest{}
	_ bin.Decoder = &AuthBindTempAuthKeyRequest{}
)

// AuthBindTempAuthKey invokes method auth.bindTempAuthKey#cdd42a05 returning error if any.
// Binds a temporary authorization key temp_auth_key_id to the permanent authorization key perm_auth_key_id. Each permanent key may only be bound to one temporary key at a time, binding a new temporary key overwrites the previous one.
// For more information, see Perfect Forward Secrecy¹.
//
// Links:
//  1) https://core.telegram.org/api/pfs
//
// See https://core.telegram.org/method/auth.bindTempAuthKey for reference.
func (c *Client) AuthBindTempAuthKey(ctx context.Context, request *AuthBindTempAuthKeyRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
