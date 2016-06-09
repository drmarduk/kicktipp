package main

import (
	"errors"
	"log"

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

func validateUser(userid int) (bool, error) {
	// TODO: check if userid matches token in db
	return false, errors.New("Not implemented")
}

func checkExistingUser(name string) bool {
	//db, err := sql.Open("sqlite3", "file:data/kicktipp.db")
	//if err != nil {
	//		t.Fatal(err)
	//}
	stmt, err := db.Prepare("SELECT count(*) FROM user WHERE name = ?;")
	if err != nil {
		log.Println("checkExistingUser: Error while preparing statement. " + err.Error())
		return false
	}
	defer stmt.Close()

	result := stmt.QueryRow(name)
	var count int = 0
	if err := result.Scan(&count); err != nil {
		log.Println("checkExistingUser: Error while scanning row. " + err.Error())
		return false
	}
	if count > 1 {
		return false
	}
	return true
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
	if checkExistingUser(name) {
		log.Printf("NewUser: User %s already exists.\n", name)
		return nil, errors.New("user already exists")
	}
	return &User{Name: name, Email: email, Password: password}, nil
}
