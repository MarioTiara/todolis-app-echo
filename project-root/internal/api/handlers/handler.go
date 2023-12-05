package handlers

import (
	"net/http"
	"strconv"

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

	task, err := h.taskService.Create(taskRequest)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to create task"})
	}
	return c.JSON(201, task.ID)
}
func (h *handlers) GetTaskByIDHandler(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}
	task, err := h.taskService.FindByID(uint(id))
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to load task"})
	}
	return c.JSON(200, task)
}
