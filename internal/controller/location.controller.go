package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/service"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type LocationController struct {
	locationService service.ILocationService
}

func NewLocationController(locationService service.ILocationService) *LocationController {
	return &LocationController{
		locationService: locationService,
	}
}

// GetDistrict is a function to get district
func (l *LocationController) GetDistrict(c *gin.Context) error {
	result := l.locationService.GetDistrict(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Get District Successfully", result)
	return nil
}

// GetProvince is a function to get province
func (l *LocationController) GetProvince(c *gin.Context) error {
	result := l.locationService.GetProvince(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Get Province Successfully", result)
	return nil
}

// GetWard is a function to get ward
func (l *LocationController) GetWard(c *gin.Context) error {
	result := l.locationService.GetWard(c)
	if result == nil {
		return nil
	}

	response.Ok(c, "Get Ward Successfully", result)
	return nil
}
