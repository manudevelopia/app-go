package main

import (
	"github.com/manudevelopia/app-go/src/internal/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", greeting)
	e.GET("/users", controller.UserAll)
	e.GET("/users/:id", controller.UserById)
	e.Static("/doc", "cmd/api/swagger")
	e.Logger.Fatal(e.Start(":1323"))
}

func greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
