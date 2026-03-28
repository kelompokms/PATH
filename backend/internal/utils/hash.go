package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func NewHashCode() string {
	randStr := rand.Text()

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s:%s:%d", GetSecret(), randStr, time.Now().UnixNano())))
	return hex.EncodeToString(h.Sum(nil))[:6]
}
