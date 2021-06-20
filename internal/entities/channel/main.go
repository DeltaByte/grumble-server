package channel

import (
	"time"

	// internal
	pb "github.com/grumblechat/server/gen/go/channel"
	"github.com/grumblechat/server/internal/helpers"

	// external
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

func Decode(encoded []byte) (*pb.Channel, error) {
	channel := &pb.Channel{}
	if err := proto.Unmarshal(encoded, channel); err != nil {
		return nil, err
	}
	return channel, nil
}

func New(chanType pb.ChannelType) *pb.Channel {
	chn := &pb.Channel{
		Id: ksuid.New().String(),
	}

	// voice specific defaults
	if chanType == pb.ChannelType_CHANNEL_TYPE_VOICE {
		chn.Bitrate = 64
	}

	return chn
}

func Save(db *bolt.DB, chn *pb.Channel) error {
	// update timestamps
	now := time.Now()
	helpers.TouchTimestamp(chn.CreatedAt, now, true)
	helpers.TouchTimestamp(chn.UpdatedAt, now, false)

	// encode to binary
	enc, err := proto.Marshal(chn)
	if err != nil {
		return err
	}

	// get ID
	id, err := helpers.ParseKSUIDBytes(chn.Id)
	if err != nil {
		return err
	}

	// persist to DB
	return db.Update(func(tx *bolt.Tx) error {
		// persist the channel
		bkt := tx.Bucket([]byte("channels"))
		err = bkt.Put(id, enc)
		if err != nil {
			return err
		}

		// create a bucket for messages
		msgBkt := tx.Bucket([]byte("messages"))
		_, err = msgBkt.CreateBucketIfNotExists(id)

		// assumed that err is either an error or nil by this point
		return err
	})
}
