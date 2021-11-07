package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

type ApiConfig interface {
	Port() string
}

// api holds the config for the API
type apiConfig struct {
	env *viper.Viper
}

// Port will returns api running port
func (config *apiConfig) Port() string {
	config.env.AutomaticEnv()
	port := config.env.GetString("api_port")
	return port
}

func NewApiConfig(env *viper.Viper) ApiConfig {
	logger.Log.Info("Reading API config...")
	return &apiConfig{
		env: env,
	}
}
