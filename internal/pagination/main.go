package pagination

import (
	"bytes"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	bolt "go.etcd.io/bbolt"
)

const (
	cursorHeader = "X-Pagination-Cursor"
	cursorQueryParam = "cursor"
	countHeader  = "X-Pagination-Count"
	countQueryParam = "count"
	CountDefault = 50
)

func New(ctx echo.Context) (*Pagination, error) {
	cursor, err := ParseCursor(ctx)
	if (err != nil) { return nil, err }

	count, err := ParseCount(ctx)
	if (err != nil) { return nil, err }

	return &Pagination{
		Cursor: cursor,
		Count: count,
	}, nil
}

func SetHeaders(ctx echo.Context, pgn *Pagination) {
	ctx.Response().Header().Set(cursorHeader, pgn.Cursor.String())
	ctx.Response().Header().Set(countHeader, strconv.FormatUint(uint64(pgn.Count), 10))
}

type Pagination struct {
	Cursor ksuid.KSUID `query:"cursor"`
	Count  uint16      `query:"count" validate:"min=1,max=1000"`
}

func (pgn *Pagination) InitCursor(cursor *bolt.Cursor) (key []byte, value []byte) {
	if (!pgn.Cursor.IsNil()) {
		// seek to the specified key, Bolt automatically goes to the next if it isn't found
		seekKey := pgn.Cursor.Bytes()
		k, v := cursor.Seek(seekKey)

		// manually go to the next key if the found one is the same as the pagination cursor
		if (bytes.Equal(seekKey, k)) {
			k, v = cursor.Next()
		}

		return k, v
	}

	log.Print("using first")
	return cursor.First()
}
