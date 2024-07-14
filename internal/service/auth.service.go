package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Register(c *gin.Context) *gin.H {
	// get data from body
	reqBody := models.UserRequestBody{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestError(c, response.ErrCodeValidation)
		return nil
	}

	// get user by email
	if _, err := repo.GetDetailUserByEmail(global.MDB, reqBody.Email); err != nil {
		response.BadRequestError(c, response.ErrCodeResourceConflict)
		return nil
	}

	return nil
}
