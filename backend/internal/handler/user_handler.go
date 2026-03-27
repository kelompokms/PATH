package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"path_project/internal/db"
	"path_project/internal/utils"
	"regexp"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type tokenStruct struct {
	UserID int   `json:"user_id"`
	Exp    int64 `json:"exp"`
}

func (app *App) register(w http.ResponseWriter, r *http.Request) {
	numberRegex := regexp.MustCompile(`^[0-9]+$`)

	nama := r.FormValue("nama")
	email := r.FormValue("email")
	telepon := r.FormValue("telepon")
	password := r.FormValue("password")

	// validasi standar input
	if nama == "" || email == "" || telepon == "" || password == "" {
		http.Error(w, "Input tidak lengkap!", http.StatusBadRequest)
		return
	}

	if len(nama) >= 128 || len(email) >= 255 || len(telepon) >= 32 || len(password) >= 255 {
		http.Error(w, "Input tidak valid!", http.StatusBadRequest)
		return
	}

	// validasi nomer telepon
	if numberRegex.MatchString(telepon) == false {
		http.Error(w, "telepon tidak berupa angka!", http.StatusBadRequest)
		return
	}

	// validasi email
	_, err := mail.ParseAddress(email)
	if err != nil {
		http.Error(w, "email tidak valid!", http.StatusBadRequest)
		return
	}

	hashed_password, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// insert ke database
	userId, err := app.db.CreatePengguna(r.Context(), db.CreatePenggunaParams{
		Nama:     nama,
		Email:    email,
		Telepon:  telepon,
		Password: hashed_password,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "email sudah terdaftarkan!", http.StatusBadRequest)
		return
	}

	_, tokenString, _ := app.tokenAuth.Encode(map[string]interface{}{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 168).Unix(),
	})

	utils.SetCookie(w, tokenString)

	w.Write([]byte("Sukses mendaftarkan pengguna"))
}

func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "gagal parsing token", http.StatusInternalServerError)
		return
	}

	userId := int32(claims["user_id"].(float64))
	res, err := app.db.GetPengguna(r.Context(), userId)
	if err != nil {
		http.Error(w, "gagal mengambil data pengguna", http.StatusInternalServerError)
		return
	}

	jsonString, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "gagal menampilkan hasil", http.StatusInternalServerError)
		return
	}
	w.Write(jsonString)
}

func (app *App) putUser(w http.ResponseWriter, r *http.Request) {
}
