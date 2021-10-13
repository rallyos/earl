package shorten

import (
	"bytes"
	"crypto/rand"
	"math"
	"math/big"
)

type Url struct {
	ID       uint64
	ShortURL string
	LongURL  string
}

// TODO URL Struct?
// TODO: Shorten Interface that has to implement shorten method?

var b62 [62]byte = [62]byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

//
func Shorten(url string) string {
	id := generateRand()
	shortUrl := hashToB62(id)

	u := Url{
		ID:       id,
		ShortURL: shortUrl.String(),
		LongURL:  url,
	}

	return ""
}

func hashToB62(n uint64) bytes.Buffer {
	var r int
	var buf bytes.Buffer

	for i := 0; n > 0 && i < 7; i++ {
		n, r = n/62, int(n%62)
		b := b62[r]
		buf.WriteByte(b)
	}

	return buf
}

func generateRand() uint64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	return r.Uint64()
}
