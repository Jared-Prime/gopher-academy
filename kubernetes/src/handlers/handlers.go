package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"sync/atomic"
	"time"
)

func Router() *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("readiness set to false by default")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("readiness set to true after 10 seconds")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/health", health)
	r.HandleFunc("/ready", ready(isReady))
	return r
}
