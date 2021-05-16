package message

import (
	"gitlab.com/grumblechat/server/internal/pagination"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func GetAll(db *bolt.DB, channelID *ksuid.KSUID, pgn *pagination.Pagination) ([]*Message, error) {
	var messages []*Message

	err := db.View(func(tx *bolt.Tx) (error) {
		dbb := channelBucket(tx, channelID)
		csr := dbb.Cursor()
		var ctr uint16 = 1

		// iterate over all messages, decode, and add to result
		for k, v := pgn.InitCursor(csr); ctr <= pgn.Count && k != nil; k, v = pgn.MoveCursor(csr) {
			decoded, err := Decode(v)
			if (err != nil) { return err }
			messages = append(messages, decoded)
			ctr++
		}

		k, _ := csr.Prev()
		endKey, err := ksuid.FromBytes(k)
		pgn.Cursor = endKey
		return err
	})

	return messages, err
}