package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/service"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type GroupController struct {
	groupService service.IGroupService
}

func NewGroupController(groupService service.IGroupService) *GroupController {
	return &GroupController{
		groupService: groupService,
	}
}

func (g *GroupController) GetGroup(c *gin.Context) error {
	result := g.groupService.GetGroup(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Get Group Successfully", result)
	return nil
}

func (g *GroupController) UpdateGroup(c *gin.Context) error {
	result := g.groupService.UpdateGroup(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Update Group Successfully", result)
	return nil
}

func (g *GroupController) CreateGroup(c *gin.Context) error {
	result := g.groupService.CreateGroup(c)
	if result == nil {
		return nil
	}

	response.Created(c, "Create Group Successfully", result)
	return nil
}

func (g *GroupController) SearchGroup(c *gin.Context) error {
	result := g.groupService.SearchGroup(c)
	if result == nil {
		return nil
	}

	response.ListDataResponse(c, "Search Group Successfully", result.Data, result.MetaData)
	return nil
}

func (g *GroupController) JoinGroup(c *gin.Context) error {
	result := g.groupService.JoinGroup(c)
	if !result {
		return nil
	}

	response.Ok(c, "Join Group Successfully", result)
	return nil
}

func (g *GroupController) LeaveGroup(c *gin.Context) error {
	result := g.groupService.LeaveGroup(c)
	if !result {
		return nil
	}

	response.Ok(c, "Leave Group Successfully", result)
	return nil
}
