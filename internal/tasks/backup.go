package tasks

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/deltabyte/grumble-server/internal/config"

	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func directory(cfg *config.Config) string {
	path := cfg.Paths.Backup
	now := time.Now()

	if cfg.Backup.Group {
		date := now.Format("2006-01-02")
		path = filepath.Join(path, date)
	}

	return path
}

func filename(name string, suffix string) string {
	fileDate := time.Now().Format("2006-01-02T1504")

	if suffix != "" {
		suffix = "-" + suffix
	}

	return fmt.Sprintf("%s-%s%s.gz", name, fileDate, suffix)
}

func uniqueFile(cfg *config.Config, name string) string {
	filePath := filepath.Join(directory(cfg), filename(name, ""))

	if _, err := os.Stat(filePath); os.IsExist(err) {
		id := ksuid.New().String()
		filePath = filepath.Join(directory(cfg), filename(name, id))
	}

	return filePath
}

func BackupDatabase(ctx TaskContext) error {
	filePath := uniqueFile(ctx.cfg, "db")

	// open file writer
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	//compress data
	gz := gzip.NewWriter(file)
	defer gz.Close()

	// stream DB through compressor and into file
	err = ctx.db.View(func(tx *bolt.Tx) error {
		ctx.log.Info("Starting DB backup")
		size, err := tx.WriteTo(gz)
		ctx.log.With(
			"file", filePath,
			"size", size,
		).Info("Completed DB backup")
		return err
	})

	return err
}
