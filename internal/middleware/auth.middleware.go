package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// check token is set
		if token == "" {
			response.UnauthorizedError(c, 401, "Token is required")
		}
	}
}
