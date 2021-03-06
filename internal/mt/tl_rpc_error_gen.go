// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// RPCError represents TL type `rpc_error#2144ca19`.
type RPCError struct {
	// ErrorCode field of RPCError.
	ErrorCode int
	// ErrorMessage field of RPCError.
	ErrorMessage string
}

// RPCErrorTypeID is TL type id of RPCError.
const RPCErrorTypeID = 0x2144ca19

// String implements fmt.Stringer.
func (r *RPCError) String() string {
	if r == nil {
		return "RPCError(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RPCError")
	sb.WriteString("{\n")
	sb.WriteString("\tErrorCode: ")
	sb.WriteString(fmt.Sprint(r.ErrorCode))
	sb.WriteString(",\n")
	sb.WriteString("\tErrorMessage: ")
	sb.WriteString(fmt.Sprint(r.ErrorMessage))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *RPCError) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_error#2144ca19 as nil")
	}
	b.PutID(RPCErrorTypeID)
	b.PutInt(r.ErrorCode)
	b.PutString(r.ErrorMessage)
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCError) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_error#2144ca19 to nil")
	}
	if err := b.ConsumeID(RPCErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_error#2144ca19: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_error#2144ca19: field error_code: %w", err)
		}
		r.ErrorCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_error#2144ca19: field error_message: %w", err)
		}
		r.ErrorMessage = value
	}
	return nil
}

// Ensuring interfaces in compile-time for RPCError.
var (
	_ bin.Encoder = &RPCError{}
	_ bin.Decoder = &RPCError{}
)
