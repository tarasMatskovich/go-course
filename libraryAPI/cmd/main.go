package main

import (
	"library"
	"library/config"
	"library/pkg/handler"
	"library/pkg/repository"
	"library/pkg/service"
	"library/pkg/storage"

	"github.com/sirupsen/logrus"
)

func main() {
	configPath := "./configs/config.env"
	config, err := config.New(configPath)
	if err != nil {
		logrus.Fatalf("Error on initializing config: %s", err.Error())
	}
	storage, err := storage.NewStorage(config.RepoFilePath)
	if err != nil {
		logrus.Fatalf("Error on initializing storage: %s", err.Error())
	}
	defer storage.File.Close()
	repos := repository.NewRepository(storage)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(library.Server)

	if err := server.Start(config.Port, handlers.NewRouter()); err != nil {
		logrus.Fatalf("Error ocured while running http server %s", err.Error())
	}
}
