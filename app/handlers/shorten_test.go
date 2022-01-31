package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestShorten(t *testing.T) {
	is := is.New(t)

	// TODO INIT DB?

	// requestBody := struct {
	// 	Url string
	// }{
	// 	Url: "https://shiftingphotons.dev/things/internet-historian-engoodening-of-no-mans-sky/",
	// }

	requestBody := []byte(`{url: "https://shiftingphotons.dev/things/internet-historian-engoodening-of-no-mans-sky/"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

}
