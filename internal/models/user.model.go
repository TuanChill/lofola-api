package models

import (
	"time"
)

type User struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	UserName string     `json:"username" gorm:"size:50;not null;unique"`
	Password string     `json:"password"`
	Email    string     `json:"email" gorm:"size:255;not null;unique"`
	Phone    *string    `json:"phone" gorm:"size:10"`
	FullName *string    `json:"full_name" gorm:"size:100"`
	Avatar   *string    `json:"avatar"`
	Gender   *bool      `json:"gender" gorm:"default:null"`
	IsActive bool       `json:"is_active" gorm:"default:false"`
	CreateAt time.Time  `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}

// register
type UserRequestBody struct {
	UserName string `json:"username" binding:"required,min=6,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone,omitempty" binding:"omitempty,min=10,max=10"`
}

type UserResponseBody struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// Span user Redis
type SpamUserRedisResponse struct {
	ExpireTime int  `json:"expire_spam"`
	IsSpam     bool `json:"is_spam"`
}
