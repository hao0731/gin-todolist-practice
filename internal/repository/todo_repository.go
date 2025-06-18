package repository

import "example.com/todolist/internal/model"

type CreateTodo struct {
	Title       string
	Description string
}

type TodoRepository interface {
	Create(data CreateTodo) (*model.Todo, error)
	GetAll() ([]*model.Todo, error)
	GetById(id string) (*model.Todo, error)
	Delete(id string) error
}
