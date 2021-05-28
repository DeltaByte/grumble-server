package tasks

import (
	"time"

	"github.com/deltabyte/grumble-server/internal/config"
	"github.com/deltabyte/grumble-server/internal/logging"
	"go.uber.org/zap"

	"github.com/go-co-op/gocron"
	bolt "go.etcd.io/bbolt"
)

type TaskContext struct {
	cfg *config.Config
	db  *bolt.DB
	log *zap.SugaredLogger
}

func Run(db *bolt.DB, cfg *config.Config) {
	log := logging.Task()
	defer log.Sync()

	// setup scheduler
	ctx := TaskContext{cfg, db, log}
	scheduler := gocron.NewScheduler(time.Local)

	// backups
	bkuCtx := ctx
	bkuCtx.log = log.Named("backup")
	scheduler.Every(cfg.Backup.Schedule).Do(BackupDatabase, bkuCtx)

	// start scheduler
	log.Info("starting task scheduler")
	scheduler.StartAsync()
}
