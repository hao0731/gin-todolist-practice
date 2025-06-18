package service

import (
	"example.com/todolist/internal/model"
	"example.com/todolist/internal/repository"
)

type CreateTodo struct {
	Title       string
	Description string
}

type TodoService interface {
	Create(payload CreateTodo) (*model.Todo, error)
	GetAll() ([]*model.Todo, error)
	GetById(id string) (*model.Todo, error)
	Delete(id string) error
}

type todoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{
		repository: repo,
	}
}

func (s *todoService) Create(payload CreateTodo) (*model.Todo, error) {
	return s.repository.Create(repository.CreateTodo{
		Title:       payload.Title,
		Description: payload.Description,
	})
}

func (s *todoService) GetAll() ([]*model.Todo, error) {
	todos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *todoService) GetById(id string) (*model.Todo, error) {
	todo, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *todoService) Delete(id string) error {
	return s.repository.Delete(id)
}
