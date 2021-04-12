package channel

import "github.com/segmentio/ksuid"

type Channel struct {
	ID ksuid.KSUID `json:"id"`
	Name  string   `json:"name" validate:"max=100,required"`

	// text only
	Topic string   `json:"topic" validate:"max=1024"`
	NSFW  bool     `json:"nsfw"`

	// voice only
	Bitrate uint8  `json:"bitrate"`
}