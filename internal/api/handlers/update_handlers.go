package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
)

// [METHOD:POST/PUT] Mengubah data list/sub list dengan kritera input diatas.
func (h *Handler) Update(c echo.Context) error {
	var taskRequest dtos.UpdateTaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	_, err := h.service.TaskService().Update(taskRequest)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "failed to update"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": fmt.Sprintf("%d updated", taskRequest.ID),
	})
}
