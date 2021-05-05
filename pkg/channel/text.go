package channel

import "github.com/segmentio/ksuid"

type TextChannel struct {
	ID ksuid.KSUID `json:"id"`
	Type  string   `json:"type" validate:"oneof:text voice,required"`
	Name  string   `json:"name" validate:"max=100,required"`
	Topic string   `json:"topic" validate:"max=1024"`
	NSFW  bool     `json:"nsfw" default:"false"`
}

func CreateText () *TextChannel {
	return &TextChannel{
		ID: ksuid.New(),
		Type: "text",
	}
}