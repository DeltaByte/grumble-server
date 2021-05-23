package channelsController

import (
	"net/http"

	"github.com/grumblechat/server/internal/channel"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func listHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res, err := channel.GetAll(db)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
