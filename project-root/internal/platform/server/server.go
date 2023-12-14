package server

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/config"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/internal/api/routes"
	"github.com/marioTiara/todolistapp/internal/api/services"
)

type Server struct {
	config  config.Config
	service services.Service
	e       *echo.Echo
	handler *handlers.Handler
	eGroup  *echo.Group
}

func NewServer(config config.Config, service services.Service) (*Server, error) {
	server := &Server{
		config:  config,
		service: service,
		e:       echo.New(),
		handler: handlers.NewHandlers(service),
	}
	return server, nil
}

func (server *Server) UseJWT() {
	server.eGroup = server.e.Group("/v1/restricted")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	server.eGroup.Use(echojwt.WithConfig(config))
	server.eGroup.GET("", server.handler.Restricted)
}

func (server *Server) SetRoutes() {
	server.e.GET("/", server.handler.Accessible)
	server.e.POST("/v1/login", server.handler.Login)
	routes.SetRoutes(server.eGroup, server.handler)
}
func (server *Server) Start() {
	if err := server.e.Start(":8080"); err != nil {
		panic("failed to start the server")
	}
}
