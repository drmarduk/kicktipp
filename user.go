package main

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string

	Predictions []Prediction // Anzahl an getippten Spielen, bzw die Tipps
	Points      int
}

func checkExistingUser(name string) (bool, error) {
	stmt, err := db.Prepare("SELECT count(*) FROM user WHERE name = ?;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	result := stmt.QueryRow(name)
	var count int = 0
	if err := result.Scan(&count); err != nil {
		return false, err
	}
	if count > 1 {
		return false, nil
	}
	return true, nil
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("empty username")
	}
	if email == "" {
		return nil, errors.New("empty email")
	}
	if password == "" {
		return nil, errors.New("empty password")
	}
	if chk, err := checkExistingUser(name); !chk {
		if err != nil { // error while db connection foo
			return nil, err
		}
		return nil, errors.New("user already exists")
	}
	var err error
	password, err = HashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{Name: name, Email: email, Password: password}, nil
}
