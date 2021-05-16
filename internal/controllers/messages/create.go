package messagesController

import (
	"net/http"

	"github.com/grumblechat/server/pkg/message"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

type createHandlerDTO struct {
	ChannelID ksuid.KSUID `json:"channel_id"`
	Body      string      `json:"body" validate:"min=1,max=2048,required"`
	TTL       uint32      `json:"ttl"`
}

func createHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// parse channelID
		channelID, err := ksuid.Parse(ctx.Param("channelID"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// validate channel exists and is the correct type
		if err := validateChannel(db, channelID); err != nil {
			return err
		}

		// bind request to Data Transfer Object
		dto := &createHandlerDTO{}
		if err := ctx.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// copy fields from DTO to actual message
		msg := message.New(channelID)
		if err := copier.CopyWithOption(msg, dto, copierOptions); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// validate message
		if err := ctx.Validate(msg); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// persist message
		if err := msg.Save(db); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, msg)
	}
}
