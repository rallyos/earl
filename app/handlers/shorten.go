package handlers

import (
	"earl/app/shorten"
	"earl/app/validate"
	"encoding/json"
	"fmt"
	"net/http"
)

type urlWrapper struct {
	Url string
}

//
func Shorten(w http.ResponseWriter, r *http.Request) {
	response := json.NewEncoder(w)
	var urlWrap urlWrapper

	if err := json.NewDecoder(r.Body).Decode(&urlWrap); err != nil {
		halt(&urlWrap, response, w, err)
		return
	}

	if _, err := validate.IsUrl(urlWrap.Url); err != nil {
		halt(&urlWrap, response, w, err)
		return
	}

	shortUrl, err := shorten.Shorten(urlWrap.Url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	setScheme(r)
	urlWrap.Url = fmt.Sprintf("%s://%s/%s", r.URL.Scheme, r.Host, shortUrl)

	w.WriteHeader(http.StatusCreated)
	response.Encode(urlWrap)
}

//
func halt(urlWrap *urlWrapper, response *json.Encoder, w http.ResponseWriter, err error) {
	urlWrap.Url = err.Error()
	response.Encode(urlWrap)
}

func setScheme(r *http.Request) {
	if r.TLS == nil {
		r.URL.Scheme = "http"
	} else {
		r.URL.Scheme = "https"
	}
}
