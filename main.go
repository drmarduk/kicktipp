package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("", nil))

}
