package main

import (
	"github.com/labstack/echo/v4"
	"github.com/manudevelopia/app-go/src/internal/controller"
)

func main() {
	e := echo.New()
	e.Static("/", "src/static/swagger")
	e.GET("/v1/users", controller.UserAll)
	e.GET("/v1/users/:id", controller.UserById)
	e.Logger.Fatal(e.Start(":1323"))
}
