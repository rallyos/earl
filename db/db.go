// Package db provides ...
package db

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func Connect() {
	db, err := bolt.Open("earl.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			return fmt.Errorf("create bucket %s", err)
		}
		return nil
	})

	// defer db.Close()

	DB = db
}
