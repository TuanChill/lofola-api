package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/internal/service"
	"github.com/tuanchill/lofola-api/pkg/helpers"
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

	helpers.SetHeaderResponse(c.Writer, constants.AuthorizationHeader, strings.Join([]string{"Bearer", result.Token.AccessToken}, " "))
	helpers.SetHeaderResponse(c.Writer, constants.RefreshTokenHeader, result.Token.RefreshToken)

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

func (a *AuthController) Logout(c *gin.Context) error {
	result := service.NewAuthService().Logout(c)
	if !result {
		return nil
	}
	response.Ok(c, "Logout successfully", result)
	return nil
}

func (a *AuthController) ResetPassword(c *gin.Context) error {
	result := service.NewAuthService().ResetPassword(c)
	if !result {
		return nil
	}
	response.Ok(c, "Reset password successfully", result)
	return nil
}

func (a *AuthController) RefreshToken(c *gin.Context) error {
	result := service.NewAuthService().RefreshToken(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Get new access token successfully", result)
	return nil
}
