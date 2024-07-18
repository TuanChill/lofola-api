package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/controller"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func AuthRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", utils.AsyncHandler(controller.NewAuthController().Register))
		auth.POST("/login", utils.AsyncHandler(controller.NewAuthController().Login))
	}
}
