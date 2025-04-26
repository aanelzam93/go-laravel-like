package middlewares

import (
	"time"
	"github.com/gin-gonic/gin"
	"sync"
)

var visitors = make(map[string]time.Time)
var mu sync.Mutex

func RateLimiterMiddleware(limit time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		lastRequest, found := visitors[ip]
		if found && time.Since(lastRequest) < limit {
			mu.Unlock()
			c.JSON(429, gin.H{
				"success": false,
				"code": 429,
				"message": "Too Many Requests",
			})
			c.Abort()
			return
		}
		visitors[ip] = time.Now()
		mu.Unlock()
		c.Next()
	}
}