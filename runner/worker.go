package runner

import (
	"context"
	"log"
	"sync"

	"github.com/tech-thinker/go-cookiecutter/app/initializer"
	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/instance"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/queue/jobs"
	"github.com/tech-thinker/go-cookiecutter/queue/router"
)

type Worker interface {
	Go(ctx context.Context, shutDownChannel chan *bool, wg *sync.WaitGroup)
}

type worker struct {
	config   config.Configuration
	instance instance.Instance
}

func (runner *worker) Go(ctx context.Context, shutDownChannel chan *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	logger.Log.Info("Starting worker...")

	// init services
	services := initializer.Init(runner.config, runner.instance)
	jobs, err := jobs.Init(runner.config, runner.instance, services)
	if err != nil {
		log.Fatal(err)
	}

	// init worker handlers
	err = router.Init(ctx, jobs, shutDownChannel)
	if err != nil {
		log.Fatal(err)
	}

	for {
		shutDown := <-shutDownChannel
		if *shutDown {
			return
		}
	}
}

func NewWorker(
	config config.Configuration,
	instance instance.Instance,
) Worker {
	return &worker{
		config:   config,
		instance: instance,
	}
}
