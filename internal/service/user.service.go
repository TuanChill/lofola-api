package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/pkg/helpers"
	"github.com/tuanchill/lofola-api/pkg/logger"
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

func (u *UserService) UpdateAvatar(c *gin.Context) *string {
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := c.Request.FormFile("avatar")
	if err != nil {
		response.BadRequestError(c, response.ErrCodeInvalidRequest, "No file provided")
		return nil
	}

	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, "Cannot read file")
		return nil
	}

	fileType := http.DetectContentType(buff)
	if fileType != "image/jpeg" && fileType != "image/jpg" && fileType != "image/png" {
		response.BadRequestError(c, response.ErrCodeInvalidRequest, "The provided file format is not allowed. Please upload a JPEG or PNG image")
		return nil
	}

	// Reset the file pointer so that the entire file can be read
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, err.Error())
		return nil
	}

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll(constants.PathUploadAvatar, os.ModePerm)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, "Cannot create uploads folder")
		return nil
	}

	// Create a new file in the uploads directory

	filePath := fmt.Sprintf("%s/%d%s", constants.PathUploadAvatar, time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	dst, err := os.Create(filePath)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, "Cannot create file")
		return nil
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, "Cannot copy file")
		return nil
	}

	// Get user_id info from token
	payload := helpers.GetPayload(c)

	// update user avatar
	if err := repo.UpdateAvatar(global.MDB, payload.ID, filePath); err != nil {
		logger.LogError(fmt.Sprintf("Cannot update avatar: %s", err.Error()))
		response.InternalServerError(c, response.ErrCodeDBConnection, err.Error())
		return nil
	}

	go func() {
		// Resize image
		if err := helpers.ReSizeImageForAvatar(filePath, filePath); err != nil {
			logger.LogError(fmt.Sprintf("Cannot resize image: %s", err.Error()))
		}
	}()

	return &filePath
}
