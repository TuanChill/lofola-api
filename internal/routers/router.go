package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//init middleware
	r.Use(middleware.LoggerMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			AuthRouter(v1)
		}
	}

	return r
}
