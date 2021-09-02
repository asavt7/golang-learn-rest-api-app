package services

import (
	"github.com/asavt7/todo/pkg/domain"
	"github.com/asavt7/todo/pkg/repos"
)

type TodoItemService struct {
	repo     repos.TodoItem
	listRepo repos.TodoList
}

func (t *TodoItemService) Update(userId int, itemId int, input domain.UpdateTodoItem) error {
	return t.repo.Update(userId, itemId, input)
}

func (t *TodoItemService) Delete(userId int, itemId int) error {
	return t.repo.Delete(userId, itemId)
}

func (t *TodoItemService) GetById(userId int, itemId int) (domain.TodoItem, error) {
	return t.repo.GetById(userId, itemId)
}

func (t *TodoItemService) GetAllItems(userId int, listId int) ([]domain.TodoItem, error) {
	return t.repo.GetAllItems(userId, listId)

}

func NewTodoItemService(repo repos.TodoItem, listRepo repos.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (t *TodoItemService) Create(userId int, listId int, input domain.TodoItem) (int, error) {
	_, err := t.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return t.repo.Create(listId, input)
}
