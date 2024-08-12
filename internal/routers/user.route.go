package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/middleware"
	"github.com/tuanchill/lofola-api/internal/wire"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func UserRouter(r *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	user := r.Group("/user")
	{
		user.Use(middleware.AuthenMiddleware())
		user.GET("/profile", utils.AsyncHandler(userController.GetProfile))
		user.POST("/profile", utils.AsyncHandler(userController.UpdateProfile))
		user.POST("/set-avatar", utils.AsyncHandler(userController.SetAvatar))
	}
}
