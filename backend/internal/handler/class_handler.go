package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"path_project/internal/db"
	"path_project/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (app *App) getClasses(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	user_id := int32(claims["user_id"].(float64))

	res, err := app.db.ListKelas(r.Context(), user_id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mengambil informasi kelas", http.StatusInternalServerError)
		return
	}
	if len(res) <= 0 {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	jsonString, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal membuat respon", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jsonString)
}

func (app *App) getClass(w http.ResponseWriter, r *http.Request) {
	kode := chi.URLParam(r, "code")
	res, err := app.db.GetKelas(r.Context(), kode)
	if err != nil {
		log.Println(err)
		http.Error(w, "kelas tidak ditemukan", http.StatusNotFound)
		return
	}

	jsonString, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, "gagal membuat respon", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jsonString)
}

func (app *App) createClass(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := int32(claims["user_id"].(float64))

	log.Println("user ID:", userID)

	nama := r.FormValue("nama")
	bagian := r.FormValue("bagian")
	deskripsi := r.FormValue("deskripsi")

	if nama == "" {
		http.Error(w, "input tidak lengkap!", http.StatusBadRequest)
		return
	}

	if len(nama) > 64 || len(bagian) > 320 {
		http.Error(w, "nama kelas atau bagian terlalu panjang!", http.StatusBadRequest)
		return
	}

	if len(nama) < 3 {
		http.Error(w, "nama kelas atau subjek terlalu pendek!", http.StatusBadRequest)
		return
	}

	validUserID, err := app.db.ValidatePenggunaID(r.Context(), userID)
	if err != nil {
		http.Error(w, "gagal memvalidasi pengguna", http.StatusInternalServerError)
		return
	}

	kode := utils.NewHashCode()
	res, err := app.db.GetKelasCode(r.Context(), kode)
	if err != nil {
		log.Println(err.Error())
	}

	// TODO: will refactor
	for len(res) >= 1 {
		kode = utils.NewHashCode()
		res, err = app.db.GetKelasCode(r.Context(), kode)
		if err != nil {
			log.Println(err.Error())
		}
	}

	err = app.db.CreateKelas(r.Context(), db.CreateKelasParams{
		Nama:      nama,
		Bagian:    pgtype.Text{String: bagian, Valid: true},
		Deskripsi: pgtype.Text{String: deskripsi, Valid: true},
		Pengajar:  validUserID,
		Kode:      kode,
	})

	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "gagal membuat kelas: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Write([]byte("berhasil membuat kelas!"))
}

func (app *App) putClass(w http.ResponseWriter, r *http.Request) {
}
