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

// DataJSON represents TL type `dataJSON#7d748d04`.
// Represents a json-encoded object
//
// See https://core.telegram.org/constructor/dataJSON for reference.
type DataJSON struct {
	// JSON-encoded object
	Data string
}

// DataJSONTypeID is TL type id of DataJSON.
const DataJSONTypeID = 0x7d748d04

// String implements fmt.Stringer.
func (d *DataJSON) String() string {
	if d == nil {
		return "DataJSON(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DataJSON")
	sb.WriteString("{\n")
	sb.WriteString("\tData: ")
	sb.WriteString(fmt.Sprint(d.Data))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DataJSON) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dataJSON#7d748d04 as nil")
	}
	b.PutID(DataJSONTypeID)
	b.PutString(d.Data)
	return nil
}

// Decode implements bin.Decoder.
func (d *DataJSON) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dataJSON#7d748d04 to nil")
	}
	if err := b.ConsumeID(DataJSONTypeID); err != nil {
		return fmt.Errorf("unable to decode dataJSON#7d748d04: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dataJSON#7d748d04: field data: %w", err)
		}
		d.Data = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DataJSON.
var (
	_ bin.Encoder = &DataJSON{}
	_ bin.Decoder = &DataJSON{}
)
