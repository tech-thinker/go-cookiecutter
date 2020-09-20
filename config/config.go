package config

type configuration struct {
	apiConfig api
	pgConfig  postgres
}

var config = &configuration{}

// Load loads the configuration into config object
func Load() {
	config.apiConfig.load()
	config.pgConfig.load()
}

// Api returns the configuration for api server
func Api() api {
	return config.apiConfig
}

// Postgres returns the configuration for postgresql database
func Postgres() postgres {
	return config.pgConfig
}
