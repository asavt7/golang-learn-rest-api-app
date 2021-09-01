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
}

type TodoItem interface {
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
		repo:          repo}
}
