package handlers

import (
	"earl/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

var server *app.Server

func TestPing(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest(http.MethodGet, "/api/ping", nil)
	w := httptest.NewRecorder()

	server.Router.ServeHTTP(w, req)

	expectedCode, expectedBody := http.StatusOK, "pong"
	statusCode, body := w.Result().StatusCode, w.Result().Body

	is.Equal(statusCode, expectedCode)
	is.Equal(body, expectedBody)
}
