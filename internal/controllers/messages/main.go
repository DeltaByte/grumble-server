package messagesController

import (
	"net/http"

	"github.com/grumblechat/server/pkg/channel"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

var copierOptions = copier.Option{
	IgnoreEmpty: true,
}

func BindRoutes(db *bolt.DB, router *echo.Group) {
	router.GET("", listHandler(db))
	router.POST("", createHandler(db))
	router.PATCH("/:id", updateHandler(db))
	router.DELETE("/:id", deleteHandler(db))
}

type messageDTO struct {
	ChannelID ksuid.KSUID `json:"channel_id"`
	Body      string      `json:"body"`
	TTL       uint32      `json:"ttl"`
}

func getChannel(db *bolt.DB, channelID string) (*channel.TextChannel, *echo.HTTPError) {
	id, err := ksuid.Parse(channelID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	chn, err := channel.Find(db, id)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if chn == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Channel ID not recognized")
	}

	if chn.GetType() != "text" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Messages are only available for 'text' channel types")
	}

	textChannel := chn.(*channel.TextChannel)
	return textChannel, nil
}
