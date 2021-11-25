package workers

import (
	"context"

	"github.com/astaxie/beego/orm"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/queue"

	"github.com/tech-thinker/go-cookiecutter/vendors"
)

type PingWorker interface {
	Invoke(ctx context.Context, shutDownChannel chan *bool) error
}

type pingWorker struct {
	eventChannel chan *queue.Message
	eventQueue   queue.Queue

	db             orm.Ormer
	modelValidator vendors.ModelValidator
}

func (w *pingWorker) Invoke(ctx context.Context, shutDownChannel chan *bool) error {
	var (
		groupError string = "INVOKE_PING_WORKER"
	)

	w.eventChannel = make(chan *queue.Message)

	err := w.eventQueue.Subscribe(ctx, w.eventChannel)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	go func(ctx context.Context, ch chan *queue.Message, shutDownChannel chan *bool) {
		for {
			select {
			case message := <-ch:
				logger.Log.Infof("An event has popped in to the message send queue...")

				err := w.handleOnPingReceive(ctx, message.Data)
				if err != nil {
					logger.Log.WithError(err).Error(groupError)
				}
				err = message.Ack(ctx)
				if err != nil {
					logger.Log.WithError(err).Error(groupError)
				}
			case shutDown := <-shutDownChannel:
				if *shutDown {
					logger.Log.Info("Shutting down message sending handler...")
					return
				}
			}
		}
	}(ctx, w.eventChannel, shutDownChannel)

	logger.Log.Info("Intialized message sending handler...")
	return nil
}

func (w *pingWorker) handleOnPingReceive(ctx context.Context, data []byte) error {
	logger.Log.Infof("Handling ping event...")
	logger.Log.Infof("Data: %v", string(data))
	return nil
}

func NewPingWorker(
	eventQueue queue.Queue,
	db orm.Ormer,
	modelValidator vendors.ModelValidator,
) PingWorker {
	return &pingWorker{
		eventQueue:     eventQueue,
		db:             db,
		modelValidator: modelValidator,
	}
}
