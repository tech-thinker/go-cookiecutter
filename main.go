package main

import (
	"github.com/mrasif/gomvc/instance"
)

func main() {
	instance.Init()
	defer instance.Destroy()
}
