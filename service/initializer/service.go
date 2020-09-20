package initializer

import (
	"github.com/tech-thinker/go-cookiecutter/instance"
	"github.com/tech-thinker/go-cookiecutter/repository"
	"github.com/tech-thinker/go-cookiecutter/service"
)

// Services is interface for all service entrypoint
type Services interface {
	TodoService() service.Todo
}

type services struct {
	todoService service.Todo
}

func (svc *services) TodoService() service.Todo {
	return svc.todoService
}

// Init initializes services repo
func Init() Services {
	db := instance.DB()
	return &services{
		todoService: service.NewTodo(
			repository.NewTodoRepo(db),
		),
	}
}
