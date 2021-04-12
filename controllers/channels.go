package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	"gitlab.com/grumblechat/server/internal/channel"
)

func BindChannelRoutes(routeGroup *echo.Group) {
	routeGroup.GET("/", listHandler)
	routeGroup.POST("/", createHandler)
}

func listHandler(ctx echo.Context) error {
	//
}

func createHandler(ctx echo.Context) error {
	newChannel := &channel.Channel{
		ID: ksuid.New(),
	}

	// unmarshal request body
	if err := ctx.Bind(newChannel); err != nil {
		return ctx.JSON(http.StatusBadRequest, newChannel)
	}

	// TODO: validate data

	// TODO: write data to store

	// TODO: respond to client
}