package main

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPasswort(t *testing.T) {
	tests := []struct {
		in string
	}{
		{"Test"},
	}

	for _, tt := range tests {
		hash, err := HashPassword(tt.in)
		if err != nil {
			t.Error(err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(tt.in))
		if err != nil {
			t.Errorf("HashPassword in lib")
		}

	}

}
