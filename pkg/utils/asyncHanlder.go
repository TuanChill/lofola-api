package utils

import "github.com/gin-gonic/gin"

func AsyncHandler(handler func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			c.Error(err)
			c.Next()
		}
	}
}
