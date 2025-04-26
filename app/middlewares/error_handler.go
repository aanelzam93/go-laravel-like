package middlewares

import (
	"net/http"
	"runtime/debug"

	"go-laravel-like/app/logs"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryWithLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error("Panic recovered",
					zap.Any("error", r),
					zap.ByteString("stack", debug.Stack()),
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"code":    500,
					"message": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}