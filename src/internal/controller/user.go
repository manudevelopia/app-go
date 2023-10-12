package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/manudevelopia/app-go/src/internal/service"
	"net/http"
)

func UserAll(c echo.Context) error {
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, service.UserAll(email))
}

func UserById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, service.UserById(id))
}
