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
	"Id":          z.String().UUID().Required(z.Message("Please provide a valid ID.")),
	"Title":       z.String().Required(z.Message("Please provide a title.")),
	"Description": z.String().Max(100, z.Message("Description cannot exceed 100 characters.")).Optional(),
	"Completed":   z.Bool().Default(false).Optional(),
})

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var CreateTodoRequestSchema = TodoSchema.Pick("Title", "Description")

type CreateTodoResponse struct {
	Todo Todo `json:"todo"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type GetTodoResponse struct {
	Todo Todo `json:"todo"`
}
