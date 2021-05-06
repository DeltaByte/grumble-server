package channel

import (
	"bytes"
	"encoding/gob"

	"github.com/dgraph-io/badger/v3"
	"github.com/segmentio/ksuid"
)

type VoiceChannel struct {
	ID      ksuid.KSUID `json:"id"`
	Type    string      `json:"type" validate:"oneof:text voice,required"`
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

func (vc *VoiceChannel) Save(db *badger.DB) error {
	err := db.Update(func(txn *badger.Txn) error {
		// byte-encode the channel
		enc, err := vc.Encode()
		if err != nil { return err }

		// persist the channel
		err = txn.Set(vc.ID.Bytes(), enc)
		return err
	})
	return err
}
