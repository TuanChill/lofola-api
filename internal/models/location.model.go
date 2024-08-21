package models

type Province struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:255;not null"`
	Slug string `json:"slug" gorm:"size:255;not null"`
}

type District struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"size:255;not null"`
	ProvinceID uint   `json:"province_id" gorm:"not null"`
}

type Ward struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"size:255;not null"`
	DistrictID uint   `json:"district_id" gorm:"not null"`
}

type ProvinceParam struct {
	SearchParam
}

type DistrictParam struct {
	SearchParam
	ProvinceID int `form:"province_id" binding:"omitempty"`
}

type WardParam struct {
	SearchParam
	DistrictID int `form:"district_id" binding:"omitempty"`
}

type ProvinceSearchResult struct {
	Data  []Province `json:"data"`
	Total int64      `json:"total"`
}

type DistrictSearchResult struct {
	Data  []District `json:"data"`
	Total int64      `json:"total"`
}

type WardSearchResult struct {
	Data  []Ward `json:"data"`
	Total int64  `json:"total"`
}

type ProvinceResponse struct {
	Data     []Province `json:"data"`
	MetaData MetaData   `json:"meta_data"`
}

type DistrictResponse struct {
	Data     []District `json:"data"`
	MetaData MetaData   `json:"meta_data"`
}

type WardResponse struct {
	Data     []Ward   `json:"data"`
	MetaData MetaData `json:"meta_data"`
}
