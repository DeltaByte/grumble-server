package channel

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/grumblechat/server/pkg/helpers"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func NewVoice() *VoiceChannel {
	return &VoiceChannel{
		ID:   ksuid.New(),
		Type: "voice",
		Bitrate: 64,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type VoiceChannel struct {
	ID        ksuid.KSUID `json:"id" copier:"-"`
	Type      string      `json:"type" validate:"eq=voice,required"`
	Name      string      `json:"name" validate:"max=100,required"`
	Bitrate   uint8       `json:"bitrate" validate:"min=4,max=255"`
	CreatedAt time.Time   `json:"created_at" copier:"-"`
	UpdatedAt time.Time   `json:"updated_at" copier:"-"`
}

func (vc *VoiceChannel) GetType() string {
	return vc.Type
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

func (vc *VoiceChannel) Decode(enc []byte) error {
	buf := bytes.NewBuffer(enc)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(vc)
	return err
}


func (vc *VoiceChannel) Save(db *bolt.DB) error {
	// update timestamps
	now := time.Now()
	vc.CreatedAt = helpers.TouchTimestamp(vc.CreatedAt, now, true)
	vc.UpdatedAt = helpers.TouchTimestamp(vc.UpdatedAt, now, false)

	// persist to DB
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
