package messagesController

import (
	"net/http"

	"github.com/grumblechat/server/pkg/message"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func updateHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// bind request to Data Transfer Object
		dto := &messageDTO{}
		if err := ctx.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// validate channel exists and is the correct type
		chn, chnErr := getChannel(db, ctx.Param("channelID"))
		if chnErr != nil {
			return chnErr
		}

		// parse messageID
		id, err := ksuid.Parse(ctx.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// get message
		msg, err := message.Find(db, chn.ID, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if msg == nil {
			return echo.NewHTTPError(http.StatusNotFound, "Channel ID not recognized.")
		}

		// copy fields from DTO to actual message
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

		return ctx.JSON(http.StatusOK, msg)
	}
}
