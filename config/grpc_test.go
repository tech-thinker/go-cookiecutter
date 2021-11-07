package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type GrpcTestSuite struct {
	suite.Suite

	mockViper  *viper.Viper
	grpcConfig Grpc
}

func (suite *GrpcTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *GrpcTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.grpcConfig = NewGrpcConfig(suite.mockViper)
}

func (suite *GrpcTestSuite) TestPortShouldReturnEmptyIfEnvNotPresent() {
	var emptyPort string
	os.Setenv("API_PORT", "")
	port := suite.grpcConfig.Port()
	suite.Equal(emptyPort, port)
}

func (suite *GrpcTestSuite) TestPortShouldNotReturnEmptyIfEnvNotPresent() {
	var envPort string = "3000"
	os.Setenv("API_PORT", envPort)
	port := suite.grpcConfig.Port()
	suite.NotEqual(envPort, port)
}

func TestGrpcConfig(t *testing.T) {
	suite.Run(t, &GrpcTestSuite{})
}
