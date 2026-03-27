package utils

import (
	"log"
	"os"
)

func GetDatabaseUrl() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL tidak boleh kosong!")
	}
	return url
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Println("PORT kosong. Default ke 3000")
	}
	return port
}

func GetSecret() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET tidak boleh kosong!")
	}
	return secret
}

func GetDevMode() bool {
	devMode := os.Getenv("DEV")
	switch devMode {
	case "0":
		return false
	case "1":
		return true
	default:
		log.Println("DEV kosong. Default ke true")
		return true
	}
}
