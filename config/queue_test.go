package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type QueueConfigTestSuite struct {
	suite.Suite

	mockViper *viper.Viper
	sut       QueueConfig
}

func (suite *QueueConfigTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *QueueConfigTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.sut = NewQueueConfig(suite.mockViper)
}

func (suite *QueueConfigTestSuite) Test_NatsClientName_ShouldReturnEmpty_IfEnvNotPresent() {
	var emptyNatsClientName string
	os.Setenv("NATS_CLIENT_NAME", "")
	natsClientName := suite.sut.NatsClientName()
	suite.Equal(emptyNatsClientName, natsClientName)
}

func (suite *QueueConfigTestSuite) Test_NatsClientName_ShouldNotReturnEmpty_IfEnvPresent() {
	var envNatsClientName string = "go-cookiecutter"
	os.Setenv("NATS_CLIENT_NAME", envNatsClientName)
	natsClientName := suite.sut.NatsClientName()
	suite.Equal(envNatsClientName, natsClientName)
}

func TestQueueConfig(t *testing.T) {
	suite.Run(t, &QueueConfigTestSuite{})
}
