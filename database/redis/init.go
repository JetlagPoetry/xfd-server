package redis

import (
	"github.com/go-redis/redis"
	"os"
)

var RedisClient *redis.Client

func Init() {
	options := redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PASSWORD"),
		MaxRetries: 3,
	}
	RedisClient = redis.NewClient(&options)
}
