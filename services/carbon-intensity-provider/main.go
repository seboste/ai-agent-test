package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting carbon intensity provider service...")
	// router := mux.NewRouter()
	// log.Fatal(http.ListenAndServe(":8080", router))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
