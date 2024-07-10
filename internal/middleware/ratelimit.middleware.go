package middleware

import "github.com/gin-gonic/gin"

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check rate limit
	}
}
