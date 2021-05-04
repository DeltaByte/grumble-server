package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	channelsController "gitlab.com/grumblechat/server/controllers/channels"
	"gitlab.com/grumblechat/server/internal/config"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
  }
  return nil
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

	// init framework and register global middleware
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// report errors to sentry
	if (config.EnableSentry) {
		app.Use(sentryEcho.New(sentryEcho.Options{
			Repanic: true,
		}))
	}

	// setup validation
	app.Validator = &CustomValidator{
		validator: validator.New(),
	}

	// bind controller routes
	channelsController.BindRoutes(app.Group("/channels"))

	// start server
	app.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}