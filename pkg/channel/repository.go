package channel

import (
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func GetAll(db *bolt.DB) ([]Channel, error) {
	var channels []Channel

	err := db.View(func(tx *bolt.Tx) (error) {
		dbb := tx.Bucket([]byte(BoltBucketName))

		// iterate over all channels, decode, and add to result
		dbb.ForEach(func(k, v []byte) error {
			decoded, err := Decode(v)
			if (err != nil) { return err }
			channels = append(channels, decoded)
			return nil
		})

		return nil
	})

	return channels, err
}

func Find(db *bolt.DB, id ksuid.KSUID) (Channel, error) {
	var channel Channel

	err := db.View(func(tx *bolt.Tx) error {
		dbb := tx.Bucket([]byte(BoltBucketName))

		// get by ID
		res := dbb.Get(id.Bytes())
		if (res == nil) {
			channel = nil
			return nil
		}

		// decode channel
		decoded, err := Decode(res)
		if (err != nil) { return err }

		channel = decoded
		return nil
	})

	return channel, err
}