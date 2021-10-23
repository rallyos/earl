// Package db provides ...
package db

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func Connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RDB = rdb
}
