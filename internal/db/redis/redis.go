package redis

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return client
}
