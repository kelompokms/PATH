package handler

import (
	"net/http"
	"path_project/internal/db"

	"github.com/go-chi/chi/v5"
)

func (app *App) createPost(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	nama := r.FormValue("nama")
	deskripsi := r.FormValue("deskripsi")
	tipe := r.FormValue("tipe")

	if code == "" || nama == "" || deskripsi == "" {
		http.Error(w, "URL tidak valid!", http.StatusBadRequest)
		return
	}

	if tipe != "kuis" || tipe != "materi" || tipe != "tugas" {
		http.Error(w, "tipe tidak valid!", http.StatusBadRequest)
		return
	}

	app.db.CreatePost(r.Context(), db.CreatePostParams{
		Nama:      nama,
		Deskripsi: deskripsi,
		KodeKelas: code,
		Tipe:      db.TipeMateriKuis,
	})

	w.Write([]byte("GET"))
}
