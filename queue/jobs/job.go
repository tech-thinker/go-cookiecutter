package jobs

import (
	"github.com/tech-thinker/go-cookiecutter/app/initializer"
	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/constants"
	"github.com/tech-thinker/go-cookiecutter/instance"
	"github.com/tech-thinker/go-cookiecutter/queue"
	"github.com/tech-thinker/go-cookiecutter/queue/workers"
	"github.com/tech-thinker/go-cookiecutter/vendors"
)

type Jobs interface {
	PingWorker() workers.PingWorker
}

type jobs struct {
	pingWorker workers.PingWorker
}

func (j *jobs) PingWorker() workers.PingWorker {
	return j.pingWorker
}

func Init(
	config config.Configuration,
	instance instance.Instance,
	services initializer.Services,
) (Jobs, error) {
	db := instance.DB()
	validation := instance.Validator()

	modelValidator := vendors.NewModelValidator(validation)

	onPingReceiveQueue, err := queue.NewQueue(
		constants.EventStream,
		constants.NatsOnPingReceived,
		config,
		instance,
	)
	if err != nil {
		return nil, err
	}
	pingWorker := workers.NewPingWorker(
		onPingReceiveQueue,
		db,
		modelValidator,
	)

	return &jobs{
		pingWorker: pingWorker,
	}, nil
}
