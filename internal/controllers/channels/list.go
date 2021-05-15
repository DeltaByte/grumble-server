package channelsController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/grumblechat/server/pkg/channel"
	bolt "go.etcd.io/bbolt"
)

func listHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res, err := channel.GetAll(db)

		if (err != nil) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, res)
	}
}