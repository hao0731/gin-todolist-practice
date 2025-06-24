package handler

import (
	"net/http"

	"example.com/todolist/internal/infrastructure/http_exception"
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
	payload := ctx.MustGet(validator.VALIDATED_PAYLOAD_KEY).(*model.CreateTodoRequest)

	todo, err := h.todoService.Create(service.CreateTodo{
		Title:       payload.Title,
		Description: payload.Description})

	if err != nil {
		exception := http_exception.NewInternalServerErrorException("Failed to create todo", nil)
		ctx.JSON(exception.StatusCode, exception)
		return
	}

	ctx.JSON(http.StatusCreated, model.CreateTodoResponse{Todo: *todo})
}

func (h *TodoHandler) GetTodoById(ctx *gin.Context) {
	id, hasId := ctx.Params.Get("id")

	if !hasId {
		exception := http_exception.NewBadRequestException("Todo ID is required", nil)
		ctx.JSON(http.StatusBadRequest, exception)
		return
	}

	todo, err := h.todoService.GetById(id)
	if err != nil {
		exception := http_exception.NewInternalServerErrorException("Failed to retrieve todo", nil)
		ctx.JSON(exception.StatusCode, exception)
		return
	}

	ctx.JSON(http.StatusOK, model.GetTodoResponse{Todo: *todo})
}

func (h *TodoHandler) GetAllTodos(ctx *gin.Context) {
	todos, err := h.todoService.GetAll()
	if err != nil {
		exception := http_exception.NewInternalServerErrorException("Failed to retrieve todos", nil)
		ctx.JSON(exception.StatusCode, exception)
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
		exception := http_exception.NewBadRequestException("Todo ID is required", nil)
		ctx.JSON(http.StatusBadRequest, exception)
		return
	}

	err := h.todoService.Delete(id)
	if err != nil {
		exception := http_exception.NewInternalServerErrorException("Failed to delete todo", nil)
		ctx.JSON(exception.StatusCode, exception)
		return
	}

	ctx.Status(http.StatusNoContent)
}
