package main

import (
	"github.com/tech-thinker/go-cookiecutter/api/router"
	"github.com/tech-thinker/go-cookiecutter/config"
	"github.com/tech-thinker/go-cookiecutter/instance"
	"github.com/tech-thinker/go-cookiecutter/logger"
	"github.com/tech-thinker/go-cookiecutter/service/initializer"
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
