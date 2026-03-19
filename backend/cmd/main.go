package main

import (
	"context"
	"log"
	"path_project/internal/db"
	"path_project/internal/handler"
	"path_project/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	pool, err := pgxpool.New(context.Background(), utils.GetDatabaseUrl())
	if err != nil {
		log.Fatal("Failed to create database pool:", err.Error())
	}

	queries := db.New(pool)

	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(utils.RequestLoggerConfig))

	app := handler.NewApp(e, queries)
	app.Start()
}
