package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/service"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// register controller
func (a *AuthController) Register(c *gin.Context) error {
	result := service.NewAuthService().Register(c)
	if result == nil {
		return nil
	}
	response.Created(c, "Register successfully", result, nil)
	return nil
}

// register controller
func (a *AuthController) Login(c *gin.Context) error {
	result := service.NewAuthService().Login(c)
	if result == nil {
		return nil
	}

	c.Header("Authorization", strings.Join([]string{"Bearer", result.Token.AccessToken}, " "))
	c.Header("RefreshToken", result.Token.RefreshToken)

	response.Ok(c, "Login successfully", gin.H{
		"email":    result.Email,
		"userName": result.UserName,
	})
	return nil
}

func (a *AuthController) VerifyOtp(c *gin.Context) error {
	result := service.NewAuthService().Verify(c)
	if !result {
		return nil
	}
	response.Ok(c, "Verify successfully", result)
	return nil
}

func (a *AuthController) ResendOtp(c *gin.Context) error {
	result := service.NewAuthService().ResendOtp(c)
	if !result {
		return nil
	}
	response.Ok(c, "Send OTP successfully", result)
	return nil
}
