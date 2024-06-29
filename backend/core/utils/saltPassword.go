package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func saltPassword(password []byte, cost int) ([]byte, error) {
	// Generate a random salt with desired cost (higher cost = more secure, but slower)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
