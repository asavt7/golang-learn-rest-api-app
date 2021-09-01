package main

import (
	todo "github.com/asavt7/todo"
	"github.com/asavt7/todo/pkg/handlers"
	"github.com/asavt7/todo/pkg/repos"
	"github.com/asavt7/todo/pkg/services"
	"log"
)

func main() {

	s := new(todo.Server)
	repo := repos.NewRepo()
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	err := s.Run("8000", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error while running server %s", err.Error())
		return
	}
}
