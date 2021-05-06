package channel

import (
	"bytes"
	"encoding/gob"

	"github.com/dgraph-io/badger/v3"
	"github.com/segmentio/ksuid"
)

type TextChannel struct {
	ID    ksuid.KSUID `json:"id"`
	Type  string      `json:"type" validate:"oneof:text voice,required"`
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

func (tc *TextChannel) Save(db *badger.DB) error {
	err := db.Update(func(txn *badger.Txn) error {
		// byte-encode the channel
		enc, err := tc.Encode()
		if err != nil { return err }

		// persist the channel
		err = txn.Set(tc.ID.Bytes(), enc)
		return err
	})
	return err
}
