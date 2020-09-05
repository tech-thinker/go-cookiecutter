package instance

import (
	"sync"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	"github.com/mrasif/gomvc/config"
	"github.com/mrasif/gomvc/logger"
)

type instance struct {
	db orm.Ormer
}

var singleton = &instance{}
var once sync.Once

// Init initializes the instance
func Init() {
	pgConfig := config.Postgres()
	once.Do(func() {

		// Postgresql database configuration
		logger.Log.Info("Database connecting...")
		err := orm.RegisterDriver("postgres", orm.DRPostgres)
		if err != nil {
			logger.Log.Fatal(err)
		}
		err = orm.RegisterDataBase("default", "postgres", pgConfig.ConnectionURL())
		if err != nil {
			logger.Log.Fatal(err)
		}
		singleton.db = orm.NewOrm()
		singleton.db.Using("default")
		// RunSyncdb used to auto generate table structure in database, it may produce error if there is no models
		orm.RunSyncdb("default", true, true)
		logger.Log.Info("Database connected successfully...")
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
