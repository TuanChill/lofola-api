package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/wire"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func AuthRouter(r *gin.RouterGroup) {
	authController, _ := wire.InitAuthRouterHandler()

	auth := r.Group("/auth")
	{
		auth.POST("/register", utils.AsyncHandler(authController.Register))
		auth.POST("/login", utils.AsyncHandler(authController.Login))
		auth.POST("/send-otp", utils.AsyncHandler(authController.ResendOtp))
		auth.POST("/verify-otp", utils.AsyncHandler(authController.VerifyOtp))
		auth.POST("/logout", utils.AsyncHandler(authController.Logout))
		auth.POST("/reset-password", utils.AsyncHandler(authController.ResetPassword))
		auth.POST("/refresh-token", utils.AsyncHandler(authController.RefreshToken))
	}
}
