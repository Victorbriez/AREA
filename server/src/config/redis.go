package config

import (
	"github.com/go-redis/redis/v8"
	"os"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:   0,
	})
}
