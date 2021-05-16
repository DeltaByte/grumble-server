package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/grumblechat/server/internal/backup"
	"github.com/grumblechat/server/internal/config"
	channelsController "github.com/grumblechat/server/internal/controllers/channels"
	messagesController "github.com/grumblechat/server/internal/controllers/messages"
	"github.com/grumblechat/server/internal/validation"
	"github.com/grumblechat/server/pkg/channel"
	"github.com/grumblechat/server/pkg/message"

	"github.com/getsentry/sentry-go"
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	bolt "go.etcd.io/bbolt"
)

func initDB(path string) *bolt.DB {
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

func main() {
	// load config
	cfg := config.Load()

	// initialize Sentry client
	err := sentry.Init(sentry.ClientOptions{
		Dsn: cfg.Sentry.DSN,
	})
	if err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}

	// init framework
	app := echo.New()
	app.HideBanner = true
	app.Validator = validation.Echo()
	app.Pre(middleware.AddTrailingSlash())

	// register global middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// report errors to sentry
	if cfg.Sentry.Enable {
		app.Use(sentryEcho.New(sentryEcho.Options{
			Repanic: true,
		}))
	}

	// load database
	db := initDB(cfg.Storage.Database)
	defer db.Close()

	// schedule tasks
	tasks := gocron.NewScheduler(time.Local)
	tasks.Every(cfg.Backup.Schedule).Do(backup.Database, cfg, db)

	// bind controller routes
	channelsController.BindRoutes(db, app.Group("/channels"))
	messagesController.BindRoutes(db, app.Group("/channels/:channelID/messages"))

	// start server
	tasks.StartAsync()
	app.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
}
