package proto

import (
	"errors"

	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
)

// MessageContainerTypeID is TL type id of MessageContainer.
const MessageContainerTypeID = 0x73f1f8dc

// MessageContainer contains slice of messages.
type MessageContainer struct {
	Messages []Message
}

// Decode implements bin.Decoder.
func (m *MessageContainer) Decode(b *bin.Buffer) error {
	if err := b.ConsumeID(MessageContainerTypeID); err != nil {
		return xerrors.Errorf("failed to consume id of message container: %w", err)
	}
	n, err := b.Int()
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		var msg Message
		if err := msg.Decode(b); err != nil {
			return err
		}
		m.Messages = append(m.Messages, msg)
	}
	return nil
}

// Message is element of MessageContainer.
type Message struct {
	ID    int64
	SeqNo int
	Bytes int
	Body  []byte
}

// Decode implements bin.Decoder.
func (m *Message) Decode(b *bin.Buffer) error {
	{
		v, err := b.Long()
		if err != nil {
			return err
		}
		m.ID = v
	}
	{
		v, err := b.Int()
		if err != nil {
			return err
		}
		m.SeqNo = v
	}
	{
		v, err := b.Int()
		if err != nil {
			return err
		}
		m.Bytes = v
	}
	if m.Bytes < 0 || m.Bytes > 1024*1024 {
		return errors.New("message length is too big")
	}
	m.Body = make([]byte, m.Bytes)
	return b.ConsumeN(m.Body, m.Bytes)
}
