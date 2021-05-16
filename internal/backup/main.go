package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/grumblechat/server/internal/config"
	"github.com/segmentio/ksuid"
)

func directory(cfg *config.Config) string {
	path := cfg.Storage.Backup
	now  := time.Now()

	if (cfg.Backup.Group) {
		date := now.Format("2006-01-02")
		path = filepath.Join(path, date)
	}

	return path
}

func filename(name string, suffix string) string {
	fileDate := time.Now().Format("2006-01-02T1504")

	if (suffix != "") {
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