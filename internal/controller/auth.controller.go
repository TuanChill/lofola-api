package controller

import (
	"fmt"

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
	fmt.Println(result)
	response.Created(c, "Register successfully", result, nil)
	return nil
}
