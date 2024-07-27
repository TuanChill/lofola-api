package repo

import (
	"time"

	"github.com/tuanchill/lofola-api/internal/models"
	"gorm.io/gorm"
)

// GetDetailUserByEmail get user by email
func GetDetailUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	record := db.Where("email = ?", email).First(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

// GetDetailUserByUsername get user by username
func GetDetailUserByUsername(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	record := db.Where("user_name = ?", username).First(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

// CreateUser create new user form data
func CreateUser(db *gorm.DB, data models.UserRequestBody) (models.User, error) {
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
func ActiveUser(db *gorm.DB, user models.User) error {
	record := db.Model(&user).Update("is_active", true)

	record.Update("update_at", time.Now())

	if record.Error != nil {
		return record.Error
	}

	return nil
}

// ChangePassword change password of user
func ChangePassword(db *gorm.DB, user models.User, newPassword string) error {
	record := db.Model(&user).Update("password", newPassword)

	record.Update("update_at", time.Now())

	if record.Error != nil {
		return record.Error
	}

	return nil
}

// UpdateAvatar update avatar of user
func UpdateUser(db *gorm.DB, userID int, data *models.UserProfileUpdateRq) error {

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
func GetInfoUser(db *gorm.DB, userId int) (models.UserInfo, error) {
	var userRes models.UserInfo
	record := db.Model(&models.User{}).Where("id = ?", userId).Scan(&userRes)

	if record.Error != nil {
		return models.UserInfo{}, record.Error
	}

	return userRes, nil
}

// UpdateAvatar update avatar of user from user_id
func UpdateAvatar(db *gorm.DB, userID int, avatar string) error {
	record := db.Model(&models.User{}).Where("id = ?", userID).Update("avatar", avatar)

	record.Update("update_at", time.Now())

	if record.Error != nil {
		return record.Error
	}

	return nil
}
