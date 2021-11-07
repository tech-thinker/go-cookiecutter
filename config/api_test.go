package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ApiConfigTestSuite struct {
	suite.Suite

	mockViper *viper.Viper
	apiConfig ApiConfig
}

func (suite *ApiConfigTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *ApiConfigTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.apiConfig = NewApiConfig(suite.mockViper)
}

func (suite *ApiConfigTestSuite) TestPort_ShouldReturnEmpty_IfEnvNotPresent() {
	var emptyPort string
	os.Setenv("API_PORT", "")
	port := suite.apiConfig.Port()
	suite.Equal(emptyPort, port)
}

func (suite *ApiConfigTestSuite) TestPort_ShouldNotReturnEmpty_IfEnvNotPresent() {
	var emptyPort string
	os.Setenv("API_PORT", "3000")
	port := suite.apiConfig.Port()
	suite.NotEqual(emptyPort, port)
}

func TestApiConfig(t *testing.T) {
	suite.Run(t, new(ApiConfigTestSuite))
}
