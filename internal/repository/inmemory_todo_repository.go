package repository

import (
	"example.com/todolist/internal/model"
	"github.com/google/uuid"
)

type InMemoryTodoRepository struct {
	todos []*model.Todo
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{todos: []*model.Todo{}}
}

func (r *InMemoryTodoRepository) Create(data CreateTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Id:          uuid.NewString(),
		Title:       data.Title,
		Description: data.Description,
		Completed:   false,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *InMemoryTodoRepository) GetAll() ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *InMemoryTodoRepository) GetById(id string) (*model.Todo, error) {
	for _, todo := range r.todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return nil, nil
}

func (r *InMemoryTodoRepository) Delete(id string) error {
	for i, todo := range r.todos {
		if todo.Id == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return nil
}
