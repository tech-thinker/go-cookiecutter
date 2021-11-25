package repository

import (
	"context"
	"encoding/json"

	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/constants"
	"github.com/tech-thinker/go-cookiecutter/instance"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/queue"
)

type EventRepo interface {
	Publish(ctx context.Context, subject string, data interface{}) error
}

type eventRepo struct {
	config   config.Configuration
	instance instance.Instance
}

// Publish is used to send message to the queue
// ctx: context.Context
// subject: constants string
// data: interface{} is the data to be sent to the queue
func (repo *eventRepo) Publish(ctx context.Context, subject string, data interface{}) error {
	var (
		groupError string = "PUBLISH_EVENT"
		message    queue.Message
		err        error
		byteData   []byte
	)

	byteData, err = json.Marshal(data)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	message = queue.Message{
		Data: byteData,
	}

	// Prepare Queue before sending message
	queue, err := queue.NewQueue(
		constants.EventStream,
		subject,
		repo.config,
		repo.instance,
	)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	// Sending message
	err = queue.Publish(ctx, &message)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	return nil
}

func NewEventRepo(
	config config.Configuration,
	instance instance.Instance,
) EventRepo {
	return &eventRepo{
		config:   config,
		instance: instance,
	}
}
