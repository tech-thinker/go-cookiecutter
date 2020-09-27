package runner

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/tech-thinker/go-cookiecutter/api/router"
	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/service/initializer"
)

// API is the interface for rest api runner
type API interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type api struct {
}

func (runner *api) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	apiConfig := config.Api()
	logger.Log.Infof("Starting API server on %v...", apiConfig.Port())
	services := initializer.Init()

	routerV1 := router.Init(services)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", apiConfig.Port()),
		Handler:      routerV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()

}

// NewAPI returns an instance of the REST API runner
func NewAPI() API {
	return &api{}
}
