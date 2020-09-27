package runner

import (
	"fmt"
	"context"
	"sync"

	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/grpc/router"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/service/initializer"
)

// GRPC is the interface for gRPC runner
type GRPC interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type grpc struct {
}

func (runner *grpc) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	grpcConfig := config.Grpc()
	logger.Log.Infof("Starting gRPC server on %v...", grpcConfig.Port())
	services := initializer.Init()

	err := router.Init(services, fmt.Sprintf(":%s",grpcConfig.Port()))
	if err != nil {
		logger.Log.WithError(err).Error("GRPC Runner")
	}
}

// NewGRPC returns an instance of the gRPC runner
func NewGRPC() GRPC {
	return &grpc{}
}
