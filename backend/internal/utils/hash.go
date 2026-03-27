package utils

import (
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
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s:%d", GetSecret(), time.Now().UnixNano())))
	return hex.EncodeToString(h.Sum(nil))[:6]
}
