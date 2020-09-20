package service

import (
	"github.com/tech-thinker/go-cookiecutter/db/models"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/repository"
)

// TodoService is interface for todoService
type Todo interface {
	Create(doc models.Todo) (models.Todo, error)
	List(query models.TodoQuery) ([]models.Todo, error)
}

type todo struct {
	todoRepo repository.TodoRepo
}

func (svc *todo) Create(doc models.Todo) (models.Todo, error) {
	groupError := "CREATE_TODO_SERVICE"
	err := svc.todoRepo.Save(&doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return doc, err
	}
	return doc, nil
}

func (svc *todo) List(query models.TodoQuery) ([]models.Todo, error) {
	groupError := "LIST_TODO_SERVICE"
	todos, _, err := svc.todoRepo.FindAll(query)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return todos, err
	}
	return todos, nil
}

// NewTodoService initializes todoService
func NewTodo(
	todoRepo repository.TodoRepo,
) Todo {
	return &todo{
		todoRepo: todoRepo,
	}
}
