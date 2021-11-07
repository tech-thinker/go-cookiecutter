package service

import (
	"context"

	"github.com/tech-thinker/go-cookiecutter/app/models"
	"github.com/tech-thinker/go-cookiecutter/app/repository"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/vendors"
)

// Todo is interface for todo service
type TodoSvc interface {
	Create(ctx context.Context, doc models.Todo) (models.Todo, error)
	List(ctx context.Context, query models.TodoQuery) ([]models.Todo, error)
}

type todoSvc struct {
	todoRepo        repository.TodoRepo
	modelsValidator vendors.ModelValidator
}

func (svc *todoSvc) Create(ctx context.Context, doc models.Todo) (models.Todo, error) {
	groupError := "CREATE_TODO_SERVICE"
	errs := svc.modelsValidator.Struct(doc)
	if len(errs) > 0 {
		logger.Log.WithError(errs[0]).Error(groupError)
		return doc, errs[0]
	}
	err := svc.todoRepo.Save(ctx, &doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return doc, err
	}
	return doc, nil
}

func (svc *todoSvc) List(ctx context.Context, query models.TodoQuery) ([]models.Todo, error) {
	groupError := "LIST_TODO_SERVICE"
	todos, _, err := svc.todoRepo.FindAll(ctx, query)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return todos, err
	}
	return todos, nil
}

// NewTodo initializes todo service
func NewTodoSvc(
	todoRepo repository.TodoRepo,
	modelValidator vendors.ModelValidator,
) TodoSvc {
	return &todoSvc{
		todoRepo:        todoRepo,
		modelsValidator: modelValidator,
	}
}
