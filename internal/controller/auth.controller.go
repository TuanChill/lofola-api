package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/service"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authServer service.IAuthService) *AuthController {
	return &AuthController{
		authService: authServer,
	}
}

// Register controller
func (a *AuthController) Register(c *gin.Context) error {
	result := a.authService.Register(c)
	if result == nil {
		return nil
	}
	response.Created(c, "Register successfully", result)
	return nil
}

// Login controller
func (a *AuthController) Login(c *gin.Context) error {
	result := a.authService.Login(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Login successfully", result)
	return nil
}

func (a *AuthController) VerifyOtp(c *gin.Context) error {
	result := a.authService.Verify(c)
	if !result {
		return nil
	}
	response.Ok(c, "Verify successfully", result)
	return nil
}

func (a *AuthController) ResendOtp(c *gin.Context) error {
	result := a.authService.ResendOtp(c)
	if !result {
		return nil
	}
	response.Ok(c, "Send OTP successfully", result)
	return nil
}

func (a *AuthController) Logout(c *gin.Context) error {
	result := a.authService.Logout(c)
	if !result {
		return nil
	}
	response.Ok(c, "Logout successfully", result)
	return nil
}

func (a *AuthController) ResetPassword(c *gin.Context) error {
	result := a.authService.ResetPassword(c)
	if !result {
		return nil
	}
	response.Ok(c, "Reset password successfully", result)
	return nil
}

func (a *AuthController) RefreshToken(c *gin.Context) error {
	result := a.authService.RefreshToken(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Get new access token successfully", result)
	return nil
}
