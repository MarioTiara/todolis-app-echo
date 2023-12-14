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
	configuration, _ := config.LoadConfig("../config")
	postgress := database.NewPostGressDB(configuration)
	store := storages.NewLocalStoarge("../uploads")
	uow := repository.NewUnitOfWork(postgress.GetDB())
	service := services.NewServices(uow, store)

	server, _ := server.NewServer(configuration, service)
	server.UseJWT()
	server.SetRoutes()
	server.Start()

}
