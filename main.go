package main

import (
	"github.com/mrasif/gomvc/api/router"
	"github.com/mrasif/gomvc/api/service"
	"github.com/mrasif/gomvc/config"
	"github.com/mrasif/gomvc/instance"
	"github.com/mrasif/gomvc/logger"
)

func main() {
	config.Load()
	instance.Init()
	defer instance.Destroy()

	dependencies := service.Init()
	err := router.Init(dependencies).Run(":3000")
	if err != nil {
		logger.Log.Fatal(err)
	}
}
