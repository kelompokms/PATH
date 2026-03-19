package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (app *App) getClasses(c *echo.Context) error {
	return c.String(http.StatusOK, "get class")
}

func (app *App) getClass(c *echo.Context) error {
	code := c.Param("code")
	return c.String(http.StatusOK, "get class "+code)
}

func (app *App) putClass(c *echo.Context) error {
	code := c.Param("code")
	postId := c.Param("postId")
	title := c.FormValue("title")
	if title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "berikan judul!")
	}

	description := c.FormValue("description")
	if description == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "berikan deskripsi!")
	}
	return c.String(http.StatusOK, "edit post"+postId+" on class "+code)
}

func (app *App) postClass(c *echo.Context) error {
	code := c.Param("code")
	title := c.FormValue("title")
	if title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "berikan judul!")
	}

	description := c.FormValue("description")
	if description == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "berikan deskripsi!")
	}
	return c.String(http.StatusOK, "new post on class "+code)
}
