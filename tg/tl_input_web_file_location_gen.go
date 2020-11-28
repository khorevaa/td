// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// InputWebFileLocation represents TL type `inputWebFileLocation#c239d686`.
type InputWebFileLocation struct {
	// URL field of InputWebFileLocation.
	URL string
	// AccessHash field of InputWebFileLocation.
	AccessHash int64
}

// InputWebFileLocationTypeID is TL type id of InputWebFileLocation.
const InputWebFileLocationTypeID = 0xc239d686

// Encode implements bin.Encoder.
func (i *InputWebFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebFileLocation#c239d686 as nil")
	}
	b.PutID(InputWebFileLocationTypeID)
	b.PutString(i.URL)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWebFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebFileLocation#c239d686 to nil")
	}
	if err := b.ConsumeID(InputWebFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWebFileLocation#c239d686: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileLocation#c239d686: field url: %w", err)
		}
		i.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileLocation#c239d686: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputWebFileLocationClass.
func (i InputWebFileLocation) construct() InputWebFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputWebFileLocation.
var (
	_ bin.Encoder = &InputWebFileLocation{}
	_ bin.Decoder = &InputWebFileLocation{}

	_ InputWebFileLocationClass = &InputWebFileLocation{}
)

// InputWebFileGeoPointLocation represents TL type `inputWebFileGeoPointLocation#9f2221c9`.
type InputWebFileGeoPointLocation struct {
	// GeoPoint field of InputWebFileGeoPointLocation.
	GeoPoint InputGeoPointClass
	// AccessHash field of InputWebFileGeoPointLocation.
	AccessHash int64
	// W field of InputWebFileGeoPointLocation.
	W int
	// H field of InputWebFileGeoPointLocation.
	H int
	// Zoom field of InputWebFileGeoPointLocation.
	Zoom int
	// Scale field of InputWebFileGeoPointLocation.
	Scale int
}

// InputWebFileGeoPointLocationTypeID is TL type id of InputWebFileGeoPointLocation.
const InputWebFileGeoPointLocationTypeID = 0x9f2221c9

// Encode implements bin.Encoder.
func (i *InputWebFileGeoPointLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebFileGeoPointLocation#9f2221c9 as nil")
	}
	b.PutID(InputWebFileGeoPointLocationTypeID)
	if i.GeoPoint == nil {
		return fmt.Errorf("unable to encode inputWebFileGeoPointLocation#9f2221c9: field geo_point is nil")
	}
	if err := i.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputWebFileGeoPointLocation#9f2221c9: field geo_point: %w", err)
	}
	b.PutLong(i.AccessHash)
	b.PutInt(i.W)
	b.PutInt(i.H)
	b.PutInt(i.Zoom)
	b.PutInt(i.Scale)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWebFileGeoPointLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebFileGeoPointLocation#9f2221c9 to nil")
	}
	if err := b.ConsumeID(InputWebFileGeoPointLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: %w", err)
	}
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field geo_point: %w", err)
		}
		i.GeoPoint = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field w: %w", err)
		}
		i.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field h: %w", err)
		}
		i.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field zoom: %w", err)
		}
		i.Zoom = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field scale: %w", err)
		}
		i.Scale = value
	}
	return nil
}

// construct implements constructor of InputWebFileLocationClass.
func (i InputWebFileGeoPointLocation) construct() InputWebFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputWebFileGeoPointLocation.
var (
	_ bin.Encoder = &InputWebFileGeoPointLocation{}
	_ bin.Decoder = &InputWebFileGeoPointLocation{}

	_ InputWebFileLocationClass = &InputWebFileGeoPointLocation{}
)

// InputWebFileLocationClass represents InputWebFileLocation generic type.
//
// Example:
//  g, err := DecodeInputWebFileLocation(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputWebFileLocation: // inputWebFileLocation#c239d686
//  case *InputWebFileGeoPointLocation: // inputWebFileGeoPointLocation#9f2221c9
//  default: panic(v)
//  }
type InputWebFileLocationClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputWebFileLocationClass
}

// DecodeInputWebFileLocation implements binary de-serialization for InputWebFileLocationClass.
func DecodeInputWebFileLocation(buf *bin.Buffer) (InputWebFileLocationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputWebFileLocationTypeID:
		// Decoding inputWebFileLocation#c239d686.
		v := InputWebFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWebFileLocationClass: %w", err)
		}
		return &v, nil
	case InputWebFileGeoPointLocationTypeID:
		// Decoding inputWebFileGeoPointLocation#9f2221c9.
		v := InputWebFileGeoPointLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWebFileLocationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputWebFileLocationClass: %w", bin.NewUnexpectedID(id))
	}
}