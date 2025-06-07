package repository

import (
    "context"
    "github.com/redis/go-redis/v9"
    "time"
    "errors"
)

var ctx = context.Background()
var rdb *redis.Client

func SetRedisClient(client *redis.Client) {
	rdb = client
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	return rdb
}

func SaveToCache(shortCode string, originalURL string) error {
    if rdb == nil {
		return errors.New("Redis client not initialized")
	}
    return rdb.Set(ctx, shortCode, originalURL, 24*time.Hour).Err()
}

func GetFromCache(shortCode string) (string, error) {
    if rdb == nil {
		return "",errors.New("Redis client not initialized")
	}
    return rdb.Get(ctx, shortCode).Result()
}