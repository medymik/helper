package helper

import (
	"golang.org/x/crypto/bcrypt"
)

const COST = 10

func HashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), COST)
	return string(bytes), err
}

func CheckPassword(plainPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
	return err == nil
}
