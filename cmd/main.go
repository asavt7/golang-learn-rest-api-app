package main

import (
	"context"
	todo "github.com/asavt7/todo"
	"github.com/asavt7/todo/pkg/handlers"
	"github.com/asavt7/todo/pkg/repos"
	"github.com/asavt7/todo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		err = s.Run(viper.GetString("port"), handler.InitRoutes())
		if err != nil {
			logrus.Fatalf("error while running server %s", err.Error())
			return
		}
	}()

	logrus.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Server shutting down")

	err = s.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("error occurred while server shutting down : %s", err.Error())
	}
	logrus.Info("Server shut down")

	err = db.Close()
	if err != nil {
		logrus.Errorf("error occurred while closing db connection : %s", err.Error())
	}
	logrus.Info("DB connection closed")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
