package service

import (
	"context"
	"os"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
	"github.com/tech-thinker/go-cookiecutter/app/models"
	"github.com/tech-thinker/go-cookiecutter/app/repository"
	"github.com/tech-thinker/go-cookiecutter/vendors"
)

type TodoSvcTestSuite struct {
	suite.Suite
	ctx context.Context
	svc TodoSvc

	mockValidator vendors.ModelValidatorMock
	mockTodoRepo  *repository.TodoRepoMock
}

func (suite *TodoSvcTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *TodoSvcTestSuite) SetupTest() {
	suite.setupConfig()
	suite.ctx = context.Background()

	suite.mockValidator = vendors.NewModelValidatorMock()

	suite.mockTodoRepo = new(repository.TodoRepoMock)

	suite.svc = NewTodoSvc(suite.mockTodoRepo, &suite.mockValidator)
}

func (suite *TodoSvcTestSuite) TestCreate_Success() {
	task := fake.Title()
	todo := models.Todo{
		Task: &task,
	}

	suite.mockValidator.On("Struct", todo).Return(nil)
	suite.mockTodoRepo.On("Save", suite.ctx, &todo).Return(todo, nil)

	result, err := suite.svc.Create(suite.ctx, todo)
	suite.NoError(err)
	suite.Equal(todo, result)
}

func TestTodoSvc(t *testing.T) {
	suite.Run(t, new(TodoSvcTestSuite))
}
