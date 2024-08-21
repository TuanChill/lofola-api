package models

type Restaurant struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"size:255;not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Address     string `json:"address" gorm:"not null"`
	WardID      uint   `json:"ward_id" gorm:"not null"`
	LocationURL string `json:"location_url" gorm:"size:255;not null"`
}
