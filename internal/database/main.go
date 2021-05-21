package database

import (
	"path/filepath"

	"github.com/grumblechat/server/pkg/channel"
	"github.com/grumblechat/server/pkg/message"

	bolt "go.etcd.io/bbolt"
)

func Init(path string) *bolt.DB {
	// open BoltDB
	dbPath := filepath.Join(path, "grumble.db")
	db, err := bolt.Open(dbPath, 0666, nil)

	if err != nil {
		panic("Failed to open database")
	}

	// ensure that buckets exist
	err = db.Update(func(tx *bolt.Tx) error {
		// channels
		_, err := tx.CreateBucketIfNotExists([]byte(channel.BoltBucketName))
		if err != nil {
			return err
		}

		// messages
		_, err = tx.CreateBucketIfNotExists([]byte(message.BoltBucketName))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic("Failed to migrate DB")
	}

	return db
}