package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/manudevelopia/app-go/src/internal/service"
)

func UserAll(c echo.Context) error {
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, service.UserAll(email))
}

func UserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, exists := service.UserById(id)
	if exists {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNotFound, "")
}
