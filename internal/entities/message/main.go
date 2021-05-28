package message

import (
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

const BoltBucketName = "messages"

func New(channelID ksuid.KSUID) *Message {
	return &Message{
		ID: ksuid.New(),
		ChannelID: channelID,
		TTL:  0,
	}
}

func Decode(enc []byte) (*Message, error) {
	msg := &Message{}
	err := msg.Decode(enc)
	return msg, err
}

// NOTE: the BoltDB bucket should have been created as part of saving the Channel
func channelBucket(tx *bolt.Tx, channelID ksuid.KSUID) *bolt.Bucket {
	msgBucket := tx.Bucket([]byte(BoltBucketName))
	return msgBucket.Bucket(channelID.Bytes())
}