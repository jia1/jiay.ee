package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/nitishm/go-rejson/v4"
	"os"
)

func RedisClients() (*redis.Client, *rejson.Handler) {
	godotenv.Load()
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(opts)
	// https://github.com/nitishm/go-rejson/issues/52
	redisJSONHandler := rejson.NewReJSONHandler()
	redisJSONHandler.SetGoRedisClient(redisClient)
	return redisClient, redisJSONHandler
}
