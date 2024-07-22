package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/middleware"
)

func NewRouter() *gin.Engine {
	// set up mode for gin
	var r *gin.Engine
	mode := global.Config.Server.Mode
	if mode == constants.DevMode {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	//init middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.Cors())

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
