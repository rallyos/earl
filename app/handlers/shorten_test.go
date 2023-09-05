package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

// TODO: Do we need to initialize the DB?
func TestShorten(t *testing.T) {
	is := is.New(t)

	requestBody := []byte(`{url: "https://rallyo.io/things/internet-historian-engoodening-of-no-mans-sky/"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

}
