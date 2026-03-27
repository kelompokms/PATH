package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func (app *App) registerRoutes() {

	app.r.Group(func(rr chi.Router) {
		rr.Use(jwtauth.Verifier(app.tokenAuth))
		rr.Use(jwtauth.Authenticator(app.tokenAuth))
		rr.Get("/", app.getIndex)
	})

	app.r.Get("/user", app.getUser)
	app.r.Patch("/user", app.putUser)
	app.r.Get("/class", app.getClasses)
	app.r.Get("/class/:code", app.getClass)
	app.r.Post("/class/:code", app.postClass)
	app.r.Put("/class/:code/:postId", app.putClass)
	app.r.Post("/login", app.login)
	app.r.Post("/logout", app.logout)
	app.r.Post("/register", app.register)
}

func (app *App) getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello PATH!"))
}
