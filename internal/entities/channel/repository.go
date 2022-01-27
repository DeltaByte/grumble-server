package channel

import (
	pb "github.com/grumblechat/server/gen/go/channel"
	"github.com/grumblechat/server/internal/helpers"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

type ChannelRepository struct{}

func GetAll(db *bolt.DB) ([]*pb.Channel, error) {
	var channels []*pb.Channel

	err := db.View(func(tx *bolt.Tx) error {
		dbb := tx.Bucket([]byte("channels"))

		// iterate over all channels, decode, and add to result
		dbb.ForEach(func(k, v []byte) error {
			decoded, err := Decode(v)
			if err != nil {
				return err
			}
			channels = append(channels, decoded)
			return nil
		})

		return nil
	})

	return channels, err
}

func Find(db *bolt.DB, id ksuid.KSUID) (*pb.Channel, error) {
	var channel *pb.Channel

	err := db.View(func(tx *bolt.Tx) error {
		dbb := tx.Bucket([]byte("channels"))

		// get by ID
		res := dbb.Get(id.Bytes())
		if res == nil {
			channel = nil
			return nil
		}

		// decode channel
		decoded, err := Decode(res)
		if err != nil {
			return err
		}

		channel = decoded
		return nil
	})

	return channel, err
}

func Delete(db *bolt.DB, chn *pb.Channel) error {
	// get ID
	id, err := helpers.ParseKSUIDBytes(chn.Id)
	if err != nil {
		return err
	}

	// remove from DB
	return db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("channels"))

		// delete message bucket
		msgBucket := tx.Bucket([]byte("messages"))
		if err := msgBucket.DeleteBucket(id); err != nil {
			return err
		}

		// delete self
		err := bkt.Delete(id)

		// assumed that err is either an error or nil by this point
		return err
	})
}
