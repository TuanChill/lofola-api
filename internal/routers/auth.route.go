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
		auth.POST("/send-otp", utils.AsyncHandler(controller.NewAuthController().ResendOtp))
		auth.POST("/verify-otp", utils.AsyncHandler(controller.NewAuthController().VerifyOtp))
		auth.POST("/logout", utils.AsyncHandler(controller.NewAuthController().Logout))
		auth.POST("/reset-password", utils.AsyncHandler(controller.NewAuthController().ResetPassword))
	}
}
