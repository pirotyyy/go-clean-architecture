package cache

import (
	"github.com/redis/go-redis/v9"
	"os"
)

type CacheHandler struct {
	RedisConn *redis.Client
}

func CacheConnector() *CacheHandler {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	return &CacheHandler{
		RedisConn: redisClient,
	}
}
