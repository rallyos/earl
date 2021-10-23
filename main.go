package main

import (
	"earl/app/handlers"
	"earl/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/api/ping", handlers.Ping).Methods(http.MethodGet)
	r.HandleFunc("/api/shorten", handlers.Shorten).Methods(http.MethodPost)
	r.HandleFunc("/{shortUrl:[a-zA-Z0-9]+}", handlers.Redirect).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}
