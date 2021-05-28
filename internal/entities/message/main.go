package message

import (
	"time"

	"github.com/deltabyte/grumble-server/internal/helpers"
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


func BatchSave(db *bolt.DB, messages []*Message) error {
	return db.Batch(func(tx *bolt.Tx) error {
		now := time.Now()
		buckets := make(map[ksuid.KSUID]*bolt.Bucket)

		// save each channel
		for _, msg := range messages {
			// update timestamps
			msg.CreatedAt = helpers.TouchTimestamp(msg.CreatedAt, now, true)
			msg.UpdatedAt = helpers.TouchTimestamp(msg.UpdatedAt, now, false)

			// get channel-specific bucket
			bkt, ok := buckets[msg.ChannelID]
			if (!ok) {
				bkt = channelBucket(tx, msg.ChannelID)
				buckets[msg.ChannelID] = bkt
			}

			// persist the channel
			if err := msg.bktSave(bkt); err != nil {
				return err
			}
		}

		return nil
	})
}