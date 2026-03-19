package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (app *App) getUser(c *echo.Context) error {
	return c.String(http.StatusOK, "get user")
}

func (app *App) putUser(c *echo.Context) error {
	return c.String(http.StatusOK, "put user")
}
