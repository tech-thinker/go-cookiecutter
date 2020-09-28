package instance

import (
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"

	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

type instance struct {
	db       orm.Ormer
	validate *validator.Validate
}

var singleton = &instance{}
var once sync.Once

// Init initializes the instance
func Init() {
	pgConfig := config.Postgres()
	once.Do(func() {

		// Validator initialization
		singleton.validate = validator.New()

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

// Validator returns the validator
func Validator() *validator.Validate {
	return singleton.validate
}
