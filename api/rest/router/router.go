package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/go-cookiecutter/api/rest/controllers"
	"github.com/tech-thinker/go-cookiecutter/app/initializer"
)

// Init sets router
func Init(dependencies initializer.Services) *gin.Engine {
	router := gin.Default()

	ping := controllers.NewPingController(dependencies)
	todo := controllers.NewTodoController(dependencies)
	router.GET("/ping", ping.Ping)
	router.GET("/todos", todo.List)
	router.POST("/todos", todo.Create)

	return router
}
