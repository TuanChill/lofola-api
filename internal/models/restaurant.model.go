package models

import "time"

type Restaurant struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"size:255;not null"`
	Description string     `json:"description" gorm:"type:text;not null"`
	Address     string     `json:"address" gorm:"not null"`
	WardID      uint       `json:"ward_id" gorm:"not null"`
	LocationURL string     `json:"location_url" gorm:"size:255;not null"`
	Phone       string     `json:"phone" gorm:"size:20;not null"`
	OpenTime    TimeOnly   `json:"open_time" gorm:"type:time;not null"`
	CloseTime   TimeOnly   `json:"close_time" gorm:"type:time;not null"`
	IsActive    bool       `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
