package main

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	buf := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(buf, 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
