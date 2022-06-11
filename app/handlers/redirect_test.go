package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shiftingphotons/earl/app/shorten"

	"github.com/matryer/is"
)

// TODO: Do we need to initialize the DB?
func TestRedirect(t *testing.T) {
	is := is.New(t)

	longUrl := "https://shiftingphotons.dev/things/internet-historian-engoodening-of-no-mans-sky/"
	shortUrl, err := createURL(longUrl)
	is.NoErr(err)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", shortUrl), nil)
	w := httptest.NewRecorder()

	server.Router.ServeHTTP(w, req)

	expectedCode := http.StatusTemporaryRedirect
	statusCode := w.Result().StatusCode
	redirectUrl, err := w.Result().Location()
	is.NoErr(err)

	is.Equal(statusCode, expectedCode)
	is.Equal(redirectUrl.String(), longUrl)
}

func createURL(longUrl string) (string, error) {
	shortenedUrl, err := shorten.Shorten(longUrl)
	return shortenedUrl, err
}
