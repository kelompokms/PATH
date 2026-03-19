package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (app *App) login(c *echo.Context) error {
	username := c.FormValue("username")
	if username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username tidak boleh kosong!")
	}

	password := c.FormValue("password")
	if password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "password tidak boleh kosong!")
	}

	return c.String(http.StatusOK, "login dengan username="+username+", password="+password)
}

func (app *App) logout(c *echo.Context) error {
	return c.String(http.StatusOK, "logout")
}
