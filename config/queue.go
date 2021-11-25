package config

import "github.com/spf13/viper"

type QueueConfig interface {
	NatsClientName() string
	NatsURL() string
	NatsUsername() string
	NatsPassword() string
}

// queueConfig holds the config for the queue
type queueConfig struct {
	env *viper.Viper
}

// NatsClientName will returns nats client name
func (config *queueConfig) NatsClientName() string {
	config.env.AutomaticEnv()
	return config.env.GetString("nats_client_name")
}

// NatsURL will returns nats url
func (config *queueConfig) NatsURL() string {
	config.env.AutomaticEnv()
	return config.env.GetString("nats_url")
}

// NatsUsername will returns nats username
func (config *queueConfig) NatsUsername() string {
	config.env.AutomaticEnv()
	return config.env.GetString("nats_username")
}

// NatsPassword will returns nats password
func (config *queueConfig) NatsPassword() string {
	config.env.AutomaticEnv()
	return config.env.GetString("nats_password")
}

func NewQueueConfig(env *viper.Viper) QueueConfig {
	return &queueConfig{
		env: env,
	}
}
