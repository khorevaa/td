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

// MessagesInactiveChats represents TL type `messages.inactiveChats#a927fec5`.
// Inactive chat list
//
// See https://core.telegram.org/constructor/messages.inactiveChats for reference.
type MessagesInactiveChats struct {
	// When was the chat last active
	Dates []int
	// Chat list
	Chats []ChatClass
	// Users mentioned in the chat list
	Users []UserClass
}

// MessagesInactiveChatsTypeID is TL type id of MessagesInactiveChats.
const MessagesInactiveChatsTypeID = 0xa927fec5

// String implements fmt.Stringer.
func (i *MessagesInactiveChats) String() string {
	if i == nil {
		return "MessagesInactiveChats(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesInactiveChats")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range i.Dates {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range i.Chats {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range i.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *MessagesInactiveChats) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.inactiveChats#a927fec5 as nil")
	}
	b.PutID(MessagesInactiveChatsTypeID)
	b.PutVectorHeader(len(i.Dates))
	for _, v := range i.Dates {
		b.PutInt(v)
	}
	b.PutVectorHeader(len(i.Chats))
	for idx, v := range i.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *MessagesInactiveChats) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.inactiveChats#a927fec5 to nil")
	}
	if err := b.ConsumeID(MessagesInactiveChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field dates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field dates: %w", err)
			}
			i.Dates = append(i.Dates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesInactiveChats.
var (
	_ bin.Encoder = &MessagesInactiveChats{}
	_ bin.Decoder = &MessagesInactiveChats{}
)
