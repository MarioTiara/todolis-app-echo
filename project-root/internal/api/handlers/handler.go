package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/services"
	"gorm.io/gorm"
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

// 5.[METHOD:POST] Menambahkan data list.
func (h *handlers) PostTaskHandler(c echo.Context) error {

	contentType := c.Request().Header.Get("Content-Type")
	isMultipart := isMultipartRequest(contentType)

	if isMultipart {
		return h.HandleMultipleFileUpload(c)
	}
	var taskRequest dtos.AddTaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	fmt.Println(taskRequest)
	task, err := h.taskService.Create(taskRequest)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to create task"})
	}
	return c.JSON(201, task.ID)
}

// 2. [METHOD:GET] Menampilkan data detail list by list id.
// 4. [METHOD:GET] Menampilkan data detail sub list by sub list id.
func (h *handlers) GetTaskByIDHandler(c echo.Context) error {

	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	//preload flag
	var preloadFlag bool
	preloadSubTaskParam := c.QueryParam("preloadSubTasks")

	if preloadSubTaskParam != "" {
		preloadFlag, _ = strconv.ParseBool(preloadSubTaskParam)
	}

	task, err := h.taskService.FindByID(uint(id), preloadFlag)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusNoContent)
	}
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to load task"})
	}
	return c.JSON(200, task)
}

// 1. [METHOD:GET] Menampilkan data all list ( include pagination, filter[Search By: title, description] ) dengan atau tanpa preload sub list (dynamic)
func (h *handlers) GetAllList(c echo.Context) error {
	//pagination parameters
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	//Search filters
	title := c.QueryParam("title")
	description := c.QueryParam("description")

	//preload flag
	var preloadFlag bool
	preloadSubTaskParam := c.QueryParam("preloadSubTasks")

	if preloadSubTaskParam != "" {
		preloadFlag, _ = strconv.ParseBool(preloadSubTaskParam)
	}

	task, err := h.taskService.FilterTask(title, description, page, pageSize, preloadFlag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err})
	}

	return c.JSON(http.StatusOK, task)

}

// 3.[METHOD:GET] Menampilkan data all sub list by list id ( include pagination, filter[Search By: title, description] )
func (h *handlers) GetAllSubListByParentID(c echo.Context) error {
	strID := c.Param("parentID")
	parentID, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	//pagination parameters
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	//Search filters
	title := c.QueryParam("title")
	description := c.QueryParam("description")

	subTasks, err := h.taskService.FindSubTaskByTaskID(title, description, uint(parentID), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err})
	}

	return c.JSON(http.StatusOK, subTasks)
}

// [METHOD:DELETE] Menghapus data list/sub list.
func (h *handlers) Delete(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid Input"})
	}

	err = h.taskService.Delete(uint(id))
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": err})
	}
	return c.NoContent(http.StatusNoContent)
}

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

	updateTask, err := h.taskService.Update(taskRequest, uint(id))
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": err})
	}

	return c.JSON(http.StatusOK, updateTask)

}

func (h *handlers) HandleMultipleFileUpload(c echo.Context) error {
	taskID, _ := strconv.ParseUint(c.FormValue("taskID"), 10, 64)
	form, err := c.MultipartForm()
	var filesDetail []models.Files
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "file upload failed"})
	}

	files := form.File["files"]

	//Iterate the files each uploaded file
	for _, file := range files {
		data, _ := h.taskService.SaveFile(uint(taskID), file)
		filesDetail = append(filesDetail, *data)
	}

	return c.JSON(http.StatusOK, filesDetail)
}
func isMultipartRequest(contentType string) bool {
	return len(contentType) >= len("multipart/form-data") && contentType[:len("multipart/form-data")] == "multipart/form-data"
}
