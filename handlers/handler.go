package handlers

import (
	"dms_mixer/components"

	"github.com/labstack/echo/v4"
)

type Hello struct {
}

func (h Hello) HandleHelloShow(c echo.Context) error {

	return components.Page().Render(c.Request().Context(), c.Response().Writer) 
}

