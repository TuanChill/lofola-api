package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/service"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// get info profile
func (u *UserController) GetProfile(c *gin.Context) error {
	result := service.NewUserService().GetInfoProfile(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Get profile successfully", result)
	return nil
}

// profile update
func (u *UserController) UpdateProfile(c *gin.Context) error {
	result := service.NewUserService().UpdateProfile(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Update profile successfully", result)
	return nil
}

// update avatar
func (u *UserController) SetAvatar(c *gin.Context) error {
	result := service.NewUserService().UpdateAvatar(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Update avatar successfully", result)
	return nil
}
