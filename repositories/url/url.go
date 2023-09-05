package url

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rallyos/earl/db"
)

var ctx = context.Background()

func Create(shortUrl, longUrl string) error {
	lifecycle := getRedisLifecycle()
	if err := db.RDB.Set(ctx, shortUrl, longUrl, lifecycle).Err(); err != nil {
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

func getRedisLifecycle() time.Duration {
	lifecycleEnv := os.Getenv("EARL_LIFECYCLE")

	if len(lifecycleEnv) == 0 {
		lifecycleEnv = "72"
	}

	lifecycle, err := strconv.Atoi(lifecycleEnv)
	if err != nil {
		panic("Lifecycle should be a number of hours.")
	}

	return time.Hour * time.Duration(lifecycle)
}
