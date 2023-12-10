package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// [METHOD:DELETE] Menghapus data list/sub list.
func (h *handlers) Delete(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid Input"})
	}

	err = h.service.TaskService().Delete(uint(id))
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": err})
	}
	return c.NoContent(http.StatusNoContent)
}
