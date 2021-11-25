package initializer

import (
	"github.com/tech-thinker/go-cookiecutter/app/repository"
	"github.com/tech-thinker/go-cookiecutter/app/service"
	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/instance"
	"github.com/tech-thinker/go-cookiecutter/vendors"
)

// Services is interface for all service entrypoint
type Services interface {
	TodoService() service.TodoSvc
	EventRepo() repository.EventRepo
}

type services struct {
	todoService service.TodoSvc
	eventRepo   repository.EventRepo
}

func (svc *services) TodoService() service.TodoSvc {
	return svc.todoService
}

func (svc *services) EventRepo() repository.EventRepo {
	return svc.eventRepo
}

// Init initializes services repo
func Init(config config.Configuration, instance instance.Instance) Services {
	db := instance.DB()
	validation := instance.Validator()
	modelValidator := vendors.NewModelValidator(validation)

	eventRepo := repository.NewEventRepo(config, instance)

	return &services{
		todoService: service.NewTodoSvc(
			repository.NewTodoRepo(db),
			modelValidator,
		),
		eventRepo: eventRepo,
	}
}
