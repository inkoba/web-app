package main

import (
	"github.com/inkoba/web-app/internal/config"
	"github.com/inkoba/web-app/internal/handlers"
	"github.com/inkoba/web-app/internal/repository"
	"github.com/inkoba/web-app/internal/server"
	"github.com/inkoba/web-app/internal/service"
)

func main() {
	logger := service.InitLog()
	logger.Info("Starting config for app")

	conf := config.GetInstance().Get()

	conn := repository.NewAuth(conf.Mongo.URI).GetClient()

	repository := repository.NewRepository(logger, conn)
	//services
	service := service.NewService(logger, repository)
	//handlers
	handlers := handlers.NewHandler(logger, service)
	//server
	server := server.NewServer(logger, handlers)
	server.Initialize()
}
