package models

import "time"

type Article struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"size:255;not null"`
	Content     string     `json:"content" gorm:"type:text;not null"`
	AuthorID    uint       `json:"author_id" gorm:"not null"`
	IsPublished bool       `json:"is_published" gorm:"default:true"`
	CreateAt    time.Time  `json:"create_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateAt    *time.Time `json:"update_at"`
}

type ImgOfArticle struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ArticleID uint   `json:"article_id" gorm:"not null"`
	ImgURL    string `json:"img_url" gorm:"size:255;not null"`
}
