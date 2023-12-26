package handlers

import "github.com/marioTiara/todolistapi/internal/app/services"

type Handler struct {
	service services.Service
}

func NewHandlers(service services.Service) *Handler {
	return &Handler{service: service}
}
