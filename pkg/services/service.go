package services

import (
	"github.com/asavt7/todo/pkg/domain"
	"github.com/asavt7/todo/pkg/repos"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(id int, input domain.TodoList) (int, error)
	GetAllLists(id int) ([]domain.TodoList, error)
	GetListById(userId, listId int) (domain.TodoList, error)
	Delete(userId, listId int) error
	Update(userId int, listId int, input domain.UpdateTodoListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, input domain.TodoItem) (int, error)
	GetAllItems(userId int, listId int) ([]domain.TodoItem, error)
	GetById(userId int, itemId int) (domain.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input domain.UpdateTodoItem) error
}

type Service struct {
	repo *repos.Repo

	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repos.Repo) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
		repo:          repo}
}
