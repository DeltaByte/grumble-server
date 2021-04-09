package controllers

import "github.com/labstack/echo/v4"

func BindChannelRoutes(routeGroup *echo.Group) {
	routeGroup.GET("/", listHandler)
}

func listHandler(ctx echo.Context) error {
	//
}