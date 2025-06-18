package handler

import (
	"net/http"

	"example.com/todolist/internal/infrastructure/error_handler"
	"example.com/todolist/internal/infrastructure/validator"
	"example.com/todolist/internal/model"
	"example.com/todolist/internal/service"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (h *TodoHandler) CreateTodo(ctx *gin.Context) {
	var payload model.CreateTodoRequest
	payloadValidator := validator.NewRequestValidator[model.CreateTodoRequest](ctx)
	errorDetail := payloadValidator.Validate(&payload)
	if errorDetail != nil {
		ctx.Error(error_handler.NewHttpException(http.StatusBadRequest, "Invalid request payload", errorDetail))
		return
	}

	todo, err := h.todoService.Create(service.CreateTodo{
		Title:       payload.Title,
		Description: payload.Description})

	if err != nil {
		ctx.Error(error_handler.NewHttpException(http.StatusInternalServerError, "Failed to create todo", nil))
		return
	}

	ctx.JSON(http.StatusCreated, model.CreateTodoResponse{Todo: *todo})
}

func (h *TodoHandler) GetTodoById(ctx *gin.Context) {
	id, hasId := ctx.Params.Get("id")

	if !hasId {
		ctx.Error(error_handler.NewHttpException(http.StatusBadRequest, "Todo ID is required", nil))
		return
	}

	todo, err := h.todoService.GetById(id)
	if err != nil {
		ctx.Error(error_handler.NewHttpException(http.StatusInternalServerError, "Failed to retrieve todo", nil))
		return
	}

	ctx.JSON(http.StatusOK, model.GetTodoResponse{Todo: *todo})
}

func (h *TodoHandler) GetAllTodos(ctx *gin.Context) {
	todos, err := h.todoService.GetAll()
	if err != nil {
		ctx.Error(error_handler.NewHttpException(http.StatusInternalServerError, "Failed to retrieve todos", nil))
		return
	}

	todoVals := make([]model.Todo, len(todos))
	for i, t := range todos {
		if t != nil {
			todoVals[i] = *t
		}
	}
	ctx.JSON(http.StatusOK, model.GetTodosResponse{Todos: todoVals})
}

func (h *TodoHandler) DeleteTodo(ctx *gin.Context) {
	id, hasId := ctx.Params.Get("id")

	if !hasId {
		ctx.Error(error_handler.NewHttpException(http.StatusBadRequest, "Todo ID is required", nil))
		return
	}

	err := h.todoService.Delete(id)
	if err != nil {
		ctx.Error(error_handler.NewHttpException(http.StatusInternalServerError, "Failed to delete todo", nil))
		return
	}

	ctx.Status(http.StatusNoContent)
}
