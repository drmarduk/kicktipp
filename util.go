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

func ComparePassword(plaintext, hash string) (bool, error) {
	_hash := []byte(hash)
	_plain := []byte(plaintext)

	err := bcrypt.CompareHashAndPassword(_hash, _plain)
	if err != nil {
		return false, err
	}
	return true, nil
}
