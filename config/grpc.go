package config

import (
	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

type Grpc interface {
	Port() string
}

// api holds the config for the API
type grpc struct {
	env *viper.Viper
}

// Port will returns api running port
func (config *grpc) Port() string {
	config.env.AutomaticEnv()
	return config.env.GetString("grpc_port")
}

func NewGrpcConfig(env *viper.Viper) Grpc {
	logger.Log.Info("Reading API config...")
	return &grpc{
		env: env,
	}
}
