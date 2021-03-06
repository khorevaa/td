package tgtest

import (
	"context"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/transport"
)

func (s *Server) writeUnencrypted(ctx context.Context, conn transport.Conn, data bin.Encoder) error {
	b := &bin.Buffer{}
	if err := data.Encode(b); err != nil {
		return err
	}
	msg := proto.UnencryptedMessage{
		MessageID:   int64(proto.NewMessageID(s.clock(), proto.MessageServerResponse)),
		MessageData: b.Copy(),
	}
	b.Reset()
	if err := msg.Encode(b); err != nil {
		return err
	}

	return conn.Send(ctx, b)
}

func (s *Server) readUnencrypted(ctx context.Context, conn transport.Conn, data bin.Decoder) error {
	b := &bin.Buffer{}
	if err := conn.Recv(ctx, b); err != nil {
		return err
	}

	return s.decodeUnencrypted(b, data)
}

func (s *Server) decodeUnencrypted(b *bin.Buffer, data bin.Decoder) error {
	var msg proto.UnencryptedMessage
	if err := msg.Decode(b); err != nil {
		return err
	}
	if err := s.checkMsgID(msg.MessageID); err != nil {
		return err
	}
	b.ResetTo(msg.MessageData)

	return data.Decode(b)
}
