package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/grumblechat/server/controllers"
)

func main() {
	// load config
	config := LoadConfig()

	// init framework
	app := echo.New()
	app.Use(middleware.Recover())

	// bind controller routes
	controllers.BindChannelRoutes(app.Group("/channels"))

	// start server
	app.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}