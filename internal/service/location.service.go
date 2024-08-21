package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/pkg/helpers"
	"github.com/tuanchill/lofola-api/pkg/response"
)

type ILocationService interface {
	GetProvince(c *gin.Context) *models.ProvinceResponse
	GetDistrict(c *gin.Context) *models.DistrictResponse
	GetWard(c *gin.Context) *models.WardResponse
}

type locationService struct {
	locationRepo repo.ILocationRepo
}

func NewLocationService(locationRepo repo.ILocationRepo) ILocationService {
	return &locationService{
		locationRepo: locationRepo,
	}
}

// GetDistrict implements ILocationService.
func (l *locationService) GetDistrict(c *gin.Context) *models.DistrictResponse {
	var param models.DistrictParam

	if err := helpers.ValidateRequest(c, &param); err != nil {
		return nil
	}

	result, err := l.locationRepo.GetDistrict(global.MDB, param)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, err.Error())
		return nil
	}

	return &models.DistrictResponse{
		Data: result.Data,
		MetaData: models.MetaData{
			Total: result.Total,
			Page:  param.Page,
			Limit: param.Limit,
		},
	}
}

// GetProvince implements ILocationService.
func (l *locationService) GetProvince(c *gin.Context) *models.ProvinceResponse {
	var param models.ProvinceParam
	if err := helpers.ValidateRequest(c, &param); err != nil {
		return nil
	}

	result, err := l.locationRepo.GetProvince(global.MDB, param)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, err.Error())
		return nil
	}

	return &models.ProvinceResponse{
		Data: result.Data,
		MetaData: models.MetaData{
			Total: result.Total,
			Page:  param.Page,
			Limit: param.Limit,
		},
	}
}

// GetWard implements ILocationService.
func (l *locationService) GetWard(c *gin.Context) *models.WardResponse {
	var param models.WardParam
	if err := helpers.ValidateRequest(c, &param); err != nil {
		return nil
	}

	result, err := l.locationRepo.GetWard(global.MDB, param)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, err.Error())
		return nil
	}

	return &models.WardResponse{
		Data: result.Data,
		MetaData: models.MetaData{
			Total: result.Total,
			Page:  param.Page,
			Limit: param.Limit,
		},
	}
}
