package todo

import (
	context "context"

	models "github.com/mrasif/gomvc/db/models"
	"github.com/mrasif/gomvc/logger"
	"github.com/mrasif/gomvc/service/initializer"
)

type todo struct {
	dependencies initializer.Services
}

func (s *todo) AddNew(ctx context.Context, message *NewTodo) (*Todo, error) {
	groupError := "ADD_NEW_TODO_GRPC"
	data := models.Todo{
		Task: &message.Task,
	}
	todo, err := s.dependencies.TodoService().Create(data)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, err
	}

	res := &Todo{
		Id:        todo.ID,
		CreatedAt: *todo.CreatedAt,
		UpdatedAt: *todo.UpdatedAt,
		Task:      *todo.Task,
		Done:      todo.Done,
	}
	return res, nil
}

func Init(dependencies initializer.Services) TodoServiceServer {
	return &todo{
		dependencies: dependencies,
	}

}
