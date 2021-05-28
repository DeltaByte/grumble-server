package channelsController

import (
	"net/http"

	"github.com/grumblechat/server/internal/entities/channel"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func deleteHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// parse ID
		id, err := ksuid.Parse(ctx.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// get channel
		chn, err := channel.Find(db, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if chn == nil {
			return echo.NewHTTPError(http.StatusNotFound, "Channel ID not recognized.")
		}

		// delete channel
		if err := chn.Delete(db); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, "Channel deleted.")
	}
}
