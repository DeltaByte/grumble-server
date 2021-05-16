package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/grumblechat/server/internal/config"
	"github.com/grumblechat/server/internal/validation"
	"github.com/grumblechat/server/pkg/channel"
	"github.com/grumblechat/server/pkg/message"
	channelsController "github.com/grumblechat/server/internal/controllers/channels"
	messagesController "github.com/grumblechat/server/internal/controllers/messages"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	sentryEcho "github.com/getsentry/sentry-go/echo"
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
	config := config.Load()

	// initialize Sentry client
	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.SentryDSN,
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
	if config.EnableSentry {
		app.Use(sentryEcho.New(sentryEcho.Options{
			Repanic: true,
		}))
	}

	// load database
	db := initDB(config.Storage.Database)
	defer db.Close()

	// bind controller routes
	channelsController.BindRoutes(db, app.Group("/channels"))
	messagesController.BindRoutes(db, app.Group("/channels/:channelID/messages"))

	// start server
	app.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}
