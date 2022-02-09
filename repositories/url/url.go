package url

import (
	"context"
	"earl/db"
	"time"
)

const EXPIRATION = time.Minute * 7200

var ctx = context.Background()

func Create(shortUrl, longUrl string) error {
	if err := db.RDB.Set(ctx, shortUrl, longUrl, EXPIRATION).Err(); err != nil {
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
