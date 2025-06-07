package repository

import (
    "context"
    "github.com/redis/go-redis/v9"
    "time"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        DB:   0,
    })
}

func SaveToCache(shortCode string, originalURL string) error {
    return rdb.Set(ctx, shortCode, originalURL, 24*time.Hour).Err()
}

func GetFromCache(shortCode string) (string, error) {
    return rdb.Get(ctx, shortCode).Result()
}