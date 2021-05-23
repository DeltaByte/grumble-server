package tasks

import (
	"time"

	"github.com/grumblechat/server/internal/config"
	"github.com/grumblechat/server/internal/logging"
	"go.uber.org/zap"

	"github.com/go-co-op/gocron"
	bolt "go.etcd.io/bbolt"
)

type TaskContext struct {
	cfg *config.Config
	db *bolt.DB
	log *zap.SugaredLogger
}

func Run(db *bolt.DB, cfg *config.Config) {
	log := logging.Task()
	ctx := TaskContext{ cfg, db, log }
	scheduler := gocron.NewScheduler(time.Local)

	// backups
	bkuCtx := ctx; bkuCtx.log = log.Named("backup")
	scheduler.Every(cfg.Backup.Schedule).Do(BackupDatabase, bkuCtx)

	// start scheduler
	log.Info("starting task scheduler")
	scheduler.StartAsync()
}