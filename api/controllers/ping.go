package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrasif/gomvc/api/service"
)

type PingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct {
	dependencies service.Services
}

func (c *pingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server working fine.",
	})
}

func NewPingController(dependencies service.Services) PingController {
	return &pingController{
		dependencies: dependencies,
	}
}
