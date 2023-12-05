package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/services"
)

type handlers struct {
	taskService services.TaskService
}

func NewHandlers(taskService services.TaskService) *handlers {
	return &handlers{taskService: taskService}
}

func (h *handlers) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func (h *handlers) PostTaskHandler(c echo.Context) error {
	var taskRequest dtos.AddTaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	return c.JSON(201, taskRequest)
}