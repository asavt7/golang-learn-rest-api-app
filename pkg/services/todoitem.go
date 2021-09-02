package services

import (
	"github.com/asavt7/todo/pkg/domain"
	"github.com/asavt7/todo/pkg/repos"
)

type TodoItemService struct {
	repo repos.TodoItem
	listRepo repos.TodoList
}

func NewTodoItemService(repo repos.TodoItem, listRepo repos.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (t *TodoItemService) Create(userId int, listId int, input domain.TodoItem) (int, error) {
	_, err := t.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	
	return t.repo.Create(listId,input)
}





