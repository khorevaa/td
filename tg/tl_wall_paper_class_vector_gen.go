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

// WallPaperClassVector is a box for Vector<WallPaper>
type WallPaperClassVector struct {
	// Elements of Vector<WallPaper>
	Elems []WallPaperClass
}

// String implements fmt.Stringer.
func (vec *WallPaperClassVector) String() string {
	if vec == nil {
		return "WallPaperClassVector(nil)"
	}
	var sb strings.Builder
	sb.WriteString("WallPaperClassVector")
	sb.WriteByte('[')
	for _, e := range vec.Elems {
		sb.WriteString(fmt.Sprint(e) + ",\n")
	}
	sb.WriteByte(']')
	return sb.String()
}

// Encode implements bin.Encoder.
func (vec *WallPaperClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<WallPaper> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *WallPaperClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<WallPaper> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for WallPaperClassVector.
var (
	_ bin.Encoder = &WallPaperClassVector{}
	_ bin.Decoder = &WallPaperClassVector{}
)
