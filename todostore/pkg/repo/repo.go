package repo

import (
	"gotodo/todostore/pkg/models"
	
)

type TodoRepo struct {
	todos []models.Todo
}

func (r *TodoRepo) GetAll() []models.Todo {
	return r.todos
}

func (r *TodoRepo) Add(todo models.Todo) {
	r.todos = append(r.todos, todo)
}