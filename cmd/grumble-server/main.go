package main

import (
	"fmt"
	"log"
	"time"

	"github.com/grumblechat/server/internal/backup"
	"github.com/grumblechat/server/internal/config"
	channelsController "github.com/grumblechat/server/internal/controllers/channels"
	messagesController "github.com/grumblechat/server/internal/controllers/messages"
	"github.com/grumblechat/server/internal/database"
	"github.com/grumblechat/server/internal/middleware"
	"github.com/grumblechat/server/internal/validation"

	"github.com/getsentry/sentry-go"
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// load config
	cfg := config.Load()

	// initialize Sentry client
	sentryOpts := sentry.ClientOptions{ Dsn: cfg.Sentry.DSN }
	if err := sentry.Init(sentryOpts); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}

	// init framework
	app := echo.New()
	app.HideBanner = true
	app.Validator = validation.Echo()
	app.Pre(echoMiddleware.AddTrailingSlash())

	// register global middleware
	app.Use(echoMiddleware.Logger())
	app.Use(echoMiddleware.Recover())
	app.Use(echoMiddleware.RequestID())
	app.Use(middleware.PoweredBy())

	// report errors to sentry
	if cfg.Sentry.Enabled {
		app.Use(sentryEcho.New(sentryEcho.Options{
			Repanic: true,
		}))
	}

	// load database
	db := database.Init(cfg.Paths.Database)
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
