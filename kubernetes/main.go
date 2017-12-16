package main

import (
	"log"
	"net/http"
)

import "./handlers"

func main() {
	log.Print("Starting the service...")

	router := handlers.Router()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
