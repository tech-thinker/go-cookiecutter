package models

import "time"

// Base is a model is used common attributes of models
type Base struct {
	ID        int64  `json:"id" orm:"column(id)"`
	CreatedAt time.Time `json:"updated_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"created_at" orm:"auto_now;type(datetime)"`
}

type Pagination struct {
	Page  *int64 `json:"page"`
	Limit *int64 `json:"limit"`
}

type Order struct {
	OrderBy   *string
	SortOrder *string
}
