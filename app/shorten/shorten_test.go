package shorten

import (
	"testing"

	"github.com/matryer/is"
	"github.com/rallyos/earl/db"
)

func init() {
	db.Connect()
}

func TestShorten(t *testing.T) {
	is := is.New(t)

	url := "https://rallyo.io/things/internet-historian-engoodening-of-no-mans-sky/"
	shortenedUrl, err := Shorten(url)

	is.NoErr(err)
	is.True(len(shortenedUrl) <= 7)
}

func TestEncodeB62(t *testing.T) {
	is := is.New(t)

	buf := encodeB62(11157)
	if buf.String() != "73C" {
		is.Equal("73C", buf.String())
	}
}

func TestGenerateRand(t *testing.T) {
	is := is.New(t)

	numOne := generateRand()
	numTwo := generateRand()

	is.True(numOne != numTwo)
}
