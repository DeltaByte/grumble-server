package database

import (
	"path/filepath"

	bolt "go.etcd.io/bbolt"
)

func initBucket(tx *bolt.Tx, name string) error {
	_, err := tx.CreateBucketIfNotExists([]byte(name))
	return err
}

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
		if err := initBucket(tx, "channels"); err != nil {
			return err
		}

		// messages
		if err := initBucket(tx, "messages"); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic("Failed to migrate DB")
	}

	return db
}