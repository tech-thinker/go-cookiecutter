package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ApiTestSuite struct {
	suite.Suite

	mockViper *viper.Viper
	apiConfig Api
}

func (suite *ApiTestSuite) setupConfig() {
	// os.Setenv("API_BUILD_ENV", "dev")
}

func (suite *ApiTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.apiConfig = NewApiConfig(suite.mockViper)
}

func (suite *ApiTestSuite) TestPortShouldReturnEmptyIfEnvNotPresent() {
	var emptyPort string
	os.Setenv("API_PORT", "")
	port := suite.apiConfig.Port()
	suite.Equal(emptyPort, port)
}

func (suite *ApiTestSuite) TestPortShouldNotReturnEmptyIfEnvNotPresent() {
	var emptyPort string
	os.Setenv("API_PORT", "3000")
	port := suite.apiConfig.Port()
	suite.NotEqual(emptyPort, port)
}

func TestApi(t *testing.T) {
	suite.Run(t, &ApiTestSuite{})
}
