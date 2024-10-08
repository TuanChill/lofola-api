package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/pkg/helpers"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// check token is set
		if token == "" || len(token) == 0 {
			response.UnauthorizedError(c, 401, "Missing authorization header")
			return
		}

		tokenStr := token[len("Bearer "):]

		_, err := helpers.VerifyToken(tokenStr, global.Config.Security.AccessTokenSecret.SecretKey)
		if err != nil {
			response.UnauthorizedError(c, 401, "Unauthorized access")
			return
		}

		c.Next()
	}
}
