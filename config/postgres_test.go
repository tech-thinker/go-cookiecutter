package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type PostgresConfigTestSuite struct {
	suite.Suite

	mockViper      *viper.Viper
	postgresConfig PostgresConfig
}

func (suite *PostgresConfigTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *PostgresConfigTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.postgresConfig = NewPostgresConfig(suite.mockViper)
}

func (suite *PostgresConfigTestSuite) TestHost_ShouldReturn_HostName() {
	os.Setenv("POSTGRES_HOST", "postgres-go-cookiecutter")

	host := suite.postgresConfig.Host()
	suite.Equal("postgres-go-cookiecutter", host)
}

func (suite *PostgresConfigTestSuite) TestPort_ShouldReturn_PortNumber() {
	os.Setenv("POSTGRES_PORT", "5432")

	port := suite.postgresConfig.Port()
	suite.Equal("5432", port)
}

func (suite *PostgresConfigTestSuite) TestDatabaseName_ShouldReturn_DatabaseName() {
	os.Setenv("POSTGRES_DB", "go-cookiecutter")

	dbName := suite.postgresConfig.Database()
	suite.Equal("go-cookiecutter", dbName)
}

func (suite *PostgresConfigTestSuite) TestUser_ShouldReturn_User() {
	os.Setenv("POSTGRES_USER", "go-cookiecutter")

	user := suite.postgresConfig.User()
	suite.Equal("go-cookiecutter", user)
}

func (suite *PostgresConfigTestSuite) TestPassword_ShouldReturn_Password() {
	os.Setenv("POSTGRES_PASSWORD", "go-cookiecutter123")

	password := suite.postgresConfig.Password()
	suite.Equal("go-cookiecutter123", password)
}

func (suite *PostgresConfigTestSuite) TestURL_ShouldReturn_If_URL_present() {
	url := "postgres://go-cookiecutter:go-cookiecutter123@postgres-go-cookiecutter:5432/go-cookiecutter?sslmode=disable"
	os.Setenv("POSTGRES_URL", url)

	pgUrl := suite.postgresConfig.ConnectionURL()
	suite.Equal(url, pgUrl)
}

func (suite *PostgresConfigTestSuite) TestURL_ShouldReturn_If_URL_Not_present() {
	url := "postgres://go-cookiecutter:go-cookiecutter123@postgres-go-cookiecutter:5432/go-cookiecutter?sslmode=disable"
	os.Setenv("POSTGRES_URL", "")
	os.Setenv("POSTGRES_HOST", "postgres-go-cookiecutter")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "go-cookiecutter")
	os.Setenv("POSTGRES_USER", "go-cookiecutter")
	os.Setenv("POSTGRES_PASSWORD", "go-cookiecutter123")

	pgUrl := suite.postgresConfig.ConnectionURL()
	suite.Equal(url, pgUrl)
}

func TestPostgresConfig(t *testing.T) {
	suite.Run(t, &PostgresConfigTestSuite{})
}
