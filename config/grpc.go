package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

// api holds the config for the API
type grpc struct {
	port string
}

// load returns the config for the API
func (config *grpc) load() {
	logger.Log.Info("Reading API config...")
	viper.SetEnvPrefix("grpc")
	viper.AutomaticEnv()

	config.port = viper.GetString("port")
}

// Port will returns api running port
func (config *grpc) Port() string {
	return config.port
}
