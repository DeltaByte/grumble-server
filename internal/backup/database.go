package backup

import (
	"compress/gzip"
	"log"
	"os"

	"github.com/grumblechat/server/internal/config"

	bolt "go.etcd.io/bbolt"
)

func Database(cfg *config.Config, db *bolt.DB) error {
	filePath := uniqueFile(cfg, "db")

	// open file writer
	file, err := os.Create(filePath)
	if (err != nil) {
		return err
	}
	defer file.Close()

	//compress data
	gz := gzip.NewWriter(file)
	defer gz.Close()
	
	// stream DB through compressor and into file
	err = db.View(func(tx *bolt.Tx) error {
		foo, err := tx.WriteTo(file)
		log.Printf("Backing up (%d)", foo)
		return err
	})

	return err
}