package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Println("tap-go is Ready to receive requests")
	log.Fatal(http.ListenAndServe(":8080", router))
}
