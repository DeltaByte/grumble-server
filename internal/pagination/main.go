package pagination

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

const (
	cursorHeader = "X-Pagination-Cursor"
	cursorQueryParam = "cursor"
	countHeader  = "X-Pagination-Count"
	countQueryParam = "count"
	CountDefault = 50
	reverseHeader = "X-Pagination-Reverse"
	reverseQueryParam = "reverse"
	moreHeader = "X-Pagination-More"
)

func New(ctx echo.Context) *Pagination {
	return &Pagination{
		Cursor: ParseCursor(ctx),
		NextCursor: ksuid.Nil,
		Count: ParseCount(ctx),
		Reverse: ParseReverse(ctx),
	}
}

func SetHeaders(ctx echo.Context, pgn *Pagination) {
	hasMore := !pgn.NextCursor.IsNil()
	ctx.Response().Header().Set(moreHeader, strconv.FormatBool(hasMore))
	if (hasMore) {
		ctx.Response().Header().Set(cursorHeader, pgn.NextCursor.String())
	}

	ctx.Response().Header().Set(countHeader, strconv.FormatUint(uint64(pgn.Count), 10))
}
