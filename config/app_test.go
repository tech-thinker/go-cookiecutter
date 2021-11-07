package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type AppConfigTestSuite struct {
	suite.Suite

	mockViper *viper.Viper
	appConfig App
}

func (suite *AppConfigTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *AppConfigTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.appConfig = NewAppConfig(suite.mockViper)
}

func (suite *AppConfigTestSuite) TestPortShouldReturnEmptyIfEnvNotPresent() {
	os.Setenv("API_BUILD_ENV", "")
	buildEnv := suite.appConfig.BuildEnv()
	suite.Empty(buildEnv)
}

func (suite *AppConfigTestSuite) TestPortShouldNotReturnEmptyIfEnvNotPresent() {
	var env string = "test"
	os.Setenv("API_BUILD_ENV", env)
	buildEnv := suite.appConfig.BuildEnv()
	suite.NotEqual(env, buildEnv)
}

func TestAppConfig(t *testing.T) {
	suite.Run(t, &AppConfigTestSuite{})
}
