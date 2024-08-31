package models

import (
	"strings"
	"time"
)

type MetaData struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

type SearchParam struct {
	KeyWord string `form:"keyword" json:"keyword" binding:"omitempty"`
	Page    int    `form:"page" json:"page" binding:"omitempty,gte=1" `
	Limit   int    `form:"limit" json:"limit" binding:"omitempty,gte=10,lte=100"`
}

type TimeOnly struct {
	time.Time
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format("15:04") + `"`), nil
}

func (t *TimeOnly) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	parsedTime, err := time.Parse("15:04", str)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
