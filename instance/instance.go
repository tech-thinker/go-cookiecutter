package instance

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nats.go"

	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

type Instance interface {
	Destroy() error
	DB() orm.Ormer
	Validator() *validator.Validate
	NatsClient() *nats.Conn
	NatsJetstreamClient() nats.JetStreamContext
}

type instance struct {
	db            orm.Ormer
	validate      *validator.Validate
	nats          *nats.Conn
	natsJetstream nats.JetStreamContext
}

// Destroy closes the connections & cleans up the instance
func (instance *instance) Destroy() error {
	return nil
}

// DB will return the instance of database
func (instance *instance) DB() orm.Ormer {
	return instance.db
}

// Validator returns the validator
func (instance *instance) Validator() *validator.Validate {
	return instance.validate
}

// NatsClient returns the nats client
func (instance *instance) NatsClient() *nats.Conn {
	return instance.nats
}

// NatsJetstreamClient returns the nats jetstream client
func (instance *instance) NatsJetstreamClient() nats.JetStreamContext {
	return instance.natsJetstream
}

// Init initializes the instance
func Init(config config.Configuration) Instance {
	instance := &instance{}

	// Validator initialization
	instance.validate = validator.New()

	// Postgresql database configuration
	logger.Log.Info("Database connecting...")
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logger.Log.Fatal(err)
	}
	err = orm.RegisterDataBase("default", "postgres", config.PostgresConfig().ConnectionURL())
	if err != nil {
		logger.Log.Fatal(err)
	}
	instance.db = orm.NewOrm()
	instance.db.Using("default")
	// RunSyncdb used to auto generate table structure in database, it may produce error if there is no models
	orm.RunSyncdb("default", true, true)

	logger.Log.Info("Database connected successfully...")

	logger.Log.Info("Connecting to nats...")
	instance.nats, err = nats.Connect(
		config.QueueConfig().NatsURL(),
		nats.Name(config.QueueConfig().NatsClientName()),
		nats.UserInfo(config.QueueConfig().NatsUsername(), config.QueueConfig().NatsPassword()),
	)
	if err != nil {
		logger.Log.Fatal(err)
	}
	instance.natsJetstream, err = instance.nats.JetStream()
	if err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("Connected to nats successfully...")

	return instance
}
