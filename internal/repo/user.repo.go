package repo

import (
	"time"

	"github.com/tuanchill/lofola-api/internal/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	GetDetailUserByEmail(db *gorm.DB, email string) (models.User, error)
	GetDetailUserByUsername(db *gorm.DB, username string) (models.User, error)
	CreateUser(db *gorm.DB, data models.UserRequestBody) (models.User, error)
	ActiveUser(db *gorm.DB, user models.User) error
	ChangePassword(db *gorm.DB, user models.User, newPassword string) error
	UpdateUser(db *gorm.DB, userID int, data *models.UserProfileUpdateRq) error
	GetInfoUser(db *gorm.DB, userId int) (models.UserInfo, error)
	UpdateAvatar(db *gorm.DB, userID int, avatar string) error
}

type userRepo struct{}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

// GetDetailUserByEmail get user by email
func (u *userRepo) GetDetailUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	record := db.Where("email = ?", email).First(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

// GetDetailUserByUsername get user by username
func (u *userRepo) GetDetailUserByUsername(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	record := db.Where("user_name = ?", username).First(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

// CreateUser create new user form data
func (u *userRepo) CreateUser(db *gorm.DB, data models.UserRequestBody) (models.User, error) {
	user := models.User{
		UserName: data.UserName,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
	}

	record := db.Create(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

// ActiveUser active user by id , change is_active to true
func (u *userRepo) ActiveUser(db *gorm.DB, user models.User) error {
	record := db.Model(&user).Update("is_active", true)

	record.Update("update_at", time.Now())

	if record.Error != nil {
		return record.Error
	}

	return nil
}

// ChangePassword change password of user
func (u *userRepo) ChangePassword(db *gorm.DB, user models.User, newPassword string) error {
	record := db.Model(&user).Update("password", newPassword)

	record.Update("update_at", time.Now())

	if record.Error != nil {
		return record.Error
	}

	return nil
}

// UpdateAvatar update avatar of user
func (u *userRepo) UpdateUser(db *gorm.DB, userID int, data *models.UserProfileUpdateRq) error {

	record := db.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"full_name": data.FullName,
		"phone":     data.Phone,
		"gender":    data.Gender,
		"birth_day": data.BirthDay,
		"update_at": time.Now(),
	})

	if record.Error != nil {
		return record.Error
	}

	return nil
}

// GetInfoUser get user info by user_id
func (u *userRepo) GetInfoUser(db *gorm.DB, userId int) (models.UserInfo, error) {
	var userRes models.UserInfo
	record := db.Model(&models.User{}).Where("id = ?", userId).Scan(&userRes)

	if record.Error != nil {
		return models.UserInfo{}, record.Error
	}

	return userRes, nil
}

// UpdateAvatar update avatar of user from user_id
func (u *userRepo) UpdateAvatar(db *gorm.DB, userID int, avatar string) error {
	record := db.Model(&models.User{}).Where("id = ?", userID).Update("avatar", avatar)

	record.Update("update_at", time.Now())

	if record.Error != nil {
		return record.Error
	}

	return nil
}
