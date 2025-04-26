
package middlewares

import (
	"time"
	"github.com/gin-gonic/gin"
	"go-laravel-like/config"
)

func RedisRateLimiterMiddleware(limit time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "ratelimit:" + ip

		val, _ := config.RedisClient.Get(config.RedisCtx, key).Result()
		if val != "" {
			c.JSON(429, gin.H{
				"success": false,
				"code": 429,
				"message": "Too Many Requests",
			})
			c.Abort()
			return
		}

		config.RedisClient.Set(config.RedisCtx, key, "1", limit)
		c.Next()
	}
}
