package runner

import (
	"context"
	"sync"

	"github.com/mrasif/gomvc/config"
	"github.com/mrasif/gomvc/grpc/router"
	"github.com/mrasif/gomvc/logger"
	"github.com/mrasif/gomvc/service/initializer"
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

	err := router.Init(services, grpcConfig.Port())
	if err != nil {
		logger.Log.WithError(err).Error("GRPC Runner")
	}
}

// NewGRPC returns an instance of the gRPC runner
func NewGRPC() API {
	return &api{}
}
