package main

import (
	"log"
	"net/http"
	"spyfall-server/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting payload channel listener")
	go handlers.ListenToPayloadChannel()

	log.Println("Starting server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}