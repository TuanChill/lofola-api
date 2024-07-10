package routers

import "github.com/gin-gonic/gin"

func AuthRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "register",
			})
		})
	}
}
