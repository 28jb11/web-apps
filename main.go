package main

import (
	"log"
	"net/http"

	"github.com/28jb11/calculator/routes"
)

func main() {
	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/greeting", routes.NameHandler)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
