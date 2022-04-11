package database

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func GetRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
}

func GetRedisClient() *redis.Client {
	rdb := redis.NewClient(GetRedisConfig())

	return rdb
}
