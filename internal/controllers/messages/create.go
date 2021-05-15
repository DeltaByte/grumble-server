package messagesController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	"gitlab.com/grumblechat/server/pkg/message"
	bolt "go.etcd.io/bbolt"
)

func createHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// parse channelID
		channelID, err := ksuid.Parse(ctx.Param("channelID"))
		if (err != nil) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// validate channel exists and is the correct type
		if err := validateChannel(db, channelID); err != nil {
			return err
		}

		// create message and bind the request body to it
		msg := message.New(channelID)
		if err := ctx.Bind(msg); err != nil {
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