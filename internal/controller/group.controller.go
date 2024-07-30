package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/service"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type GroupController struct {
}

func NewGroupController() *GroupController {
	return &GroupController{}
}

func (g *GroupController) GetGroup(c *gin.Context) error {
	result := service.NewGroupService().GetGroup(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Get Group Successfully", result)
	return nil
}

func (g *GroupController) UpdateGroup(c *gin.Context) error {
	result := service.NewGroupService().UpdateGroup(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Update Group Successfully", result)
	return nil
}

func (g *GroupController) CreateGroup(c *gin.Context) error {
	result := service.NewGroupService().CreateGroup(c)
	if result == nil {
		return nil
	}

	response.Created(c, "Create Group Successfully", result)
	return nil
}

func (g *GroupController) SearchGroup(c *gin.Context) error {
	result := service.NewGroupService().SearchGroup(c)
	if result == nil {
		return nil
	}

	response.ListDataResponse(c, "Search Group Successfully", result.Data, result.MetaData)
	return nil
}
