package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *App) getTugas(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		http.Error(w, "URL tidak valid!", http.StatusBadRequest)
		return
	}

	posts, err := app.db.ListPostTugas(r.Context(), code)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mendapatkan posts!", http.StatusBadGateway)
		return
	}

	jsonString, err := json.Marshal(posts)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mendapatkan postingan", http.StatusBadGateway)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jsonString)
}
