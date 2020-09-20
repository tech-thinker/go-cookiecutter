package main

import (
	"context"
	"os"
	"sync"

	"github.com/mrasif/gomvc/config"
	"github.com/mrasif/gomvc/instance"
	"github.com/mrasif/gomvc/runner"
	"github.com/urfave/cli"
)

func main() {
	config.Load()
	instance.Init()
	defer instance.Destroy()

	clientApp := cli.NewApp()
	clientApp.Name = "auth-app"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the service",
			Action: func(c *cli.Context) error {
				ctx := context.Background()

				var wg sync.WaitGroup
				wg.Add(1)

				go runner.NewAPI().Go(ctx, &wg)

				wg.Wait()
				return nil

			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)

	}

}
