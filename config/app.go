package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

type App interface {
	BuildEnv() string
}

// app holds the config for the APP
type app struct {
	env *viper.Viper
}

// BuildEnv will returns api running port
func (config *app) BuildEnv() string {
	config.env.AutomaticEnv()
	return config.env.GetString("app_build_env")
}

func NewAppConfig(env *viper.Viper) App {
	logger.Log.Info("Reading API config...")
	return &app{
		env: env,
	}
}
