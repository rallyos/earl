package handlers

import (
	"earl/shorten"
	"encoding/json"
	"fmt"
	"net/http"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	response := json.NewEncoder(w)
	// TODO: Is URL validation?

	urlCarrier := struct {
		Url string `json:"url"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&urlCarrier)
	if err != nil {
		urlCarrier.Url = err.Error()
		response.Encode(urlCarrier)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shortUrl, err := shorten.Shorten(urlCarrier.Url)
	if err != nil {
		urlCarrier.Url = err.Error()
		response.Encode(urlCarrier)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.TLS == nil {
		r.URL.Scheme = "http"
	} else {
		r.URL.Scheme = "https"
	}
	urlCarrier.Url = fmt.Sprintf("%s://%s/%s", r.URL.Scheme, r.Host, shortUrl)

	w.WriteHeader(http.StatusCreated)
	response.Encode(urlCarrier)
}
