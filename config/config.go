package config

type configuration struct {
	pgConfig postgres
}

var config = &configuration{}

// Load loads the configuration into config object
func Load() {
	config.pgConfig.load()
}

// Postgres returns the configuration for postgresql database
func Postgres() postgres {
	return config.pgConfig
}
