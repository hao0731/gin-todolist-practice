package main

import (
	"time"

	"strconv"

	"example.com/todolist/config"
	"example.com/todolist/internal/handler"
	"example.com/todolist/internal/infrastructure/error_handler"
	"example.com/todolist/internal/repository"
	"example.com/todolist/internal/service"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config, err := config.LoadConfig()
	var logger *zap.Logger

	if err != nil {
		logger, _ = zap.NewProduction()
		logger.Error("Failed to load configuration", zap.Error(err))
		return
	}

	if config.Env == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	defer logger.Sync()

	router := gin.New()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(error_handler.ErrorHandler())

	todoRepository := repository.NewInMemoryTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	v1 := router.Group("/api/v1")
	v1.POST("/todos", todoHandler.CreateTodo)
	v1.GET("/todos/:id", todoHandler.GetTodoById)
	v1.GET("/todos", todoHandler.GetAllTodos)
	v1.DELETE("/todos/:id", todoHandler.DeleteTodo)

	router.Run(":" + strconv.Itoa(config.Server.Port))
}
