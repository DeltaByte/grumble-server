package channelsController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/grumblechat/server/pkg/channels"
	bolt "go.etcd.io/bbolt"
)

func listHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var res []channels.Channel

		err := db.View(func(tx *bolt.Tx) (error) {
			dbb := tx.Bucket([]byte(channels.DBBucket))

			// iterate over all channels, decode, and add to result
			dbb.ForEach(func(k, v []byte) error {
				decoded, err := channels.Decode(v)
				if (err != nil) { return err }
				res = append(res, decoded)
				return nil
			})

			return nil
		})

		if (err != nil) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, res)
	}
}