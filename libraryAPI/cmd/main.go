package main

import (
	"library"
	"library/pkg/handler"
	"library/pkg/repository"
	"library/pkg/service"
	"library/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := config.New(); err != nil {
		logrus.Fatalf("Error on initializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(library.Server)

	if err := server.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error ocured while running http server %s", err.Error())
	}
}
