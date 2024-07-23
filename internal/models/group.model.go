package models

import "time"

type Group struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"size:50;not null;unique"`
	Description string     `json:"description" gorm:"size:255"`
	IsPublic    bool       `json:"is_public" gorm:"default:false"`
	OwnerID     uint       `json:"owner_id" gorm:"not null"`
	CreateAt    time.Time  `json:"create_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateAt    *time.Time `json:"update_at"`
}
