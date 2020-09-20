package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

// app holds the config for the APP
type app struct {
	buildEnv string
}

// load returns the config for the API
func (config *app) load() {
	logger.Log.Info("Reading API config...")
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	config.buildEnv = viper.GetString("build_env")
}

// Port will returns api running port
func (config *app) Port() string {
	return config.buildEnv
}
