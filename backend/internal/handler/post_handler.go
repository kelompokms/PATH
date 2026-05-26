package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"path_project/internal/db"

	"github.com/go-chi/chi/v5"
)

func (app *App) getPost(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		http.Error(w, "URL tidak valid!", http.StatusBadRequest)
		return
	}

	posts, err := app.db.ListPost(r.Context(), code)
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

func (app *App) createPost(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	nama := r.FormValue("nama")
	deskripsi := r.FormValue("deskripsi")
	tipe := r.FormValue("tipe")

	if code == "" || nama == "" || deskripsi == "" {
		http.Error(w, "URL tidak valid!", http.StatusBadRequest)
		return
	}

	if tipe != "kuis" && tipe != "materi" && tipe != "tugas" {
		http.Error(w, "Tipe materi tidak valid!", http.StatusBadRequest)
		return
	}

	err := app.db.CreatePost(r.Context(), db.CreatePostParams{
		Nama:      nama,
		Deskripsi: deskripsi,
		KodeKelas: code,
		Tipe:      db.TipeMateri(tipe),
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "gagal membuat post", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("post berhasil dibuat!"))
}
