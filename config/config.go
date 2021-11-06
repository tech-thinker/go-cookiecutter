package config

import "github.com/spf13/viper"

type Configuration interface {
	AppConfig() App
	ApiConfig() Api
	PostgresConfig() Postgress
	GrpcConfig() Grpc
}
type configuration struct {
	appConfig  App
	apiConfig  Api
	pgConfig   Postgress
	grpcConfig Grpc
}

// App returns the configuration for application
func (config *configuration) AppConfig() App {
	return config.appConfig
}

// Api returns the configuration for api server
func (config *configuration) ApiConfig() Api {
	return config.apiConfig
}

// Postgres returns the configuration for postgresql database
func (config *configuration) PostgresConfig() Postgress {
	return config.pgConfig
}

// Grpc returns the configuration for grpc service
func (config *configuration) GrpcConfig() Grpc {
	return config.grpcConfig
}

func Init(
	v *viper.Viper,
) Configuration {
	return &configuration{
		appConfig:  NewAppConfig(v),
		apiConfig:  NewApiConfig(v),
		pgConfig:   NewPostgresConfig(v),
		grpcConfig: NewGrpcConfig(v),
	}
}
