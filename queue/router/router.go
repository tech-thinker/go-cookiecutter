package router

import (
	"context"

	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/queue/jobs"
)

func Init(ctx context.Context, jobs jobs.Jobs, shutDownChannel chan *bool) error {
	logger.Log.Info("Initializing router")

	err := jobs.PingWorker().Invoke(ctx, shutDownChannel)
	if err != nil {
		return err
	}

	return nil
}
