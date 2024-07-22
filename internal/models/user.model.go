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
	CreateAt time.Time  `json:"create_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateAt *time.Time `json:"update_at"`
}

// otp
type UserResendOtp struct {
	Email string `json:"email" binding:"required,email"`
}

type UserVerifyOtp struct {
	Email string `json:"email" binding:"required,email"`
	Otp   int    `json:"otp" binding:"required,min=6,max=6,numeric`
}

// register
type UserRequestBody struct {
	UserName string `json:"username" binding:"required,min=6,max=50,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone,omitempty" binding:"omitempty,min=10,max=10"`
}

type UserResponseBody struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// login
type BodyLoginRequest struct {
	Identifier string `json:"identifier" binding:"required,min=6"`
	Password   string `json:"password" binding:"required,min=6,max=50"`
}

type LoginResponse struct {
	ID       int         `json:"id"`
	Email    string      `json:"email"`
	UserName string      `json:"userName"`
	Token    TokenStruct `json:"token"`
}

type TokenStruct struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Spam user Redis
type SpamUserRedisResponse struct {
	ExpireTime int  `json:"expire_spam"`
	IsSpam     bool `json:"is_spam"`
}

// payload token
type PayloadToken struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
}
