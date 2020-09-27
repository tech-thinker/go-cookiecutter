package service

import (
	"context"

	"github.com/tech-thinker/go-cookiecutter/db/models"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/repository"
)

// Todo is interface for todo service
type Todo interface {
	Create(ctx context.Context, doc models.Todo) (models.Todo, error)
	List(ctx context.Context, query models.TodoQuery) ([]models.Todo, error)
}

type todo struct {
	todoRepo repository.TodoRepo
}

func (svc *todo) Create(ctx context.Context, doc models.Todo) (models.Todo, error) {
	groupError := "CREATE_TODO_SERVICE"
	err := svc.todoRepo.Save(ctx, &doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return doc, err
	}
	return doc, nil
}

func (svc *todo) List(ctx context.Context, query models.TodoQuery) ([]models.Todo, error) {
	groupError := "LIST_TODO_SERVICE"
	todos, _, err := svc.todoRepo.FindAll(ctx, query)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return todos, err
	}
	return todos, nil
}

// NewTodo initializes todo service
func NewTodo(
	todoRepo repository.TodoRepo,
) Todo {
	return &todo{
		todoRepo: todoRepo,
	}
}
