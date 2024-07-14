package controller

import (
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
	response.Created(c, "Register", result, nil)
	return nil
}
