package channels

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func BindRoutes(routeGroup *echo.Group) {
	routeGroup.GET("/", listHandler)
	routeGroup.POST("/", createHandler)
}

var knownTypes = [2]string {"text", "voice"}

func listHandler(ctx echo.Context) error {
	return errors.New("Not implemented")
}

