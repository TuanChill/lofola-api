package repo

import (
	"github.com/tuanchill/lofola-api/internal/models"
	"gorm.io/gorm"
)

type ILocationRepo interface {
	GetProvince(db *gorm.DB, param models.ProvinceParam) (*models.ProvinceSearchResult, error)
	GetDistrict(db *gorm.DB, param models.DistrictParam) (*models.DistrictSearchResult, error)
	GetWard(db *gorm.DB, param models.WardParam) (*models.WardSearchResult, error)
}

type LocationRepo struct {
}

func NewLocationRepo() ILocationRepo {
	return &LocationRepo{}
}

// GetDistrict implements ILocationRepo.
func (l *LocationRepo) GetDistrict(db *gorm.DB, param models.DistrictParam) (*models.DistrictSearchResult, error) {
	district := []models.District{}
	query := db.Model(&models.District{})

	if param.ProvinceID != 0 {
		query = query.Where("province_id = ?", param.ProvinceID)
	}

	if param.KeyWord != "" {
		query = query.Where("name LIKE ?", "%"+param.KeyWord+"%")
	}

	if param.Limit != 0 {
		query = query.Limit(param.Limit)
	}

	if param.Page != 0 {
		query = query.Offset((param.Page - 1) * param.Limit)
	}

	if err := query.Find(&district).Error; err != nil {
		return nil, err
	}

	// get total group
	total := int64(0)
	if err := db.Model(&models.Group{}).Count(&total).Error; err != nil {
		return nil, err
	}

	result := models.DistrictSearchResult{
		Data:  district,
		Total: total,
	}

	return &result, nil
}

// GetProvince implements ILocationRepo.
func (l *LocationRepo) GetProvince(db *gorm.DB, param models.ProvinceParam) (*models.ProvinceSearchResult, error) {
	province := []models.Province{}
	query := db.Model(&models.Province{})

	if param.KeyWord != "" {
		query = query.Where("name LIKE ?", "%"+param.KeyWord+"%")
	}

	if param.Limit != 0 {
		query = query.Limit(param.Limit)
	}

	if param.Page != 0 {
		query = query.Offset((param.Page - 1) * param.Limit)
	}

	if err := query.Find(&province).Error; err != nil {
		return nil, err
	}

	// get total group
	total := int64(0)
	if err := db.Model(&models.Group{}).Count(&total).Error; err != nil {
		return nil, err
	}

	result := models.ProvinceSearchResult{
		Data:  province,
		Total: total,
	}

	return &result, nil
}

// GetWard implements ILocationRepo.
func (l *LocationRepo) GetWard(db *gorm.DB, param models.WardParam) (*models.WardSearchResult, error) {
	ward := []models.Ward{}
	query := db.Model(&models.Ward{})

	if param.DistrictID != 0 {
		query = query.Where("district_id = ?", param.DistrictID)
	}

	if param.KeyWord != "" {
		query = query.Where("name LIKE ?", "%"+param.KeyWord+"%")
	}

	if param.Limit != 0 {
		query = query.Limit(param.Limit)
	}

	if param.Page != 0 {
		query = query.Offset((param.Page - 1) * param.Limit)
	}

	if err := query.Find(&ward).Error; err != nil {
		return nil, err
	}

	// get total group
	total := int64(0)
	if err := db.Model(&models.Group{}).Count(&total).Error; err != nil {
		return nil, err
	}

	result := models.WardSearchResult{
		Data:  ward,
		Total: total,
	}

	return &result, nil
}
