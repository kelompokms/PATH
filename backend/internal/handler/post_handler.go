package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"path_project/internal/db"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (app *App) getPost(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal membaca ID post", http.StatusBadRequest)
		return
	}

	post, err := app.db.GetPost(r.Context(), db.GetPostParams{
		KodeKelas: code,
		ID:        int32(postID),
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mendapatkan post", http.StatusBadRequest)
		return
	}

	jsonString, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal memproses respon", http.StatusInternalServerError)
		return
	}

	w.Write(jsonString)

}

func (app *App) getPosts(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

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
		parsedTenggat, err := time.Parse("2006-01-02T15:04:05.000Z07:00", tenggat)
		if err != nil {
			log.Println(err)
			http.Error(w, "Tenggat tugas tidak valid!", http.StatusBadRequest)
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
