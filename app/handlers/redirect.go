package handlers

import (
	"earl/repositories/url"
	"net/http"
)

//
func Redirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.RequestURI[1:]

	longUrl, err := url.Get(shortUrl)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
}
