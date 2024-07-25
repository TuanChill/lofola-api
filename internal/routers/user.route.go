package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/controller"
	"github.com/tuanchill/lofola-api/internal/middleware"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func UserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.Use(middleware.AuthenMiddleware())
		user.GET("/profile", utils.AsyncHandler(controller.NewUserController().GetProfile))
		user.POST("/profile", utils.AsyncHandler(controller.NewUserController().UpdateProfile))
	}
}
