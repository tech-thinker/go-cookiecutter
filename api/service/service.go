package service

import (
	"github.com/mrasif/gomvc/api"
	"github.com/mrasif/gomvc/instance"
	"github.com/mrasif/gomvc/repository"
)

// Services is interface for all service entrypoint
type Services interface {
	TodoService() api.TodoService
}

type services struct {
	todoService api.TodoService
}

func (svc *services) TodoService() api.TodoService {
	return svc.todoService
}

// Init initializes services repo
func Init() Services {
	db := instance.DB()
	return &services{
		todoService: api.NewTodoService(
			repository.NewTodoRepo(db),
		),
	}
}
