package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tech-thinker/go-cookiecutter/app/models"
)

type TodoRepoMock struct {
	mock.Mock
}

func (repo *TodoRepoMock) Save(ctx context.Context, doc *models.Todo) error {
	args := repo.Called(ctx, doc)
	return args.Error(1)
}

func (repo *TodoRepoMock) FindOne(ctx context.Context, doc models.Todo) (models.Todo, error) {
	args := repo.Called(ctx, doc)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (repo *TodoRepoMock) Update(ctx context.Context, doc *models.Todo, fieldsToUpdate []string) error {
	args := repo.Called(ctx, doc, fieldsToUpdate)
	return args.Error(1)
}

func (repo *TodoRepoMock) FindAll(ctx context.Context, query models.TodoQuery) ([]models.Todo, int64, error) {
	args := repo.Called(ctx, query)
	return args.Get(0).([]models.Todo), args.Get(1).(int64), args.Error(2)
}

func NewTodoRepoMock() *TodoRepoMock {
	return &TodoRepoMock{}
}
