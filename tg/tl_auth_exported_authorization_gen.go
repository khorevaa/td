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

// AuthExportedAuthorization represents TL type `auth.exportedAuthorization#df969c2d`.
// Data for copying of authorization between data centres.
//
// See https://core.telegram.org/constructor/auth.exportedAuthorization for reference.
type AuthExportedAuthorization struct {
	// current user identifier
	ID int
	// authorizes key
	Bytes []byte
}

// AuthExportedAuthorizationTypeID is TL type id of AuthExportedAuthorization.
const AuthExportedAuthorizationTypeID = 0xdf969c2d

// String implements fmt.Stringer.
func (e *AuthExportedAuthorization) String() string {
	if e == nil {
		return "AuthExportedAuthorization(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthExportedAuthorization")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tBytes: ")
	sb.WriteString(fmt.Sprint(e.Bytes))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *AuthExportedAuthorization) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportedAuthorization#df969c2d as nil")
	}
	b.PutID(AuthExportedAuthorizationTypeID)
	b.PutInt(e.ID)
	b.PutBytes(e.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (e *AuthExportedAuthorization) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportedAuthorization#df969c2d to nil")
	}
	if err := b.ConsumeID(AuthExportedAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.exportedAuthorization#df969c2d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportedAuthorization#df969c2d: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportedAuthorization#df969c2d: field bytes: %w", err)
		}
		e.Bytes = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthExportedAuthorization.
var (
	_ bin.Encoder = &AuthExportedAuthorization{}
	_ bin.Decoder = &AuthExportedAuthorization{}
)
