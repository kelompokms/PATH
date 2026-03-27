package handler

import (
	"context"
	"log"
	"net/http"
	"net/mail"
	"path_project/internal/db"
	"path_project/internal/utils"
	"regexp"
)

func (app *App) register(w http.ResponseWriter, r *http.Request) {
	numberRegex := regexp.MustCompile(`^[0-9]+$`)

	nama := r.FormValue("nama")
	email := r.FormValue("email")
	telepon := r.FormValue("telepon")
	password := r.FormValue("password")

	if nama == "" || email == "" || telepon == "" || password == "" {
		http.Error(w, "Input tidak lengkap!", http.StatusBadRequest)
		return
	}

	if len(nama) >= 128 || len(email) >= 255 || len(telepon) >= 32 || len(password) >= 255 {
		http.Error(w, "Input tidak valid!", http.StatusBadRequest)
		return
	}

	if numberRegex.MatchString(telepon) == false {
		http.Error(w, "telepon tidak berupa angka!", http.StatusBadRequest)
		return
	}

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

	err = app.db.CreatePengguna(context.Background(), db.CreatePenggunaParams{
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

	w.Write([]byte("Sukses mendaftarkan pengguna"))
}

func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
}

func (app *App) putUser(w http.ResponseWriter, r *http.Request) {
}
