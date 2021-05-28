package messagesController

import (
	"net/http"

	"github.com/deltabyte/grumble-server/internal/entities/message"
	"github.com/deltabyte/grumble-server/internal/pagination"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

func listHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pgn := pagination.New(ctx)

		// parse channelID
		channelID, err := ksuid.Parse(ctx.Param("channelID"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// validate pagination
		if err := ctx.Validate(pgn); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// load messages from DB
		res, err := message.GetAll(db, channelID, *pgn)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// response
		pagination.SetHeaders(ctx, pgn)
		return ctx.JSON(http.StatusOK, res)
	}
}
