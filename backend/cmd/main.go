package main

import (
	"context"
	"log"
	"path_project/internal/db"
	"path_project/internal/handler"
	"path_project/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// membuat pool koneksi database
	pool, err := pgxpool.New(context.Background(), utils.GetDatabaseUrl())
	if err != nil {
		log.Fatal("Failed to initialize database pool:", err.Error())
	}
	queries := db.New(pool)

	// membuat jwt token
	tokenAuth := jwtauth.New("HS256", []byte(utils.GetSecret()), nil)

	// membuat router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	app := handler.NewApp(r, queries, tokenAuth)
	app.Start()
}
