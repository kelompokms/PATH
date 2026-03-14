package handler

import (
	"path_project/internal/db"

	"github.com/labstack/echo/v5"
)

type App struct {
	db *db.Queries
	e  *echo.Echo
}

func NewApp(e *echo.Echo, db *db.Queries) *App {
	return &App{
		e:  e,
		db: db,
	}
}

func (app *App) Start() {
	err := app.e.Start(":3000")
	if err != nil {
		app.e.Logger.Error("Failed to start server: ", err)
	}
}
