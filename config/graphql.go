package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

// graphQL holds the config for the GraphQL
type graphQL struct {
	port string
}

// load returns the config for the GraphQL
func (config *graphQL) load() {
	logger.Log.Info("Reading GraphQL config...")
	viper.SetEnvPrefix("graphql")
	viper.AutomaticEnv()

	config.port = viper.GetString("port")
}

// Port will returns graphQL running port
func (config *graphQL) Port() string {
	return config.port
}
