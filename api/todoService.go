package api

import (
	"github.com/mrasif/gomvc/db/models"
	"github.com/mrasif/gomvc/logger"
	"github.com/mrasif/gomvc/repository"
)

// TodoService is interface for todoService
type TodoService interface {
	Create(doc models.Todo) (models.Todo, error)
	List(query models.TodoQuery) ([]models.Todo, error)
}

type todoService struct {
	todoRepo repository.TodoRepo
}

func (svc *todoService) Create(doc models.Todo) (models.Todo, error) {
	groupError := "CREATE_TODO_SERVICE"
	err := svc.todoRepo.Save(&doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return doc, err
	}
	return doc, nil
}

func (svc *todoService) List(query models.TodoQuery) ([]models.Todo, error) {
	groupError := "LIST_TODO_SERVICE"
	todos, _, err := svc.todoRepo.FindAll(query)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return todos, err
	}
	return todos, nil
}

// NewTodoService initializes todoService
func NewTodoService(
	todoRepo repository.TodoRepo,
) TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}
