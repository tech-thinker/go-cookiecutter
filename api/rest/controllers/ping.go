package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/go-cookiecutter/app/initializer"
)

// PingController is an interface for ping controller
type PingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct {
	dependencies initializer.Services
}

func (c *pingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server working fine.",
	})
}

// NewPingController initializes ping controller with dependency
func NewPingController(dependencies initializer.Services) PingController {
	return &pingController{
		dependencies: dependencies,
	}
}
