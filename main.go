package main

import (
	"github.com/mrasif/gomvc/api/router"
	"github.com/mrasif/gomvc/config"
	"github.com/mrasif/gomvc/instance"
	"github.com/mrasif/gomvc/logger"
	"github.com/mrasif/gomvc/service/initializer"
)

func main() {
	config.Load()
	instance.Init()
	defer instance.Destroy()

	dependencies := initializer.Init()
	err := router.Init(dependencies).Run(":3000")
	if err != nil {
		logger.Log.Fatal(err)
	}
}
