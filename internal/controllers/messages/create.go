package messagesController

import (
	"net/http"

	"github.com/grumblechat/server/internal/message"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func createHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// validate channel exists and is the correct type
		chn, err := getChannel(db, ctx.Param("channelID"))
		if err != nil {
			return err
		}

		// bind request to Data Transfer Object
		dto := &messageDTO{}
		if err := ctx.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// copy fields from DTO to actual message
		msg := message.New(chn.ID)
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
