package server

import (
	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/config"
	"github.com/marioTiara/todolistapp/internal/api/routes"
	"github.com/marioTiara/todolistapp/internal/api/services"
)

type Server struct {
	config  config.Config
	service services.TaskService
}

func NewServer(config config.Config, service services.TaskService) (*Server, error) {
	server := &Server{
		config:  config,
		service: service,
	}
	return server, nil
}

func (server *Server) Start() {
	e := echo.New()
	routes.SetRoutes(e, server.service)
	// err := e.Start(server.config.ServerAddress)
	// return err
	if err := e.Start(":8080"); err != nil {
		panic("failed to start the server")
	}
}
