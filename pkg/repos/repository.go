package repos

import "github.com/asavt7/todo/pkg/domain"

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type TodoList interface {
	Create(id int, input domain.TodoList) (int, error)
	GetAllLists(id int) ([]domain.TodoList, error)
	GetListById(userId int, listId int) (domain.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input domain.UpdateTodoListInput) error
}

type TodoItem interface {
}

type Repo struct {
	Authorization
	TodoList
	TodoItem
}
