package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCheckMatch(t *testing.T) {
	db, _ = sql.Open("sqlite3", "file:data/kicktipp.db")
	tests := []struct {
		id  int
		out bool
	}{
		{1, true},
		{-1, false},
	}

	for _, tt := range tests {
		x, err := checkMatch(tt.id)
		if err != nil {
			t.Error("Error while dingsi")
		}
		if x != tt.out {
			t.Error("CheckMatch(%d): expected %t, got %t\n", tt.id, tt.out, x)
		}

	}
}
