package handler

import (
	"fmt"
	"log"
	"net/http"
	"path_project/internal/db"
	"path_project/internal/utils"
	"strings"

	"github.com/go-chi/chi/v5"
)

type App struct {
	db *db.Queries
	r  *chi.Mux
}

func NewApp(r *chi.Mux, db *db.Queries) *App {
	app := &App{
		r:  r,
		db: db,
	}
	app.registerRoutes()
	return app
}

func (app *App) Start() {
	port := utils.GetPort()

	println("\nMenampilkan semua routes:")
	if err := chi.Walk(app.r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
	println()

	log.Println("Server berjalan di PORT :" + port)
	http.ListenAndServe(":"+port, app.r)
}

func walkFunc(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	route = strings.ReplaceAll(route, "/*/", "/")
	fmt.Printf("\t%s %s\n", method, route)
	return nil
}
