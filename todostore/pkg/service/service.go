package service

import (
	"gotodo/todostore/pkg/models"
	"gotodo/todostore/pkg/repo"
)

type TodoService struct {
	TodoRepository *repo.TodoRepo
}

func (s *TodoService) GetAll() []models.Todo {
	return s.TodoRepository.GetAll()
}

func (s *TodoService) Add(todo models.Todo) {
	s.TodoRepository.Add(todo)
}