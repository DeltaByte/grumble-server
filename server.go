package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/grumblechat/server/controllers"
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
	config := LoadConfig()

	// init framework
	app := echo.New()
	app.Use(middleware.Recover())

	// setup validation
	app.Validator = &CustomValidator{
		validator: validator.New(),
	}

	// bind controller routes
	controllers.BindChannelRoutes(app.Group("/channels"))

	// start server
	app.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}