package models

import (
	"github.com/astaxie/beego/orm"
)

// Todo is a model
type Todo struct {
	Base
	Task *string `json:"task" orm:"column(task)" validate:"required"`
	Done bool    `json:"done" orm:"column(done)"`
}

// TableName It will returns table name
func (a *Todo) TableName() string {
	return "todos"
}

func init() {
	orm.RegisterModel(new(Todo))
}

// TodoQuery is non db operational model
type TodoQuery struct {
	Todo
	Pagination *Pagination
	Sort       *Order
}
