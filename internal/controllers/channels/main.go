package channels

import (
	"errors"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func BindRoutes(routeGroup *echo.Group, db *bolt.DB) {
	routeGroup.GET("/", listHandler)
	routeGroup.POST("/", createHandler(db))
}

var knownTypes = [2]string {"text", "voice"}

func listHandler(ctx echo.Context) error {
	return errors.New("not implemented")
}
