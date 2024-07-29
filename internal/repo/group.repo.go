package repo

import (
	"github.com/tuanchill/lofola-api/internal/models"
	"gorm.io/gorm"
)

// GroupRepo is a struct that defines the group repository
type GroupRepo struct {
}

// NewGroupRepo is a function that returns a new group repository
func NewGroupRepo() *GroupRepo {
	return &GroupRepo{}
}

// CreateGroup is a function that returns a group
func (g *GroupRepo) CreateGroup(db *gorm.DB, data models.Group) (*models.Group, error) {
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
func (g *GroupRepo) UpdateGroup(db *gorm.DB, data models.Group) (*models.Group, error) {
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
func (g *GroupRepo) GetGroup(db *gorm.DB, id uint) (*models.Group, error) {
	group := models.Group{}
	record := db.First(&group, id)

	if record.Error != nil {
		return nil, record.Error
	}

	return &group, nil
}
