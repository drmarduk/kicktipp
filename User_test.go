package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNewUser(t *testing.T) {
	db, _ = sql.Open("sqlite3", "file:data/kicktipp.db")
	tests := []struct {
		name, email, password, out string
	}{
		{"marduk", "mail@knilch.net", "password", ""},
		{"", "mail@knilch.net", "password", "empty username"},
		{"soda", "", "password", "empty mail"},
		{"evilpie", "mail@knilch.net", "", "empty password"},
	}

	for _, tt := range tests {
		u, err := NewUser(tt.name, tt.email, tt.password)
		if u == nil {
			if err.Error() != tt.out {
				t.Errorf("NewUser(%s, %s, %s,): want %s, got %s\n", tt.name, tt.email, tt.password, tt.out, err.Error())
			}
		}
	}
}

func TestSaveUser(t *testing.T) {

}

func TestCheckUser(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"", 0},
		{"marduk", 1},
	}

	for range tests {

	}
}
