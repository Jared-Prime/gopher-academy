package main

import (
	"log"
	"net/http"
	"os"
)

import "./handlers"

func main() {
	log.Print("Starting the service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No port set!")
	}

	router := handlers.Router()

	log.Print("The service is ready to listen and serve on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
