package handler

import "net/http"

func (app *App) registerRoutes() {
	app.r.Get("/", app.getIndex)

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
