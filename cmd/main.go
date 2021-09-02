package main

import (
	todo "github.com/asavt7/todo"
	"github.com/asavt7/todo/pkg/handlers"
	"github.com/asavt7/todo/pkg/repos"
	"github.com/asavt7/todo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(new(logrus.JSONFormatter))
	err := initConfig()
	if err != nil {
		logrus.Fatalf("error while init configs %s", err.Error())
	}

	db, err := repos.NewPostgreDb(repos.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.ssl.mode"),
	})
	if err != nil {
		logrus.Fatalf("error init database %s", err.Error())
	}

	s := new(todo.Server)
	repo := repos.NewPostgresRepo(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	err = s.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		logrus.Fatalf("error while running server %s", err.Error())
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
