package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapi/internal/app/dtos"
	"github.com/marioTiara/todolistapi/internal/app/utils"
)

// 5.[METHOD:POST] Menambahkan data list.
func (h *Handler) PostTaskHandler(c echo.Context) error {
	var taskRequest dtos.AddTaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}
	task, err := h.service.TaskService().Create(taskRequest)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to create task"})
	}
	return c.JSON(201, map[string]interface{}{"status": "success", "data": task})
}

// Upload Files
func (h *Handler) UploadTaskFilesHandler(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.FormValue("taskID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid Parameter"})
	}
	//Save file
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "file upload failed"})
	}

	var filesDetail []dtos.FileQueryModel
	files := form.File["files"]
	if len(files) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid Parameter"})
	}
	//Iterate the files each uploaded file
	for _, file := range files {
		if !utils.IsValidFileExtension(file) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "file extension should in .text or .pdf"})
		}
		data, err := h.service.FileService().SaveFile(uint(taskID), file)
		if err != nil {
			return c.JSON(500, map[string]interface{}{"error": "failed to save file"})
		}
		filesDetail = append(filesDetail, data)
	}
	return c.JSON(201, map[string]interface{}{"status": "success", "data": filesDetail})
}

func (h *Handler) PostSubTaskHandler(c echo.Context) error {
	var subTask dtos.AddSubTaskRequest
	if err := c.Bind(&subTask); err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}
	task, err := h.service.TaskService().CreateSubTask(subTask)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to create task"})
	}

	return c.JSON(201, map[string]interface{}{"status": "success", "ID": task.ID})
}
