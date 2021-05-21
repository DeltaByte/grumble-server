package channelsController

import (
	"net/http"

	"github.com/grumblechat/server/pkg/channel"
	"github.com/segmentio/ksuid"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func updateHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// parse ID
		id, err := ksuid.Parse(ctx.Param("id"))
		if (err != nil) {
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

		// bind request to Data Transfer Object
		dto := &channelDTO{}
		if err := ctx.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// prevent type switching,
		// this is already caught by validation but I want a specific error for it
		if dto.Type != chn.GetType() {
			return echo.NewHTTPError(http.StatusBadRequest, "Channel types are immutable.")
		}

		// copy data from DTO
		if err := copier.CopyWithOption(chn, dto, copierOptions); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// validate channel
		if err := ctx.Validate(chn); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// persist channel
		if err := chn.Save(db); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, chn)
	}
}
