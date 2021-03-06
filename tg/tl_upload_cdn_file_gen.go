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

// UploadCdnFileReuploadNeeded represents TL type `upload.cdnFileReuploadNeeded#eea8e46e`.
// The file was cleared from the temporary RAM cache of the CDN¹ and has to be reuploaded.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.cdnFileReuploadNeeded for reference.
type UploadCdnFileReuploadNeeded struct {
	// Request token (see CDN¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	RequestToken []byte
}

// UploadCdnFileReuploadNeededTypeID is TL type id of UploadCdnFileReuploadNeeded.
const UploadCdnFileReuploadNeededTypeID = 0xeea8e46e

// String implements fmt.Stringer.
func (c *UploadCdnFileReuploadNeeded) String() string {
	if c == nil {
		return "UploadCdnFileReuploadNeeded(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UploadCdnFileReuploadNeeded")
	sb.WriteString("{\n")
	sb.WriteString("\tRequestToken: ")
	sb.WriteString(fmt.Sprint(c.RequestToken))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *UploadCdnFileReuploadNeeded) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFileReuploadNeeded#eea8e46e as nil")
	}
	b.PutID(UploadCdnFileReuploadNeededTypeID)
	b.PutBytes(c.RequestToken)
	return nil
}

// Decode implements bin.Decoder.
func (c *UploadCdnFileReuploadNeeded) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFileReuploadNeeded#eea8e46e to nil")
	}
	if err := b.ConsumeID(UploadCdnFileReuploadNeededTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.cdnFileReuploadNeeded#eea8e46e: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.cdnFileReuploadNeeded#eea8e46e: field request_token: %w", err)
		}
		c.RequestToken = value
	}
	return nil
}

// construct implements constructor of UploadCdnFileClass.
func (c UploadCdnFileReuploadNeeded) construct() UploadCdnFileClass { return &c }

// Ensuring interfaces in compile-time for UploadCdnFileReuploadNeeded.
var (
	_ bin.Encoder = &UploadCdnFileReuploadNeeded{}
	_ bin.Decoder = &UploadCdnFileReuploadNeeded{}

	_ UploadCdnFileClass = &UploadCdnFileReuploadNeeded{}
)

// UploadCdnFile represents TL type `upload.cdnFile#a99fca4f`.
// Represent a chunk of a CDN¹ file.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.cdnFile for reference.
type UploadCdnFile struct {
	// The data
	Bytes []byte
}

// UploadCdnFileTypeID is TL type id of UploadCdnFile.
const UploadCdnFileTypeID = 0xa99fca4f

// String implements fmt.Stringer.
func (c *UploadCdnFile) String() string {
	if c == nil {
		return "UploadCdnFile(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UploadCdnFile")
	sb.WriteString("{\n")
	sb.WriteString("\tBytes: ")
	sb.WriteString(fmt.Sprint(c.Bytes))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *UploadCdnFile) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFile#a99fca4f as nil")
	}
	b.PutID(UploadCdnFileTypeID)
	b.PutBytes(c.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (c *UploadCdnFile) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFile#a99fca4f to nil")
	}
	if err := b.ConsumeID(UploadCdnFileTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.cdnFile#a99fca4f: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.cdnFile#a99fca4f: field bytes: %w", err)
		}
		c.Bytes = value
	}
	return nil
}

// construct implements constructor of UploadCdnFileClass.
func (c UploadCdnFile) construct() UploadCdnFileClass { return &c }

// Ensuring interfaces in compile-time for UploadCdnFile.
var (
	_ bin.Encoder = &UploadCdnFile{}
	_ bin.Decoder = &UploadCdnFile{}

	_ UploadCdnFileClass = &UploadCdnFile{}
)

// UploadCdnFileClass represents upload.CdnFile generic type.
//
// See https://core.telegram.org/type/upload.CdnFile for reference.
//
// Example:
//  g, err := DecodeUploadCdnFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UploadCdnFileReuploadNeeded: // upload.cdnFileReuploadNeeded#eea8e46e
//  case *UploadCdnFile: // upload.cdnFile#a99fca4f
//  default: panic(v)
//  }
type UploadCdnFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() UploadCdnFileClass
	fmt.Stringer
}

// DecodeUploadCdnFile implements binary de-serialization for UploadCdnFileClass.
func DecodeUploadCdnFile(buf *bin.Buffer) (UploadCdnFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UploadCdnFileReuploadNeededTypeID:
		// Decoding upload.cdnFileReuploadNeeded#eea8e46e.
		v := UploadCdnFileReuploadNeeded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadCdnFileClass: %w", err)
		}
		return &v, nil
	case UploadCdnFileTypeID:
		// Decoding upload.cdnFile#a99fca4f.
		v := UploadCdnFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadCdnFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UploadCdnFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// UploadCdnFile boxes the UploadCdnFileClass providing a helper.
type UploadCdnFileBox struct {
	CdnFile UploadCdnFileClass
}

// Decode implements bin.Decoder for UploadCdnFileBox.
func (b *UploadCdnFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UploadCdnFileBox to nil")
	}
	v, err := DecodeUploadCdnFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CdnFile = v
	return nil
}

// Encode implements bin.Encode for UploadCdnFileBox.
func (b *UploadCdnFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CdnFile == nil {
		return fmt.Errorf("unable to encode UploadCdnFileClass as nil")
	}
	return b.CdnFile.Encode(buf)
}
