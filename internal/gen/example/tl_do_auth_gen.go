// Code generated by gotdgen, DO NOT EDIT.

package td

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

// DoAuthRequest represents TL type `doAuth#fd2f6687`.
//
// See https://localhost:80/doc/method/doAuth for reference.
type DoAuthRequest struct {
}

// DoAuthRequestTypeID is TL type id of DoAuthRequest.
const DoAuthRequestTypeID = 0xfd2f6687

// String implements fmt.Stringer.
func (d *DoAuthRequest) String() string {
	if d == nil {
		return "DoAuthRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DoAuthRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DoAuthRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode doAuth#fd2f6687 as nil")
	}
	b.PutID(DoAuthRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DoAuthRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode doAuth#fd2f6687 to nil")
	}
	if err := b.ConsumeID(DoAuthRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode doAuth#fd2f6687: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for DoAuthRequest.
var (
	_ bin.Encoder = &DoAuthRequest{}
	_ bin.Decoder = &DoAuthRequest{}
)

// DoAuth invokes method doAuth#fd2f6687 returning error if any.
//
// See https://localhost:80/doc/method/doAuth for reference.
func (c *Client) DoAuth(ctx context.Context) (AuthClass, error) {
	var result AuthBox

	request := &DoAuthRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Auth, nil
}
