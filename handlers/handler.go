package handlers

import (
	"dms_mixer/components"

	"github.com/labstack/echo/v4"
)


func Root(c echo.Context) error {

	return components.Page().Render(c.Request().Context(), c.Response().Writer) 
}

func RenderWindow(c echo.Context) error {
	return components.DivCon("whatever").Render(c.Request().Context(), c.Response().Writer)
}

func RenderMusic(c echo.Context) error {
	return components.Music().Render(c.Request().Context(), c.Response().Writer)
}