package middleware

import (
	"net/http"
	"time"

	"github.com/ayanAhm4d/URL-shortener/redis"
	"github.com/gin-gonic/gin"
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate:" + ip

		count, err := redis.Client.Get(redis.Ctx, key).Int()
		if err == nil && count >= 10 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		redis.Client.Incr(redis.Ctx, key)
		redis.Client.Expire(redis.Ctx, key, 30*time.Minute)

		c.Next()
	}
}
