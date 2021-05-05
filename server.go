package main

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gitlab.com/grumblechat/server/internal/config"
	"gitlab.com/grumblechat/server/internal/validation"
	channelsController "gitlab.com/grumblechat/server/internal/controllers/channels"
)

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

	// bind controller routes
	channelsController.BindRoutes(app.Group("/channels"))

	// start server
	app.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}
