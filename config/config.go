package config

type configuration struct {
	appConfig     app
	apiConfig     api
	pgConfig      postgres
	grpcConfig    grpc
	graphqlConfig graphQL
}

var config = &configuration{}

// Load loads the configuration into config object
func Load() {
	config.appConfig.load()
	config.apiConfig.load()
	config.pgConfig.load()
	config.grpcConfig.load()
	config.graphqlConfig.load()
}

// App returns the configuration for application
func App() app {
	return config.appConfig
}

// Api returns the configuration for api server
func Api() api {
	return config.apiConfig
}

// Postgres returns the configuration for postgresql database
func Postgres() postgres {
	return config.pgConfig
}

// Grpc returns the configuration for grpc service
func Grpc() grpc {
	return config.grpcConfig
}

// GraphQL returns the configuration for GraphQL service
func GraphQL() graphQL {
	return config.graphqlConfig
}
