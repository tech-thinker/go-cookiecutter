package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

type Api interface {
	Port() string
}

// api holds the config for the API
type api struct {
	env *viper.Viper
}

// Port will returns api running port
func (config *api) Port() string {
	config.env.AutomaticEnv()
	port := config.env.GetString("api_port")
	return port
}

func NewApiConfig(env *viper.Viper) Api {
	logger.Log.Info("Reading API config...")
	return &api{
		env: env,
	}
}
