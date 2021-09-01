package main

import (
	todo "github.com/asavt7/todo"
	"github.com/asavt7/todo/pkg/handlers"
	"github.com/asavt7/todo/pkg/repos"
	"github.com/asavt7/todo/pkg/services"
	"github.com/spf13/viper"
	"log"
)

func main() {

	err := initConfig()
	if err != nil {
		log.Fatalf("error while init configs %s", err.Error())
	}

	s := new(todo.Server)
	repo := repos.NewRepo()
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	err = s.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("error while running server %s", err.Error())
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
