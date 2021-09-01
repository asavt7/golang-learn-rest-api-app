package services

import "github.com/asavt7/todo/pkg/repos"

type Authorization interface {
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
	return &Service{repo: repo}
}
