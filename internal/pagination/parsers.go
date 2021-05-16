package pagination

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

func ParseCursor(ctx echo.Context) ksuid.KSUID {
	var value string

	if header := ctx.Request().Header.Get(cursorHeader); header != "" {
		value = header
	}

	if query := ctx.QueryParam(cursorQueryParam); query != "" {
		value = query
	}

	parsed, err := ksuid.Parse(value)

	if (value == "" || err != nil) {
		return ksuid.Nil
	}

	return parsed
}

func ParseCount(ctx echo.Context) uint16 {
	var value string

	if header := ctx.Request().Header.Get(countHeader); header != "" {
		value = header
	}

	if query := ctx.QueryParam(countQueryParam); query != "" {
		value = query
	}

	parsed, err := strconv.ParseUint(value, 10, 16)

	if (value == "" || err != nil) {
		return CountDefault
	}

	return uint16(parsed)
}

func ParseReverse(ctx echo.Context) bool {
	var value string

	if header := ctx.Request().Header.Get(reverseHeader); header != "" {
		value = header
	}

	if query := ctx.QueryParam(reverseQueryParam); query != "" {
		value = query
	}

	parsed, err := strconv.ParseBool(value)

	if (value == "" || err != nil) {
		return false
	}

	return parsed
}