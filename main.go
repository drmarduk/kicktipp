package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func main() {
	var err error

	db, err = sql.Open("sqlite3", "file:data/kicktipp.db")
	if err != nil {
		log.Println("main: Error while opening database. " + err.Error())
		return
	}

	http.HandleFunc("/", use(indexHandler, basicAuth))
	http.HandleFunc("/main", use(mainHandler, basicAuth))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	log.Fatal(http.ListenAndServeTLS("", "data/cert.pem", "data/key.pem", nil))

}
