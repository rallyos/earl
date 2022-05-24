// Package db provides ...
package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func Connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
		Password: "",
		DB:       0,
	})

	RDB = rdb
}
