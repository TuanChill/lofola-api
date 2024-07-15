package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/internal/repo/redis"
	"github.com/tuanchill/lofola-api/pkg/helpers"
	"github.com/tuanchill/lofola-api/pkg/response"
	"github.com/tuanchill/lofola-api/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Register(c *gin.Context) *models.UserResponseBody {
	// get data from body
	reqBody := models.UserRequestBody{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestErrorWithFields(c, response.ErrCodeValidation, utils.GetObjMessage(err))
		return nil
	}

	// check spam user
	if spamResponse := redis.SpamUser(c, global.RDB, constants.SpanKey, constants.RequestThreshold); spamResponse != nil {
		if spamResponse.IsSpam {
			ttl := fmt.Sprintf("You are blocked for %d seconds", spamResponse.ExpireTime)
			response.BadRequestError(c, response.ErrIpBlackList, ttl)
			return nil
		}
	}

	// get user by email
	user, err := repo.GetDetailUserByEmail(global.MDB, reqBody.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
			return nil
		}
	}

	// if user exists
	if user.ID > 0 {
		response.BadRequestError(c, response.ErrCodeResourceConflict, "Email already exists")
		return nil
	}

	// get user by username
	oldUser, err := repo.GetDetailUserByUsername(global.MDB, reqBody.UserName)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
			return nil
		}
	}

	// if user exists
	if oldUser.ID > 0 {
		response.BadRequestError(c, response.ErrCodeResourceConflict, "UserName already exists")
		return nil
	}

	// hash password
	hashedPassword, err := helpers.HashPassword(reqBody.Password)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer)
		return nil
	}

	fmt.Println("hashedPassword", hashedPassword)

	reqBody.Password = hashedPassword

	// create user
	newUser, err := repo.CreateUser(global.MDB, reqBody)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery)
		return nil
	}

	return &models.UserResponseBody{
		UserName: newUser.UserName,
		Email:    newUser.Email,
	}
}
