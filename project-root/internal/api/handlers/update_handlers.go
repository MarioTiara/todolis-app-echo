package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
)

// [METHOD:POST/PUT] Mengubah data list/sub list dengan kritera input diatas.
func (h *handlers) Update(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid Input"})
	}

	var taskRequest dtos.AddTaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	updateTask, err := h.service.TaskService().Update(taskRequest, uint(id))
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": err})
	}

	return c.JSON(http.StatusOK, updateTask)

}
