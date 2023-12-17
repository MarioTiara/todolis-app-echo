package handlers

import "github.com/marioTiara/todolistapp/internal/api/services"

type Handler struct {
	service services.Service
}

func NewHandlers(service services.Service) *Handler {
	return &Handler{service: service}
}
