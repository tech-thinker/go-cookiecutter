package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/go-cookiecutter/app/initializer"
	"github.com/tech-thinker/go-cookiecutter/app/models"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

// TodoController is an interface for todo controller
type TodoController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
}

type todoController struct {
	dependencies initializer.Services
}

func (c *todoController) Create(ctx *gin.Context) {
	groupError := "CREATE_TODO_CONTROLLER"

	var data models.Todo
	ctx.BindJSON(&data)
	todo, err := c.dependencies.TodoService().Create(context.Background(), data)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    todo,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Created",
		"data":    todo,
	})
}

func (c *todoController) List(ctx *gin.Context) {
	groupError := "LIST_TODO_CONTROLLER"

	query := models.TodoQuery{}

	todos, err := c.dependencies.TodoService().List(context.Background(), query)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    todos,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OK",
		"data":    todos,
	})
}

// NewTodoController initializes todo controller with dependency
func NewTodoController(dependencies initializer.Services) TodoController {
	return &todoController{
		dependencies: dependencies,
	}
}
