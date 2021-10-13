package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.RequestURI[1:]
	http.Redirect(w, r, "https://shiftingphotons.dev", http.StatusTemporaryRedirect)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("shorten"))
	w.WriteHeader(http.StatusOK)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/ping", ping)
	r.HandleFunc("/api/shorten", shorten)
	r.HandleFunc("/{shortUrl:[a-zA-Z0-9]+}", redirect)

	log.Fatal(http.ListenAndServe(":8080", r))
}
