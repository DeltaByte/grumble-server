package message

import (
	"bytes"
	"encoding/gob"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

type Message struct {
	ID        ksuid.KSUID `json:"id"`
	ChannelID ksuid.KSUID `json:"channel_id"`
	Body      string      `json:"data"`
	TTL       uint32      `json:"ttl"`
}

func (msg *Message) Encode() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(msg)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (msg *Message) Decode(enc []byte) error {
	buf := bytes.NewBuffer(enc)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(msg)
	return err
}

func (msg *Message) Save(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		// byte-encode the channel
		enc, err := msg.Encode()
		if err != nil { return err }

		// persist the channel
		dbb := tx.Bucket(DBBucket(msg))
		err = dbb.Put(msg.ID.Bytes(), enc)

		return err
	})
}