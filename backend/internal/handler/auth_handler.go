package handler

import (
	"log"
	"net/http"
	"net/mail"
	"path_project/internal/utils"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// validasi standar input
	if email == "" || password == "" {
		http.Error(w, "input tidak lengkap!", http.StatusBadRequest)
		return
	}

	if len(email) > 255 || len(password) > 255 {
		http.Error(w, "input tidak valid!", http.StatusBadRequest)
		return
	}

	// validasi email
	_, err := mail.ParseAddress(email)
	if err != nil {
		http.Error(w, "email tidak valid!", http.StatusBadRequest)
		return
	}

	res, err := app.db.CheckPengguna(r.Context(), email)
	if err != nil {
		log.Println(err)
		http.Error(w, "email atau password salah!", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil {
		log.Println(err)
		http.Error(w, "gagal membuat kredensial", http.StatusBadRequest)
		return
	}

	_, tokenString, _ := app.tokenAuth.Encode(map[string]interface{}{
		"user_id": res.ID,
		"exp":     time.Now().Add(time.Hour * 168).Unix(),
	})

	utils.SetCookie(w, tokenString)

	w.Write([]byte("Berhasil login!"))
}

func (app *App) logout(w http.ResponseWriter, r *http.Request) {
	utils.ClearCookie(w)
	w.Write([]byte("Berhasil logout!"))
}

func (app *App) getAuth(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "Gagal mendapatkan token: "+err.Error(), http.StatusBadRequest)
	}

	userId := int32(claims["user_id"].(float64))

	_, err = app.db.GetPengguna(r.Context(), userId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mendapatkan pengguna", http.StatusBadRequest)
		return
	}

	w.Write([]byte("User terautentikasi!"))
}
