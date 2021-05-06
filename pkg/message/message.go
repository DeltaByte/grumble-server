package message

import "github.com/segmentio/ksuid"

type Message struct {
	ID      string `json:"id"`,
	Channel []byte `json:"channel"`,
	Payload []byte `json:"data"`,
	TTL     uint32 `json:"ttl"`,
}

func New(channel, payload []byte) *Message {
	return &Message{
		ID: ksuid.New(),
		Channel: channel,
		Payload: payload,
	}
}