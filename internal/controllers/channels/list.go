package channels

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/grumblechat/server/pkg/channel"
	bolt "go.etcd.io/bbolt"
)

func listHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var channels []channel.Channel

		err := db.View(func(tx *bolt.Tx) (error) {
			dbb := tx.Bucket([]byte(channel.DBBucket))

			// iterate over all channels, decode and add to array
			dbb.ForEach(func(k, v []byte) error {
				decoded, err := channel.Decode(v)
				if (err != nil) { return err }
				channels = append(channels, decoded)
				return nil
			})

			return nil
		})

		if (err != nil) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, channels)
	}
}