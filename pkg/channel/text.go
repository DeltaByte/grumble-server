package channel

import (
	"bytes"
	"encoding/gob"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

type TextChannel struct {
	ID    ksuid.KSUID `json:"id"`
	Type  string      `json:"type" validate:"oneof=text voice,required"`
	Name  string      `json:"name" validate:"max=100,required"`
	Topic string      `json:"topic" validate:"max=1024"`
	NSFW  bool        `json:"nsfw" default:"false"`
}

func NewText() *TextChannel {
	return &TextChannel{
		ID:   ksuid.New(),
		Type: "text",
	}
}

func (tc *TextChannel) Encode() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(tc)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (tc *TextChannel) Decode(enc []byte) error {
	buf := bytes.NewBuffer(enc)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(tc)
	return err
}

func (tc *TextChannel) Save(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		// byte-encode the channel
		enc, err := tc.Encode()
		if err != nil { return err }

		// persist the channel
		dbb := tx.Bucket([]byte(DBBucket))
		err = dbb.Put(tc.ID.Bytes(), enc)

		return err
	})
}
