package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func (app *App) registerRoutes() {
	app.r.Get("/", app.getIndex)
	app.r.Post("/login", app.login)
	app.r.Post("/logout", app.logout)
	app.r.Post("/register", app.register)

	app.r.Group(func(rp chi.Router) {
		rp.Use(jwtauth.Verifier(app.tokenAuth))
		rp.Use(jwtauth.Authenticator(app.tokenAuth))

		rp.Get("/user", app.getUser)
		rp.Patch("/user", app.patchUser)
		rp.Get("/class", app.getClasses)
		rp.Post("/class", app.createClass)
		rp.Get("/class/{code}", app.getClass)
		rp.Put("/class/{code}/{postId}", app.putClass)
	})

}

func (app *App) getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello PATH!"))
}
