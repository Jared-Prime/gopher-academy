package main

import (
	"log"
	"net/http"
	"os"
)

import "handlers"
import "version"

func main() {
	log.Print("Starting the service...")
	log.Print("commit: ", version.Commit)
	log.Print("build time: ", version.BuildTime)
	log.Print("release: ", version.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No port set!")
	}

	router := handlers.Router()

	log.Print("The service is ready to listen and serve on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
