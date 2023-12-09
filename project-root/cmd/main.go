package main

import (
	"github.com/marioTiara/todolistapp/config"
	"github.com/marioTiara/todolistapp/internal/api/services"
	"github.com/marioTiara/todolistapp/internal/platform/database"
	"github.com/marioTiara/todolistapp/internal/platform/server"
	"github.com/marioTiara/todolistapp/internal/platform/storages"
	"github.com/marioTiara/todolistapp/internal/repository"
)

func main() {
	configuration, _ := config.LoadConfig("./config")
	postgress := database.NewPostGressDB("host=localhost user=root password=secret dbname=todolistwebapi port=5432 sslmode=disable")
	store := storages.NewLocalStoarge("uploads")
	uow := repository.NewUnitOfWork(postgress.GetDB())
	service := services.NewTaskService(uow, store)

	server, _ := server.NewServer(configuration, service)
	server.Start()

}
