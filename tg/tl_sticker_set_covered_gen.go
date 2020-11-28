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

// StickerSetCovered represents TL type `stickerSetCovered#6410a5d2`.
type StickerSetCovered struct {
	// Set field of StickerSetCovered.
	Set StickerSet
	// Cover field of StickerSetCovered.
	Cover DocumentClass
}

// StickerSetCoveredTypeID is TL type id of StickerSetCovered.
const StickerSetCoveredTypeID = 0x6410a5d2

// Encode implements bin.Encoder.
func (s *StickerSetCovered) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSetCovered#6410a5d2 as nil")
	}
	b.PutID(StickerSetCoveredTypeID)
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetCovered#6410a5d2: field set: %w", err)
	}
	if s.Cover == nil {
		return fmt.Errorf("unable to encode stickerSetCovered#6410a5d2: field cover is nil")
	}
	if err := s.Cover.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetCovered#6410a5d2: field cover: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickerSetCovered) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSetCovered#6410a5d2 to nil")
	}
	if err := b.ConsumeID(StickerSetCoveredTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerSetCovered#6410a5d2: %w", err)
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickerSetCovered#6410a5d2: field set: %w", err)
		}
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetCovered#6410a5d2: field cover: %w", err)
		}
		s.Cover = value
	}
	return nil
}

// construct implements constructor of StickerSetCoveredClass.
func (s StickerSetCovered) construct() StickerSetCoveredClass { return &s }

// Ensuring interfaces in compile-time for StickerSetCovered.
var (
	_ bin.Encoder = &StickerSetCovered{}
	_ bin.Decoder = &StickerSetCovered{}

	_ StickerSetCoveredClass = &StickerSetCovered{}
)

// StickerSetMultiCovered represents TL type `stickerSetMultiCovered#3407e51b`.
type StickerSetMultiCovered struct {
	// Set field of StickerSetMultiCovered.
	Set StickerSet
	// Covers field of StickerSetMultiCovered.
	Covers []DocumentClass
}

// StickerSetMultiCoveredTypeID is TL type id of StickerSetMultiCovered.
const StickerSetMultiCoveredTypeID = 0x3407e51b

// Encode implements bin.Encoder.
func (s *StickerSetMultiCovered) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSetMultiCovered#3407e51b as nil")
	}
	b.PutID(StickerSetMultiCoveredTypeID)
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetMultiCovered#3407e51b: field set: %w", err)
	}
	b.PutVectorHeader(len(s.Covers))
	for idx, v := range s.Covers {
		if v == nil {
			return fmt.Errorf("unable to encode stickerSetMultiCovered#3407e51b: field covers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stickerSetMultiCovered#3407e51b: field covers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickerSetMultiCovered) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSetMultiCovered#3407e51b to nil")
	}
	if err := b.ConsumeID(StickerSetMultiCoveredTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: %w", err)
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: field set: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: field covers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: field covers: %w", err)
			}
			s.Covers = append(s.Covers, value)
		}
	}
	return nil
}

// construct implements constructor of StickerSetCoveredClass.
func (s StickerSetMultiCovered) construct() StickerSetCoveredClass { return &s }

// Ensuring interfaces in compile-time for StickerSetMultiCovered.
var (
	_ bin.Encoder = &StickerSetMultiCovered{}
	_ bin.Decoder = &StickerSetMultiCovered{}

	_ StickerSetCoveredClass = &StickerSetMultiCovered{}
)

// StickerSetCoveredClass represents StickerSetCovered generic type.
//
// Example:
//  g, err := DecodeStickerSetCovered(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *StickerSetCovered: // stickerSetCovered#6410a5d2
//  case *StickerSetMultiCovered: // stickerSetMultiCovered#3407e51b
//  default: panic(v)
//  }
type StickerSetCoveredClass interface {
	bin.Encoder
	bin.Decoder
	construct() StickerSetCoveredClass
}

// DecodeStickerSetCovered implements binary de-serialization for StickerSetCoveredClass.
func DecodeStickerSetCovered(buf *bin.Buffer) (StickerSetCoveredClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StickerSetCoveredTypeID:
		// Decoding stickerSetCovered#6410a5d2.
		v := StickerSetCovered{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerSetCoveredClass: %w", err)
		}
		return &v, nil
	case StickerSetMultiCoveredTypeID:
		// Decoding stickerSetMultiCovered#3407e51b.
		v := StickerSetMultiCovered{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerSetCoveredClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StickerSetCoveredClass: %w", bin.NewUnexpectedID(id))
	}
}