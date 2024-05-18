package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	salt := os.Getenv("BCRYPT_SALT")
	saltInt, _ := strconv.Atoi(salt)

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), int(saltInt))
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
