package instance

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

type instance struct {
	db orm.Ormer
}

var singleton = &instance{}
var once sync.Once

// Init initializes the instance
func Init() {
	once.Do(func() {
		singleton.db = orm.NewOrm()
	})
}

// Destroy closes the connections & cleans up the instance
func Destroy() error {
	return nil
}

// DB will return the instance of database
func DB() orm.Ormer {
	return singleton.db
}
