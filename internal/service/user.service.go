package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/pkg/helpers"
	"github.com/tuanchill/lofola-api/pkg/response"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetInfoProfile(c *gin.Context) *models.UserInfo {
	// Get user_id info from token
	payload := helpers.GetPayload(c)

	// Get user info from user_id
	user, err := repo.GetInfoUser(global.MDB, payload.ID)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBConnection, "Internal server error")
		return nil
	}

	return &models.UserInfo{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Phone:    user.Phone,
		FullName: user.FullName,
		BirthDay: user.BirthDay,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}

func (u *UserService) UpdateProfile(c *gin.Context) *models.UserInfo {
	// Get data from request and validate
	var data *models.UserProfileUpdateRq

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		if err.Error() == "EOF" {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, "No data provided")
			return nil
		}

		if len(utils.GetObjMessage(err)) == 0 {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, err.Error())
			return nil
		}
		response.BadRequestErrorWithFields(c, response.ErrCodeInvalidRequest, utils.GetObjMessage(err))
		return nil
	}

	fmt.Println(data)

	// Get user_id info from token
	payload := helpers.GetPayload(c)

	// update user info
	if err := repo.UpdateUser(global.MDB, payload.ID, data); err != nil {
		response.InternalServerError(c, response.ErrCodeDBConnection, err.Error())
		return nil
	}

	// Get user info from user_id
	user, err := repo.GetInfoUser(global.MDB, payload.ID)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBConnection, err.Error())
		return nil
	}

	return &user
}
