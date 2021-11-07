package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

type AppConfig interface {
	BuildEnv() string
}

// app holds the config for the APP
type appConfig struct {
	env *viper.Viper
}

// BuildEnv will returns api running port
func (config *appConfig) BuildEnv() string {
	config.env.AutomaticEnv()
	return config.env.GetString("app_build_env")
}

func NewAppConfig(env *viper.Viper) AppConfig {
	logger.Log.Info("Reading APP config...")
	return &appConfig{
		env: env,
	}
}
