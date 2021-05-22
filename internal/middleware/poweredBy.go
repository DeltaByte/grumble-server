package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

const versionHeader = "X-Powered-By"

func PoweredBy(version string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if version == "" {
				version = "unkown"
			}

			server := fmt.Sprintf("GrumbleChat/%s", string(version))
			ctx.Response().Header().Set(versionHeader, server)
			return next(ctx)
		}
	}
}