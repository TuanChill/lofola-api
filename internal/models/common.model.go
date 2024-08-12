package models

type MetaData struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

type SearchParam struct {
	KeyWord string `form:"keyword" json:"keyword" binding:"omitempty"`
	Page    int    `form:"page" json:"page" binding:"omitempty,gt=0" `
	Limit   int    `form:"limit" json:"limit" binding:"omitempty,gt=0,lte=50"`
}
