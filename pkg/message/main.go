package message

import (
	"time"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

const BoltBucketName = "messages"

func New(channelID ksuid.KSUID) *Message {
	return &Message{
		ID: ksuid.New(),
		ChannelID: channelID,
		TTL:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Decode(enc []byte) (*Message, error) {
	msg := &Message{}
	err := msg.Decode(enc)
	return msg, err
}

func channelBucket(tx *bolt.Tx, msg *Message) *bolt.Bucket {
	msgBucket := tx.Bucket([]byte(BoltBucketName))
	return msgBucket.Bucket(msg.ChannelID.Bytes())
}