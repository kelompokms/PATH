package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (app *App) registerRoutes() {
	app.e.GET("/", app.getIndex)

	app.e.GET("/user", app.getUser)
	app.e.PATCH("/user", app.putUser)
	app.e.GET("/class", app.getClasses)
	app.e.GET("/class/:code", app.getClass)
	app.e.POST("/class/:code", app.postClass)
	app.e.PUT("/class/:code/:postId", app.putClass)
	app.e.POST("/login", app.login)
	app.e.POST("/logout", app.logout)
}

func (app *App) getIndex(c *echo.Context) error {
	return c.String(http.StatusOK, "get index")
}
