package db

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {

	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return db
}
