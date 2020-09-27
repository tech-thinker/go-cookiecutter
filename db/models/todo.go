package models

import (
	"github.com/astaxie/beego/orm"
)

// Todo is a model
type Todo struct {
	Base
	Task *string `json:"task" orm:"column(task)"`
	Done bool    `json:"done" orm:"column(done)"`
}

func init() {
	orm.RegisterModel(new(Todo))
}

type TodoQuery struct {
	Todo
	Pagination *Pagination
	Sort       *Order
}
