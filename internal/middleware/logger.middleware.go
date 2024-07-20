package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/pkg/logger"
	"go.uber.org/zap"
)

// LoggerMiddleware tạo ra một middleware để ghi log với Zap.
func LoggerMiddleware() gin.HandlerFunc {

	logger := logger.LoggerInstance()

	return func(c *gin.Context) {
		start := time.Now()

		// next to process request
		c.Next()

		// Ghi log sau khi xử lý xong request
		logger.Info("Request::",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("clientIP", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}
