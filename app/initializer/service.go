package initializer

import (
	"github.com/tech-thinker/go-cookiecutter/app/repository"
	"github.com/tech-thinker/go-cookiecutter/app/service"
	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/instance"
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
func Init(config config.Configuration, instance instance.Instance) Services {
	db := instance.DB()
	validator := instance.Validator()

	return &services{
		todoService: service.NewTodo(
			repository.NewTodoRepo(db),
			validator,
		),
	}
}
