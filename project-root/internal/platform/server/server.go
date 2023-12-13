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
	r := server.e.Group("/restricted")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JwtCustomClaims)
		},

		SigningKey: []byte("secret"),
	}

	r.Use(echojwt.WithConfig(config))
	r.GET("", server.handler.Restricted)

}
func (server *Server) Start() {
	routes.SetRoutes(server.e, server.service)
	if err := server.e.Start(":8080"); err != nil {
		panic("failed to start the server")
	}
}
