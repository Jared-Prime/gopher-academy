package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

import "handlers"

func main() {
	log.Print("Starting the service...")
	log.Print("commit: ", handlers.Commit)
	log.Print("build time: ", handlers.BuildTime)
	log.Print("release: ", handlers.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No port set!")
	}

	router := handlers.Router()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr: ":"+port,
		Handler: router,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve on port ", port)

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM")
	}

	log.Print("The service is shutting down...")
	server.Shutdown(context.Background())
	log.Print("done!")
}
