package main

import (
	"github.com/marioTiara/todolistapi/api"
	config "github.com/marioTiara/todolistapi/configs"
	"github.com/marioTiara/todolistapi/internal/app/repository"
	"github.com/marioTiara/todolistapi/internal/app/services"
	"github.com/marioTiara/todolistapi/internal/app/storages"
)

func main() {
	configuration, _ := config.LoadConfig("../config")
	postgress := repository.NewPostGressDB(configuration.DbSource)
	store := storages.NewLocalStoarge("../uploads")
	uow := repository.NewUnitOfWork(postgress.GetDB())
	service := services.NewServices(uow, store)

	server, _ := api.NewServer(configuration, service)
	server.UseJWT()
	server.SetRoutes()
	server.Start()
}
