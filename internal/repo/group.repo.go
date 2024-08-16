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
	JoinGroup(db *gorm.DB, groupID, userID uint) error
	LeaveGroup(db *gorm.DB, groupID, userID uint) error
	GetGroupUsers(db *gorm.DB, groupID uint) ([]models.User, error)
	GetGroupByUserID(db *gorm.DB, userID uint) (*models.Group, error)
	CheckUserJoinedGroup(db *gorm.DB, groupID, userID uint) (bool, error)
	CheckGroupExits(db *gorm.DB, id uint) (bool, error)
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

// JoinGroup is a function that joins a group
func (g *groupRepo) JoinGroup(db *gorm.DB, groupID uint, userID uint) error {
	groupUser := models.GroupUser{
		GroupID: groupID,
		UserID:  userID,
	}
	record := db.Create(&groupUser)

	if record.Error != nil {
		return record.Error
	}

	return nil
}

// LeaveGroup is a function that leaves a group from a user
func (g *groupRepo) LeaveGroup(db *gorm.DB, groupID uint, userID uint) error {
	groupUser := models.GroupUser{
		GroupID: groupID,
		UserID:  userID,
	}
	record := db.Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&groupUser)

	if record.Error != nil {
		return record.Error
	}

	return nil
}

// GetGroupUsers is a function that returns a list of users in a group
func (g *groupRepo) GetGroupUsers(db *gorm.DB, groupID uint) ([]models.User, error) {
	users := []models.User{}
	query := db.Model(&models.User{}).Joins("JOIN group_users ON users.id = group_users.user_id").Where("group_users.group_id = ?", groupID)

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetGroupByUserID is a function that returns a group by user ID was joined
func (g *groupRepo) GetGroupByUserID(db *gorm.DB, userID uint) (*models.Group, error) {
	group := models.Group{}
	query := db.Model(&models.Group{}).Joins("JOIN group_users ON groups.id = group_users.group_id").Where("group_users.user_id = ?", userID)

	if err := query.First(&group).Error; err != nil {
		return nil, err
	}

	return &group, nil
}

// CheckUserJoinedGroup is a function that checks if a user has joined a group returns true if joined
func (g *groupRepo) CheckUserJoinedGroup(db *gorm.DB, groupID, userID uint) (bool, error) {
	groupUser := models.GroupUser{}
	record := db.Where("group_id = ? AND user_id = ?", groupID, userID).First(&groupUser)

	if record.Error != nil {
		if record.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, record.Error
	}

	return true, nil
}

// checkGroupExits is a function that checks if a group exists
func (g *groupRepo) CheckGroupExits(db *gorm.DB, id uint) (bool, error) {
	record := db.Raw("SELECT count(id) as count FROM `groups` WHERE id = ?", id).Row()
	var count int
	if err := record.Scan(&count); err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
