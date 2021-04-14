package channel

import "github.com/segmentio/ksuid"

type VoiceChannel struct {
	ID ksuid.KSUID `json:"id"`
	Type    string `json:"type" validate:"oneof:text voice,required"`
	Name    string `json:"name" validate:"max=100,required"`
	Bitrate uint8  `json:"bitrate"`
}

func CreateVoice () *VoiceChannel {
	return &VoiceChannel{
		ID: ksuid.New(),
		Type: "voice",
	}
}