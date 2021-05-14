package channels

import (
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func BindRoutes(routeGroup *echo.Group, db *bolt.DB) {
	routeGroup.GET("/", listHandler(db))
	routeGroup.POST("/", createHandler(db))
}
