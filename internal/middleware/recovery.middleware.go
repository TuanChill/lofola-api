package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/pkg/logger"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Convert the recovered error into a string and log it
				panicErr, ok := err.(error)
				var errMsg string
				if ok {
					errMsg = panicErr.Error()
				} else {
					errMsg = fmt.Sprintf("%v", err)
				}
				fmt.Println(errMsg)
				logger.LogError(errMsg)
				response.InternalServerError(c, response.ErrCodeInternalServer, "INTERNAL_SERVER_ERROR")
				c.Abort()
			}
		}()
		c.Next()
		c.Next()
	}
}
