package main

import (
	"dms_mixer/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return handlers.Hello.HandleHelloShow(handlers.Hello{}, c)
	})
	e.Static("/assets", "static")
	e.Logger.Fatal(e.Start(":1323"))
}
