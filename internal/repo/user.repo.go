package repo

import (
	"github.com/tuanchill/lofola-api/internal/models"
	"gorm.io/gorm"
)

func GetDetailUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	record := db.Where("email = ?", email).First(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

func GetDetailUserByUsername(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	record := db.Where("user_name = ?", username).First(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

func CreateUser(db *gorm.DB, data models.UserRequestBody) (models.User, error) {
	user := models.User{
		UserName: data.UserName,
		Password: data.Password,
		Email:    data.Email,
		Phone:    &data.Phone,
	}

	record := db.Create(&user)

	if record.Error != nil {
		return models.User{}, record.Error
	}

	return user, nil
}

func ActiveUser(db *gorm.DB, user models.User) error {
	record := db.Model(&user).Update("is_active", true)

	if record.Error != nil {
		return record.Error
	}

	return nil
}
