package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain-text password and returns a hashed version of it.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash verifies if a given plain-text password matches the hashed password.
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
