package messagesController

import (
	"net/http"

	"github.com/grumblechat/server/pkg/message"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func deleteHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
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
			return echo.NewHTTPError(http.StatusNotFound, "Message ID not recognized.")
		}

		// delete message
		if err := msg.Delete(db); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, "Message deleted.")
	}
}
