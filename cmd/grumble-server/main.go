package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	if cfg.Banner {
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
	defer log.Sync()
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
	go func() {
		host := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		log.With("address", host).Info("starting HTTP server")
		if err := app.Start(host); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 60 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	// allow up to 60secs for connections to drain and tasks to complete
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
