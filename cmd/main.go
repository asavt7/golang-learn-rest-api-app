package main

import (
	todo "github.com/asavt7/todo"
	"github.com/asavt7/todo/pkg/handlers"
	"log"
)

func main() {

	s := new(todo.Server)
	handler := new(handlers.Handler)

	err := s.Run("8000", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error while running server %s", err.Error())
		return
	}
}
