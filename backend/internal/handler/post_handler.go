package handler

import (
	"log"
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

	var TipeMateri db.TipeMateri

	switch tipe {
	case "kuis":
		TipeMateri = db.TipeMateriKuis
	case "materi":
		TipeMateri = db.TipeMateriMateri
	case "tugas":
		TipeMateri = db.TipeMateriTugas
	default:
		http.Error(w, "Tipe materi tidak valid!", http.StatusBadRequest)
		return
	}

	err := app.db.CreatePost(r.Context(), db.CreatePostParams{
		Nama:      nama,
		Deskripsi: deskripsi,
		KodeKelas: code,
		Tipe:      TipeMateri,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "gagal membuat post", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tipe))
}
