package services

import (
	"github.com/asavt7/todo/pkg/domain"
	"github.com/asavt7/todo/pkg/repos"
)

type TodoListService struct {
	repo repos.TodoList
}

func NewTodoListService(repo repos.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(id int, input domain.TodoList) (int, error) {
	return t.repo.Create(id, input)
}
