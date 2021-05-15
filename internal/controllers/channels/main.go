package channelsController

import (
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func BindRoutes(db *bolt.DB, router *echo.Group) {
	router.GET("/", listHandler(db))
	router.POST("/", createHandler(db))
}
