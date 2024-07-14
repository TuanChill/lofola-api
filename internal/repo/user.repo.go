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
