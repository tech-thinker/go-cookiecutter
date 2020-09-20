package todo

import (
	context "context"

	models "github.com/tech-thinker/go-cookiecutter/db/models"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/service/initializer"
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
		CreatedAt: todo.CreatedAt.Unix(),
		UpdatedAt: todo.UpdatedAt.Unix(),
		Task:      *todo.Task,
		Done:      todo.Done,
	}
	return res, nil
}

func (s *todo) List(input *TodoListInput, list TodoService_ListServer) error {
	groupError := "LIST_TODO_GRPC"
	query := models.TodoQuery{}
	todos, err :=s.dependencies.TodoService().List(query)
	if err!=nil {
		logger.Log.WithError(err).Error(groupError)
	}
	var res []*Todo
	for _, t := range todos {
		r := &Todo{
			Id:        t.ID,
			CreatedAt: t.CreatedAt.Unix(),
			UpdatedAt: t.UpdatedAt.Unix(),
			Task:      *t.Task,
			Done:      t.Done,
		}
		res=append(res,r)
	}
	listRes := &ListResponse{
		Items: res,
	}
	list.Send(listRes)
	return nil
}

func Init(dependencies initializer.Services) TodoServiceServer {
	return &todo{
		dependencies: dependencies,
	}

}
