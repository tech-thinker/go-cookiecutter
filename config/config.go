package config

import "github.com/spf13/viper"

type Configuration interface {
	AppConfig() AppConfig
	ApiConfig() ApiConfig
	PostgresConfig() PostgresConfig
	GrpcConfig() GrpcConfig
	QueueConfig() QueueConfig
}
type configuration struct {
	appConfig   AppConfig
	apiConfig   ApiConfig
	pgConfig    PostgresConfig
	grpcConfig  GrpcConfig
	queueConfig QueueConfig
}

// App returns the configuration for application
func (config *configuration) AppConfig() AppConfig {
	return config.appConfig
}

// Api returns the configuration for api server
func (config *configuration) ApiConfig() ApiConfig {
	return config.apiConfig
}

// Postgres returns the configuration for postgresql database
func (config *configuration) PostgresConfig() PostgresConfig {
	return config.pgConfig
}

// Grpc returns the configuration for grpc service
func (config *configuration) GrpcConfig() GrpcConfig {
	return config.grpcConfig
}

// Queue returns the configuration for queue service
func (config *configuration) QueueConfig() QueueConfig {
	return config.queueConfig
}

func Init(
	v *viper.Viper,
) Configuration {
	return &configuration{
		appConfig:   NewAppConfig(v),
		apiConfig:   NewApiConfig(v),
		pgConfig:    NewPostgresConfig(v),
		grpcConfig:  NewGrpcConfig(v),
		queueConfig: NewQueueConfig(v),
	}
}
