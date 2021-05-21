package channelsController

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

var copierOptions = copier.Option{
	IgnoreEmpty: true,
}

func BindRoutes(db *bolt.DB, router *echo.Group) {
	router.GET("", listHandler(db))
	router.POST("", createHandler(db))
	router.PUT("/:id", updateHandler(db))
}

type channelDTO struct {
	Type    string `json:"type" validate:"oneof=text voice,required"`
	Name    string `json:"name"`
	Topic   string `json:"topic"`
	NSFW    bool   `json:"nsfw"`
	Bitrate uint8  `json:"bitrate"`
}
