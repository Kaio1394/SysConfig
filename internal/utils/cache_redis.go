package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func SaveCache(key string, value string) error {
	if redisClient == nil {
		InitRedisClient()
	}
	if err := redisClient.Set(context.Background(), key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}
