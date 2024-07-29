package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/pkg/helpers"
	"github.com/tuanchill/lofola-api/pkg/response"
	"github.com/tuanchill/lofola-api/pkg/utils"
	"gorm.io/gorm"
)

type GroupService struct {
}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (g *GroupService) GetGroup(c *gin.Context) *models.GroupInfo {
	groupId := c.Query("id")
	if groupId == "" {
		response.BadRequestError(c, response.ErrCodeInvalidRequest, "ID is required")
		return nil
	}

	// get group info
	groupID, err := strconv.ParseUint(groupId, 10, 64)
	if err != nil {
		response.BadRequestError(c, response.ErrCodeInvalidRequest, "ID must be a number")
		return nil
	}

	group, err := repo.NewGroupRepo().GetGroup(global.MDB, uint(groupID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundError(c, response.ErrCodeNotFound, "Group not found")
			return nil
		}

		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return nil
	}

	// check group not public
	if !group.IsPublic {
		response.ForbiddenError(c, response.ErrCodeForbidden, "Group is not public")
		return nil
	}

	// get owner info
	owner, err := repo.GetInfoUser(global.MDB, int(group.OwnerID))
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return nil
	}

	return &models.GroupInfo{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		IsPublic:    group.IsPublic,
		Owner:       owner.UserName,
		CreateAt:    group.CreateAt,
		UpdateAt:    group.UpdateAt,
	}
}

// UpdateGroup is a function that update a group
func (g *GroupService) UpdateGroup(c *gin.Context) *models.GroupInfo {
	var reqBody models.GroupUpdateRequest

	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		if err.Error() == "EOF" {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, "No data provided")
			return nil
		}

		response.BadRequestErrorWithFields(c, response.ErrCodeInvalidRequest, utils.GetObjMessage(err))
		return nil
	}

	// check group exist
	group, err := repo.NewGroupRepo().GetGroup(global.MDB, reqBody.ID)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return nil
	}

	// owner can update group
	payload := helpers.GetPayload(c)

	fmt.Println("payload.ID", payload.ID)
	fmt.Println("group.OwnerID", group.OwnerID)
	if group.OwnerID != uint(payload.ID) {
		response.ForbiddenError(c, response.ErrCodeForbidden, "You are not owner of this group")
		return nil
	}

	// check owner is set in request exist
	user, err := repo.GetInfoUser(global.MDB, int(reqBody.OwnerID))
	if err != nil {
		response.BadRequestError(c, response.ErrCodeInvalidInput, err.Error())
		return nil
	}

	if user.ID == 0 {
		response.BadRequestError(c, response.ErrCodeInvalidInput, "Owner not exist")
		return nil
	}

	// update group
	group, err = repo.NewGroupRepo().UpdateGroup(global.MDB, models.Group{
		ID:          reqBody.ID,
		Name:        reqBody.Name,
		Description: reqBody.Description,
		IsPublic:    reqBody.IsPublic,
		OwnerID:     reqBody.OwnerID,
	})

	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return nil
	}
	return &models.GroupInfo{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		IsPublic:    group.IsPublic,
		Owner:       user.UserName,
		CreateAt:    group.CreateAt,
		UpdateAt:    group.UpdateAt,
	}
}

// CreateGroup is a function that create and  returns a group
func (g *GroupService) CreateGroup(c *gin.Context) *models.GroupInfo {
	var reqBody models.GroupCreateRequest

	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		if err.Error() == "EOF" {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, "No data provided")
			return nil
		}

		response.BadRequestErrorWithFields(c, response.ErrCodeInvalidRequest, utils.GetObjMessage(err))
		return nil
	}

	// get user_id from token
	payload := helpers.GetPayload(c)

	// create group
	group, err := repo.NewGroupRepo().CreateGroup(global.MDB, models.Group{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		IsPublic:    reqBody.IsPublic,
		OwnerID:     uint(payload.ID),
	})
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return nil
	}

	// set name owner for group
	groupInfo := models.GroupInfo{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		IsPublic:    group.IsPublic,
		Owner:       payload.UserName,
		CreateAt:    group.CreateAt,
		UpdateAt:    group.UpdateAt,
	}

	return &groupInfo
}

func (g *GroupService) GetInfoGroup(c *gin.Context) *models.GroupInfo {
	groupId := c.Query("id")
	if groupId == "" {
		response.BadRequestError(c, response.ErrCodeInvalidRequest, "Group ID is required")
		return nil
	}

	return nil
}
