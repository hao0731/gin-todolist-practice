package main

import (
	"net/http"

	"example.com/todolist/pkg/error_handler"
	"example.com/todolist/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,max=20"`
	Description string `json:"description"`
}

var todos = []Todo{}

func createTodoHandler(ctx *gin.Context) {
	var payload CreateTodoRequest
	payloadValidator := validator.New[CreateTodoRequest](ctx)
	errorDetail := payloadValidator.Validate(&payload)
	if errorDetail != nil {
		ctx.Error(error_handler.NewHttpException(http.StatusBadRequest, "Invalid request payload", errorDetail))
		return
	}
	todo := Todo{
		Id:          uuid.NewString(),
		Title:       payload.Title,
		Description: payload.Description,
		Completed:   false,
	}
	todos = append(todos, todo)
	ctx.JSON(http.StatusCreated, gin.H{
		"todo": todo,
	})
}

func getTodosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func main() {
	router := gin.Default()
	router.Use(error_handler.ErrorHandler())

	v1 := router.Group("/api/v1")
	v1.GET("/todos", getTodosHandler)
	v1.POST("/todos", createTodoHandler)

	router.Run()
}
