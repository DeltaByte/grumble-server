package message

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/grumblechat/server/internal/helpers"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

type Message struct {
	ID        ksuid.KSUID `json:"id" copier:"-"`
	ChannelID ksuid.KSUID `json:"channel_id"`
	Body      string      `json:"body" validate:"min=1,max=2048,required"`
	TTL       uint32      `json:"ttl,omitempty" validate:"max=2592000"`
	CreatedAt time.Time   `json:"created_at" copier:"-"`
	UpdatedAt time.Time   `json:"updated_at" copier:"-"`
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

// put the message into a specific DB bucket
func (msg *Message) bktSave(bkt *bolt.Bucket) error {
	// byte-encode the channel
	enc, err := msg.Encode()
	if err != nil {
		return err
	}

	// persist the channel
	err = bkt.Put(msg.ID.Bytes(), enc)

	return err
}

// persist the message to the DB
func (msg *Message) Save(db *bolt.DB) error {
	// update timestamps
	now := time.Now()
	msg.CreatedAt = helpers.TouchTimestamp(msg.CreatedAt, now, true)
	msg.UpdatedAt = helpers.TouchTimestamp(msg.UpdatedAt, now, false)

	// persist to DB
	return db.Update(func(tx *bolt.Tx) error {
		bkt := channelBucket(tx, msg.ChannelID)
		return msg.bktSave(bkt)
	})
}

func (msg *Message) Delete(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		dbb := channelBucket(tx, msg.ChannelID)

		// delete self
		err := dbb.Delete(msg.ID.Bytes())

		// assumed that err is either an error or nil by this point
		return err
	})
}
