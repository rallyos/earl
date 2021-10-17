package url

import (
	"earl/db"
	"github.com/boltdb/bolt"
)

// TODO(Research): Keep struct, or just string encode/decode a struct?
func Create(shortUrl, longUrl string) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		err := b.Put([]byte(shortUrl), []byte(longUrl))
		return err
	})

	return err
}

func Get(shortUrl string) (string, error) {
	var longUrl []byte
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		longUrl = b.Get([]byte(shortUrl))
		return nil
	})
	return string(longUrl), err
}
