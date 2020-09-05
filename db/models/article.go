package models

import "time"

// Article is a model
type Article struct {
	ID      int64     `json:"id" orm:"column(id)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
