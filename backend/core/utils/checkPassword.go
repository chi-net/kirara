package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
