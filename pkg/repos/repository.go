package repos

import "github.com/asavt7/todo/pkg/domain"

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repo struct {
	Authorization
	TodoList
	TodoItem
}
