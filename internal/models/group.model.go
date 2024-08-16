package models

import "time"

type Group struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:50;not null"`
	Description string    `json:"description" gorm:"size:255"`
	IsPublic    bool      `json:"is_public" gorm:"default:false"`
	OwnerID     uint      `json:"owner_id" gorm:"not null"`
	CreateAt    time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateAt    time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()"`
}

// GroupUser is a struct that defines the group contains user
type GroupUser struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	GroupID uint `json:"group_id" gorm:"not null"`
	UserID  uint `json:"user_id" gorm:"not null"`
}

// GroupUserRequest is a struct that defines the request body for adding user to group
type GroupCreateRequest struct {
	Name        string `json:"name" binding:"required" validate:"min=1,max=50"`
	Description string `json:"description" binding:"omitempty"`
	IsPublic    bool   `json:"is_public" binding:"omitempty"`
}

type GroupUpdateRequest struct {
	ID          uint   `json:"id" binding:"required,numeric"`
	Name        string `json:"name" binding:"required" validate:"min=1,max=50"`
	Description string `json:"description" binding:"omitempty"`
	IsPublic    bool   `json:"is_public" binding:"omitempty"`
	OwnerID     uint   `json:"owner_id" binding:"required,numeric"`
}

type GroupInfo struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	Owner       string    `json:"owner"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type GroupListResponse struct {
	Data     []GroupInfo `json:"data"`
	MetaData MetaData    `json:"meta_data"`
}

type GroupSearchResult struct {
	Data  []GroupInfo `json:"data"`
	Total int64       `json:"total"`
}

type GroupJoinRequest struct {
	GroupID uint `json:"group_id" binding:"required,numeric"`
}

type GroupLeaveRequest struct {
	GroupID uint `json:"group_id" binding:"required,numeric"`
}
