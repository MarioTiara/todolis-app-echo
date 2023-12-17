package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// [METHOD:DELETE] Menghapus data list/sub list.
func (h *Handler) DeleteTask(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid Input"})
	}

	err = h.service.TaskService().Delete(uint(id))
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to  delete Task"})
	}
	err = h.service.FileService().DeleteByTaskID(uint(id))
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to delete files"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteFile(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}
	err = h.service.FileService().DeleteByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to delete file"})
	}

	return c.NoContent(http.StatusNoContent)
}
