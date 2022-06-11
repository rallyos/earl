package url

import (
	"context"
	"fmt"
	"time"

	"github.com/shiftingphotons/earl/db"
)

const EXPIRATION = time.Minute * 120

var ctx = context.Background()

func Create(shortUrl, longUrl string) error {
	if err := db.RDB.Set(ctx, shortUrl, longUrl, EXPIRATION).Err(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func Get(shortUrl string) (string, error) {
	longUrl, err := db.RDB.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", err
	}

	return longUrl, nil
}
