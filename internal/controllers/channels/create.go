package channelsController

import (
	"net/http"

	"github.com/grumblechat/server/internal/channel"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func createHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var newChannel channel.Channel

		// bind request to Data Transfer Object
		dto := &channelDTO{}
		if err := ctx.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// run DTO validations
		if err := ctx.Validate(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// voice channel
		if dto.Type == "voice" {
			newChannel = channel.NewVoice()
			if err := copier.CopyWithOption(newChannel, dto, copierOptions); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		// text channel
		if dto.Type == "text" {
			newChannel = channel.NewText()
			if err := copier.CopyWithOption(newChannel, dto, copierOptions); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		// validate channel
		if err := ctx.Validate(newChannel); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// persist channel
		if err := newChannel.Save(db); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, newChannel)
	}
}
