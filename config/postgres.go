package config

import (
	"fmt"

	"github.com/mrasif/gomvc/logger"
	"github.com/spf13/viper"
)

// postgres is the config object for postgres database
type postgres struct {
	host     string
	port     string
	user     string
	password string
	database string
	url      string
}

// load method loads configuration
func (config *postgres) load() {
	logger.Log.Info("Reading postgresql database configuration...")
	viper.SetEnvPrefix("postgres")
	viper.AutomaticEnv()

	config.host = viper.GetString("host")
	config.port = viper.GetString("port")
	config.database = viper.GetString("db")
	config.user = viper.GetString("user")
	config.password = viper.GetString("password")
	config.url = viper.GetString("url")
}

// Host returns database hostname
func (config *postgres) Host() string {
	return config.host
}

// Port returns database port
func (config *postgres) Port() string {
	return config.port
}

// Database returns database name
func (config *postgres) Database() string {
	return config.database
}

// User returns database username
func (config *postgres) User() string {
	return config.user
}

// Password returns database password
func (config *postgres) Password() string {
	return config.password
}

// ConnectionURL returns connection url for postgresql database
func (config *postgres) ConnectionURL() string {
	if len(config.url) > 0 {
		return config.url
	}
	return fmt.Sprintf(`postgres://%v:%v@%v:%v/%v?sslmode=disable`, config.user, config.password, config.host, config.port, config.database)
}
