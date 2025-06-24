package model

import (
	z "github.com/Oudwins/zog"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var TodoSchema = z.Struct(z.Shape{
	"id":          z.String().Required(z.Message("Please provide a valid ID.")),
	"title":       z.String().Required(z.Message("Please provide a title.")),
	"description": z.String().Max(100, z.Message("Description cannot exceed 100 characters.")),
	"completed":   z.Bool().Default(false),
})

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var CreateTodoRequestSchema = TodoSchema.Pick("title", "description")

type CreateTodoResponse struct {
	Todo Todo `json:"todo"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type GetTodoResponse struct {
	Todo Todo `json:"todo"`
}
