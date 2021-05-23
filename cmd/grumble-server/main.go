package main

import (
	"fmt"
	"log"
	"time"

	"github.com/grumblechat/server/internal/config"
	channelsController "github.com/grumblechat/server/internal/controllers/channels"
	messagesController "github.com/grumblechat/server/internal/controllers/messages"
	"github.com/grumblechat/server/internal/database"
	"github.com/grumblechat/server/internal/logging"
	"github.com/grumblechat/server/internal/middleware"
	"github.com/grumblechat/server/internal/tasks"
	"github.com/grumblechat/server/internal/validation"

	"github.com/getsentry/sentry-go"
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var version string

func main() {
	// load config
	cfg := config.Load()

	if (cfg.Banner) {
		printBanner()
	}

	// initialize Sentry client
	sentryOpts := sentry.ClientOptions{Dsn: cfg.Sentry.DSN}
	if err := sentry.Init(sentryOpts); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}

	// init backend crap
	logging.Init()
	log := logging.Default()
	db := database.Init(cfg.Paths.Database)
	defer db.Close()

	// init framework
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true
	app.Validator = validation.Echo()
	app.Pre(echoMiddleware.RemoveTrailingSlash())

	// register global middleware
	app.Use(middleware.Logger(cfg))
	app.Use(echoMiddleware.Recover())
	app.Use(echoMiddleware.RequestID())
	app.Use(middleware.PoweredBy(version))

	// report errors to sentry
	if cfg.Sentry.Enabled {
		app.Use(sentryEcho.New(sentryEcho.Options{
			Repanic: true,
		}))
	}

	// bind controller routes
	channelsController.BindRoutes(db, app.Group("/channels"))
	messagesController.BindRoutes(db, app.Group("/channels/:channelID/messages"))

	// start task scheduler
	tasks.Run(db, cfg)

	// start server
	tasks.StartAsync()
	app.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
}
