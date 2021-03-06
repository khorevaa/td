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

// MessagesHighScores represents TL type `messages.highScores#9a3bfd99`.
// Highscores in a game
//
// See https://core.telegram.org/constructor/messages.highScores for reference.
type MessagesHighScores struct {
	// Highscores
	Scores []HighScore
	// Users, associated to the highscores
	Users []UserClass
}

// MessagesHighScoresTypeID is TL type id of MessagesHighScores.
const MessagesHighScoresTypeID = 0x9a3bfd99

// String implements fmt.Stringer.
func (h *MessagesHighScores) String() string {
	if h == nil {
		return "MessagesHighScores(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesHighScores")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range h.Scores {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range h.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (h *MessagesHighScores) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.highScores#9a3bfd99 as nil")
	}
	b.PutID(MessagesHighScoresTypeID)
	b.PutVectorHeader(len(h.Scores))
	for idx, v := range h.Scores {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field scores element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(h.Users))
	for idx, v := range h.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *MessagesHighScores) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.highScores#9a3bfd99 to nil")
	}
	if err := b.ConsumeID(MessagesHighScoresTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field scores: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value HighScore
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field scores: %w", err)
			}
			h.Scores = append(h.Scores, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field users: %w", err)
			}
			h.Users = append(h.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesHighScores.
var (
	_ bin.Encoder = &MessagesHighScores{}
	_ bin.Decoder = &MessagesHighScores{}
)
