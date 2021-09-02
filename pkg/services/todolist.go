package services

import (
	"github.com/asavt7/todo/pkg/domain"
	"github.com/asavt7/todo/pkg/repos"
)

type TodoListService struct {
	repo repos.TodoList
}

func (t *TodoListService) Update(userId int, listId int, input domain.UpdateTodoListInput) error {
	return t.repo.Update(userId, listId, input)
}

func (t *TodoListService) Delete(userId, listId int) error {
	return t.repo.Delete(userId, listId)
}

func (t *TodoListService) GetListById(userId, listId int) (domain.TodoList, error) {
	return t.repo.GetListById(userId, listId)
}

func (t *TodoListService) GetAllLists(id int) ([]domain.TodoList, error) {
	return t.repo.GetAllLists(id)
}

func NewTodoListService(repo repos.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(id int, input domain.TodoList) (int, error) {
	return t.repo.Create(id, input)
}
