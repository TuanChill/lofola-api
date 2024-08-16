package service

import (
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

type IGroupService interface {
	GetGroup(c *gin.Context) *models.GroupInfo
	UpdateGroup(c *gin.Context) *models.GroupInfo
	CreateGroup(c *gin.Context) *models.GroupInfo
	GetInfoGroup(c *gin.Context) *models.GroupInfo
	SearchGroup(c *gin.Context) *models.GroupListResponse
	JoinGroup(c *gin.Context) bool
	LeaveGroup(c *gin.Context) bool
}

type groupService struct {
	groupRepo repo.IGroupRepo
	userRepo  repo.IUserRepo
}

func NewGroupService(groupRepo repo.IGroupRepo, userRepo repo.IUserRepo) IGroupService {
	return &groupService{
		groupRepo: groupRepo,
		userRepo:  userRepo,
	}
}

func (g *groupService) GetGroup(c *gin.Context) *models.GroupInfo {
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

	group, err := g.groupRepo.GetGroup(global.MDB, uint(groupID))
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
	owner, err := g.userRepo.GetInfoUser(global.MDB, int(group.OwnerID))
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
func (g *groupService) UpdateGroup(c *gin.Context) *models.GroupInfo {
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
	group, err := g.groupRepo.GetGroup(global.MDB, reqBody.ID)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return nil
	}

	// owner can update group
	payload := helpers.GetPayload(c)

	if group.OwnerID != uint(payload.ID) {
		response.ForbiddenError(c, response.ErrCodeForbidden, "You are not owner of this group")
		return nil
	}

	// check owner is set in request exist
	user, err := g.userRepo.GetInfoUser(global.MDB, int(reqBody.OwnerID))
	if err != nil {
		response.BadRequestError(c, response.ErrCodeInvalidInput, err.Error())
		return nil
	}

	if user.ID == 0 {
		response.BadRequestError(c, response.ErrCodeInvalidInput, "Owner not exist")
		return nil
	}

	// update group
	group, err = g.groupRepo.UpdateGroup(global.MDB, models.Group{
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
func (g *groupService) CreateGroup(c *gin.Context) *models.GroupInfo {
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
	group, err := g.groupRepo.CreateGroup(global.MDB, models.Group{
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

func (g *groupService) GetInfoGroup(c *gin.Context) *models.GroupInfo {
	groupId := c.Query("id")
	if groupId == "" {
		response.BadRequestError(c, response.ErrCodeInvalidRequest, "Group ID is required")
		return nil
	}

	return nil
}

func (g *groupService) SearchGroup(c *gin.Context) *models.GroupListResponse {
	var reqParam models.SearchParam

	// keyword := c.Query("keyword")

	if err := helpers.ValidateRequestSearch(c, &reqParam); err != nil {
		return nil
	}

	// set default value
	if reqParam.Limit == 0 {
		reqParam.Limit = 10
	}

	if reqParam.Page == 0 {
		reqParam.Page = 1
	}

	// search group
	groups, err := g.groupRepo.SearchGroup(global.MDB, reqParam)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return nil
	}

	result := models.GroupListResponse{
		Data: groups.Data,
		MetaData: models.MetaData{
			Page:  reqParam.Page,
			Limit: reqParam.Limit,
			Total: groups.Total,
		},
	}

	return &result
}

// JoinGroup is a function that join a group service
func (g *groupService) JoinGroup(c *gin.Context) bool {
	var reqBody models.GroupJoinRequest
	if IsErr := utils.BindRequest(c, &reqBody); !IsErr {
		return false
	}

	// get user_id from token
	payload := helpers.GetPayload(c)

	// check group exist
	isExist, err := g.groupRepo.CheckGroupExits(global.MDB, reqBody.GroupID)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return false
	}
	if !isExist {
		response.NotFoundError(c, response.ErrCodeNotFound, "Group not found")
		return false
	}

	// check user already joined group
	isJoined, err := g.groupRepo.CheckUserJoinedGroup(global.MDB, reqBody.GroupID, uint(payload.ID))
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return false
	}

	if isJoined {
		response.BadRequestError(c, response.ErrCodeInvalidInput, "You already joined this group")
		return false
	}

	// join group
	err = g.groupRepo.JoinGroup(global.MDB, reqBody.GroupID, uint(payload.ID))
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return false
	}

	return true
}

// LeaveGroup is a function that leave a group service return bool value
func (g *groupService) LeaveGroup(c *gin.Context) bool {
	var reqBody models.GroupJoinRequest
	if IsErr := utils.BindRequest(c, &reqBody); !IsErr {
		return false
	}

	// check group exist
	isExist, err := g.groupRepo.CheckGroupExits(global.MDB, reqBody.GroupID)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return false
	}
	if !isExist {
		response.NotFoundError(c, response.ErrCodeNotFound, "Group not found")
		return false
	}

	// get user_id from token
	payload := helpers.GetPayload(c)

	// check user already joined group
	isJoined, err := g.groupRepo.CheckUserJoinedGroup(global.MDB, reqBody.GroupID, uint(payload.ID))
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return false
	}

	if !isJoined {
		response.BadRequestError(c, response.ErrCodeInvalidInput, "You not joined this group")
		return false
	}

	// leave group
	err = g.groupRepo.LeaveGroup(global.MDB, reqBody.GroupID, uint(payload.ID))
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery, err.Error())
		return false
	}

	return true
}
