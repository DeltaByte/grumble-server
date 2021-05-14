package channel

import (
	"bytes"
	"encoding/gob"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

type VoiceChannel struct {
	ID      ksuid.KSUID `json:"id"`
	Type    string      `json:"type" validate:"oneof=text voice,required"`
	Name    string      `json:"name" validate:"max=100,required"`
	Bitrate uint8       `json:"bitrate"`
}

func NewVoice() *VoiceChannel {
	return &VoiceChannel{
		ID:   ksuid.New(),
		Type: "voice",
	}
}

func (vc *VoiceChannel) Encode() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(vc)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (vc *VoiceChannel) Save(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		// byte-encode the channel
		enc, err := vc.Encode()
		if err != nil { return err }

		// persist the channel
		dbb := tx.Bucket([]byte("channels"))
		err = dbb.Put(vc.ID.Bytes(), enc)

		return err
	})
}
