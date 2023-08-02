package main

import (
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"os"
)

func RedisClient() {
	godotenv.Load()
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opts)
	// TODO: Fix issue
}
