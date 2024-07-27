package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/pkg/logger"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.LogError(c.Err().Error())
				response.InternalServerError(c, response.ErrCodeInternalServer, "INTERNAL_SERVER_ERROR")
				c.Abort()
			}
		}()
		c.Next()
	}
}
