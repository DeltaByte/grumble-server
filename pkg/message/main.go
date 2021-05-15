package message

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

func DBBucket(msg *Message) []byte {
	name := fmt.Sprintf("messages:%s", msg.ChannelID.String())
	return []byte(name)
}

func New(channelID ksuid.KSUID) *Message {
	return &Message{
		ID: ksuid.New(),
		ChannelID: channelID,
		TTL:  0,
	}
}

func Decode(enc []byte) (*Message, error) {
	msg := &Message{}
	err := msg.Decode(enc)
	return msg, err
}