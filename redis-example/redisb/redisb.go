package redisb

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = Connect()

func Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Set(key, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

func GetAll() []string {
	return rdb.Keys(ctx, "*").Val()
}
