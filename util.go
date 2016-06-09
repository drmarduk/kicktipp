package main

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	buf := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(buf, 23)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
