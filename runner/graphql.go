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

// GraphQL is the interface for GraphQL runner
type GraphQL interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type graphQL struct {
}

func (runner *graphQL) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	graphqlConfig := config.GraphQL()
	logger.Log.Infof("Starting GraphQL server on %v...", graphqlConfig.Port())
	services := initializer.Init()

	routerV1 := router.Init(services)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", graphqlConfig.Port()),
		Handler:      routerV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()

}

// NewGraphQL returns an instance of the GraphQL runner
func NewGraphQL() GraphQL {
	return &graphQL{}
}
