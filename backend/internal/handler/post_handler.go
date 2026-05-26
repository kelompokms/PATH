package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"path_project/internal/db"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
	tenggat := r.FormValue("tenggat")

	if code == "" || nama == "" || deskripsi == "" {
		http.Error(w, "URL tidak valid!", http.StatusBadRequest)
		return
	}

	if tipe != "kuis" && tipe != "materi" && tipe != "tugas" {
		http.Error(w, "Tipe materi tidak valid!", http.StatusBadRequest)
		return
	}

	newPost := db.CreatePostParams{
		Nama:      nama,
		Deskripsi: deskripsi,
		KodeKelas: code,
		Tenggat:   pgtype.Timestamp{Valid: false},
		Tipe:      db.TipeMateri(tipe),
	}

	if tipe == "tugas" {
		if tenggat == "" {
			http.Error(w, "Tenggat harus diisi untuk tugas!", http.StatusBadRequest)
			return
		}

		// TODO: sesuaikan layout dengan format javascript
		parsedTenggat, err := time.Parse("", tenggat)
		if err != nil {
			http.Error(w, "Gagal mendapatkan tenggat", http.StatusBadRequest)
			return
		}

		newPost.Tenggat = pgtype.Timestamp{
			Time:  parsedTenggat,
			Valid: true,
		}
	}

	err := app.db.CreatePost(r.Context(), newPost)

	if err != nil {
		log.Println(err)
		http.Error(w, "gagal membuat post", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("post berhasil dibuat!"))
}

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
