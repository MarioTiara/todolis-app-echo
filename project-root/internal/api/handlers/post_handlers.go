package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/services"
)

type handlers struct {
	service services.Service
}

func NewHandlers(service services.Service) *handlers {
	return &handlers{service: service}
}

func (h *handlers) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

// 5.[METHOD:POST] Menambahkan data list.
func (h *handlers) PostTaskHandler(c echo.Context) error {

	contentType := c.Request().Header.Get("Content-Type")
	isMultipart := isMultipartRequest(contentType)

	if isMultipart {
		return h.PostTaskFormBodyHandler(c)
	}
	return h.PostTaskJsonBodyHandler(c)
}

func (h *handlers) PostTaskFormBodyHandler(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "file upload failed"})
	}
	//Retrive JOSN text from form data
	jsonData := form.Value["data"]

	//Parse JSON
	var taskRequest dtos.AddTaskRequest
	if err := json.Unmarshal([]byte(jsonData[0]), &taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	//Create new Task
	newTask := dtos.AddTaskRequest{Title: taskRequest.Title, Description: taskRequest.Description}
	task, err := h.service.TaskService().Create(newTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "failed to create new task"})
	}

	//Save file
	var filesDetail []models.Files
	files := form.File["files"]

	//Iterate the files each uploaded file
	for _, file := range files {
		data, _ := h.service.FileService().SaveFile(uint(task.ID), file)
		filesDetail = append(filesDetail, *data)
	}

	task.Files = filesDetail
	return c.JSON(http.StatusOK, task)
}

func (h *handlers) PostTaskJsonBodyHandler(c echo.Context) error {
	var taskRequest dtos.AddTaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	fmt.Println(taskRequest)
	task, err := h.service.TaskService().Create(taskRequest)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to create task"})
	}
	return c.JSON(201, task)
}

func isMultipartRequest(contentType string) bool {
	return len(contentType) >= len("multipart/form-data") && contentType[:len("multipart/form-data")] == "multipart/form-data"
}
