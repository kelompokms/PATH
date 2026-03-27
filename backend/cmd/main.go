package main

import (
	"context"
	"log"
	"path_project/internal/db"
	"path_project/internal/handler"
	"path_project/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := pgxpool.New(context.Background(), utils.GetDatabaseUrl())
	if err != nil {
		log.Fatal("Failed to initialize database pool:", err.Error())
	}
	queries := db.New(pool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	app := handler.NewApp(r, queries)
	app.Start()
}
