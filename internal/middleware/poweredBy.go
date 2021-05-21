package middleware

import (
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

const versionHeader = "X-Powered-By"

func PoweredBy() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			version, err := ioutil.ReadFile("VERSION")
			if err != nil {
				return err
			}

			next(ctx)
			server := fmt.Sprintf("GrumbleChat/%s", string(version))
			ctx.Response().Header().Set(versionHeader, server)
			return nil
		}
	}
}