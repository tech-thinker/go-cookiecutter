package models

// Base is a model is used common attributes of models
type Base struct {
	ID        int64  `json:"id" orm:"column(id)"`
	CreatedAt *int64 `json:"updated_at" orm:"column(created_at)"`
	UpdatedAt *int64 `json:"created_at" orm:"column(updated_at)"`
}

type Pagination struct {
	Page  *int64 `json:"page"`
	Limit *int64 `json:"limit"`
}

type Order struct {
	OrderBy   *string
	SortOrder *string
}
