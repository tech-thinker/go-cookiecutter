package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type GrpcConfigTestSuite struct {
	suite.Suite

	mockViper  *viper.Viper
	grpcConfig GrpcConfig
}

func (suite *GrpcConfigTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *GrpcConfigTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.grpcConfig = NewGrpcConfig(suite.mockViper)
}

func (suite *GrpcConfigTestSuite) TestPort_ShouldReturnEmpty_IfEnvNotPresent() {
	var emptyPort string
	os.Setenv("API_PORT", "")
	port := suite.grpcConfig.Port()
	suite.Equal(emptyPort, port)
}

func (suite *GrpcConfigTestSuite) TestPort_ShouldNotReturnEmpty_IfEnvNotPresent() {
	var envPort string = "4000"
	os.Setenv("API_PORT", envPort)
	port := suite.grpcConfig.Port()
	suite.NotEqual(envPort, port)
}

func TestGrpcConfig(t *testing.T) {
	suite.Run(t, &GrpcConfigTestSuite{})
}
