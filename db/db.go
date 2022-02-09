// Package db provides ...
package db

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func Connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	RDB = rdb
}
