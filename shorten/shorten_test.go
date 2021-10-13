package shorten

import (
	"testing"
)

func TestShorten(t *testing.T) {
	url := "https://shiftingphotons.dev/things/internet-historian-engoodening-of-no-mans-sky/"
	shortenedUrl := Shorten(url)

	if len(shortenedUrl) > 7 {
		t.Errorf("Expected shortened url to have length lower than 7, got %d", len(shortenedUrl))
	}
}

func TestHashToB62(t *testing.T) {
	buf := hashToB62(11157)
	if buf.String() != "73C" {
		t.Errorf("Wrong value. Expected %s, got %s.", "73C", buf.String())
	}
}
