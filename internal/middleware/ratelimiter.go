package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiter(redisClient *redis.Client, limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			c.Error(errors.New("unauthorized user"))
			c.Abort()
			return
		}
		key := fmt.Sprintf("rate_limit:user:%v:%s", userID, c.FullPath())
		count, err := redisClient.Incr(c.Request.Context(), key).Result()
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		//setting the ttl
		if count == 1 {
			redisClient.Expire(c.Request.Context(), key, duration)
		}

		//checking the limit
		if count > int64(limit) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded , try again after some time",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
