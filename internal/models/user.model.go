package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	UserName string         `json:"username" gorm:"size:50;not null;unique"`
	Password string         `json:"password"`
	Email    string         `json:"email" gorm:"size:255;not null;unique"`
	Phone    sql.NullString `json:"phone" gorm:"size:10"`
	FullName sql.NullString `json:"full_name" gorm:"size:100"`
	Avatar   sql.NullString `json:"avatar"`
	Gender   bool           `json:"gender"`
	CreateAt time.Time      `json:"create_at"`
	UpdateAt time.Time      `json:"update_at"`
}

// register
type UserRequestBody struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
}

type UserResponseBody struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
