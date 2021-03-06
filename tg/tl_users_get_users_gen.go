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

// UsersGetUsersRequest represents TL type `users.getUsers#d91a548`.
// Returns basic user info according to their identifiers.
//
// See https://core.telegram.org/method/users.getUsers for reference.
type UsersGetUsersRequest struct {
	// List of user identifiers
	ID []InputUserClass
}

// UsersGetUsersRequestTypeID is TL type id of UsersGetUsersRequest.
const UsersGetUsersRequestTypeID = 0xd91a548

// String implements fmt.Stringer.
func (g *UsersGetUsersRequest) String() string {
	if g == nil {
		return "UsersGetUsersRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UsersGetUsersRequest")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range g.ID {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *UsersGetUsersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode users.getUsers#d91a548 as nil")
	}
	b.PutID(UsersGetUsersRequestTypeID)
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode users.getUsers#d91a548: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode users.getUsers#d91a548: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *UsersGetUsersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode users.getUsers#d91a548 to nil")
	}
	if err := b.ConsumeID(UsersGetUsersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode users.getUsers#d91a548: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode users.getUsers#d91a548: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode users.getUsers#d91a548: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for UsersGetUsersRequest.
var (
	_ bin.Encoder = &UsersGetUsersRequest{}
	_ bin.Decoder = &UsersGetUsersRequest{}
)

// UsersGetUsers invokes method users.getUsers#d91a548 returning error if any.
// Returns basic user info according to their identifiers.
//
// See https://core.telegram.org/method/users.getUsers for reference.
func (c *Client) UsersGetUsers(ctx context.Context, id []InputUserClass) ([]UserClass, error) {
	var result UserClassVector

	request := &UsersGetUsersRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
