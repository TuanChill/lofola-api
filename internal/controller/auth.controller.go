package controller

import "github.com/gin-gonic/gin"

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// register controller
func (a *AuthController) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}
