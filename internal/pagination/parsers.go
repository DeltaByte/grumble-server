package pagination

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

func ParseCursor(ctx echo.Context) (ksuid.KSUID, error) {
	var value string

	if header := ctx.Request().Header.Get(cursorHeader); header != "" {
		value = header
	}

	if query := ctx.QueryParam(cursorQueryParam); query != "" {
		value = query
	}

	if (value == "") {
		return ksuid.Nil, nil
	}

	return ksuid.Parse(value)
}

func ParseCount(ctx echo.Context) (uint16, error) {
	var value string

	if header := ctx.Request().Header.Get(countHeader); header != "" {
		value = header
	}

	if query := ctx.QueryParam(countQueryParam); query != "" {
		value = query
	}

	if (value == "") {
		return CountDefault, nil
	}

	parsed, err := strconv.ParseUint(value, 10, 16)
	if (err != nil) { return 0, err }
	return uint16(parsed), nil
}