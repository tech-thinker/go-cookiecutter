package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/viper"
	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/instance"
	"github.com/tech-thinker/go-cookiecutter/runner"

	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

func main() {
	v := viper.New()
	config := config.Init(v)

	var shutDownChannel chan *bool

	instance := instance.Init(config)
	defer instance.Destroy()

	clientApp := cli.NewApp()
	clientApp.Name = "go-cookiecutter"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the service",
			Action: func(c *cli.Context) error {
				ctx := context.Background()

				var wg sync.WaitGroup

				wg.Add(1)
				go runner.NewAPI(config, instance).Go(ctx, &wg)

				wg.Add(1)
				go runner.NewGRPC(config, instance).Go(ctx, &wg)

				wg.Wait()
				return nil

			},
		},
		{
			Name:  "start_workers",
			Usage: "Start the workers",
			Action: func(c *cli.Context) error {
				ctx := context.Background()

				var wg sync.WaitGroup

				wg.Add(1)
				go runner.NewWorker(config, instance).Go(ctx, shutDownChannel, &wg)
				wg.Wait()

				return nil
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(
		signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		<-signalChannel
		shutDown(shutDownChannel)
	}()

}

func shutDown(shutDownChannel chan *bool) {
	shutDown := true
	shutDownChannel <- &shutDown
	close(shutDownChannel)
}
