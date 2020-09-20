package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

// api holds the config for the API
type api struct {
	port string
}

// load returns the config for the API
func (config *api) load() {
	logger.Log.Info("Reading API config...")
	viper.SetEnvPrefix("api")
	viper.AutomaticEnv()

	config.port = viper.GetString("port")
}

// Port will returns api running port
func (config *api) Port() string {
	return config.port
}
