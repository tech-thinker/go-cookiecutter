package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/tech-thinker/go-cookiecutter/app/models"
	"github.com/tech-thinker/go-cookiecutter/app/repository"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

// Todo is interface for todo service
type Todo interface {
	Create(ctx context.Context, doc models.Todo) (models.Todo, error)
	List(ctx context.Context, query models.TodoQuery) ([]models.Todo, error)
}

type todo struct {
	todoRepo  repository.TodoRepo
	validator *validator.Validate
}

func (svc *todo) Create(ctx context.Context, doc models.Todo) (models.Todo, error) {
	groupError := "CREATE_TODO_SERVICE"
	err := svc.validator.Struct(doc)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		logger.Log.WithError(errs).Error(groupError)
		return doc, fmt.Errorf(`'%v' is not a valid %v`, errs[0].Value(), errs[0].StructField())
	}
	err = svc.todoRepo.Save(ctx, &doc)
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
	validator *validator.Validate,
) Todo {
	return &todo{
		todoRepo:  todoRepo,
		validator: validator,
	}
}
