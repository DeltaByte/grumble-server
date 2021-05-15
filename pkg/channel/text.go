package channel

import (
	"bytes"
	"time"
	"encoding/gob"

	"github.com/segmentio/ksuid"
	"gitlab.com/grumblechat/server/pkg/message"
	bolt "go.etcd.io/bbolt"
)

func NewText() *TextChannel {
	return &TextChannel{
		ID:   ksuid.New(),
		Type: "text",
	}
}

type TextChannel struct {
	ID        ksuid.KSUID `json:"id" copier:"-"`
	Type      string      `json:"type" validate:"eq=text,required"`
	Name      string      `json:"name" validate:"max=100,required"`
	Topic     string      `json:"topic" validate:"max=1024"`
	NSFW      bool        `json:"nsfw"`
	CreatedAt time.Time   `json:"created_at" copier:"-"`
	UpdatedAt time.Time   `json:"updated_at" copier:"-"`
}

func (tc *TextChannel) GetType() string {
	return tc.Type
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
		dbb := tx.Bucket([]byte(BoltBucketName))
		err = dbb.Put(tc.ID.Bytes(), enc)
		if (err != nil) { return err }

		// create a bucket for messages
		msgBucket := tx.Bucket([]byte(message.BoltBucketName))
		_, err = msgBucket.CreateBucketIfNotExists(tc.ID.Bytes())

		// assumed that err is either an error or nil by this point
		return err
	})
}
