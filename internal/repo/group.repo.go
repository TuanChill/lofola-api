package repo

import (
	"github.com/tuanchill/lofola-api/internal/models"
	"gorm.io/gorm"
)

type IGroupRepo interface {
	CreateGroup(db *gorm.DB, data models.Group) (*models.Group, error)
	UpdateGroup(db *gorm.DB, data models.Group) (*models.Group, error)
	GetGroup(db *gorm.DB, id uint) (*models.Group, error)
	SearchGroup(db *gorm.DB, param models.SearchParam) (*models.GroupSearchResult, error)
}

// GroupRepo is a struct that defines the group repository
type groupRepo struct {
}

// NewGroupRepo is a function that returns a new group repository
func NewGroupRepo() IGroupRepo {
	return &groupRepo{}
}

// CreateGroup is a function that returns a group
func (g *groupRepo) CreateGroup(db *gorm.DB, data models.Group) (*models.Group, error) {
	group := models.Group{
		Name:        data.Name,
		Description: data.Description,
		IsPublic:    data.IsPublic,
		OwnerID:     data.OwnerID,
	}
	record := db.Create(&group)

	if record.Error != nil {
		return nil, record.Error
	}

	// Reload the group from the database to ensure all fields are updated
	if err := db.First(&group, group.ID).Error; err != nil {
		return nil, err
	}

	return &group, nil
}

// UpdateGroup is a function that returns a group
func (g *groupRepo) UpdateGroup(db *gorm.DB, data models.Group) (*models.Group, error) {
	group := models.Group{
		Name:        data.Name,
		Description: data.Description,
		IsPublic:    data.IsPublic,
		OwnerID:     data.OwnerID,
	}
	record := db.Where("id = ?", data.ID).Updates(&group).First(&group)

	if record.Error != nil {
		return nil, record.Error
	}

	return &group, nil
}

// GetGroup is a function that returns a group
func (g *groupRepo) GetGroup(db *gorm.DB, id uint) (*models.Group, error) {
	group := models.Group{}
	record := db.First(&group, id)

	if record.Error != nil {
		return nil, record.Error
	}

	return &group, nil
}

func (g *groupRepo) SearchGroup(db *gorm.DB, param models.SearchParam) (*models.GroupSearchResult, error) {
	groups := []models.GroupInfo{}
	query := db.Model(&models.Group{})

	// not select not public group
	query = query.Where("is_public = ?", true)

	if param.KeyWord != "" {
		query = query.Where("name LIKE ?", "%"+param.KeyWord+"%").Where("description LIKE ?", "%"+param.KeyWord+"%")
	}

	if param.Limit != 0 {
		query = query.Limit(param.Limit)
	}

	if param.Page != 0 {
		query = query.Offset((param.Page - 1) * param.Limit)
	}
	query = query.Joins("JOIN users ON `groups`.owner_id = users.id").Select("`groups`.*, users.user_name as owner")

	if err := query.Find(&groups).Error; err != nil {
		return nil, err
	}

	// get total group
	total := int64(0)
	if err := db.Model(&models.Group{}).Count(&total).Error; err != nil {
		return nil, err
	}

	result := models.GroupSearchResult{
		Data:  groups,
		Total: total,
	}

	return &result, nil
}
