package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/services"
	"github.com/marioTiara/todolistapp/internal/api/utils"
)

type taskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *taskHandler {
	return &taskHandler{taskService}
}
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mario Tiara",
		"bio":  "A Full Stack SOftware Enginner",
	})
}

// [METHOD:POST] Menambahkan data list.
func (h *taskHandler) PostTasksHandler(c *gin.Context) {
	var taskRequest dtos.AddTaskRequest
	err := c.ShouldBindJSON((&taskRequest))

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	task, err := h.taskService.Create(taskRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}

func (h *taskHandler) GetTaskByIDHandler(c *gin.Context) {
	StrID := c.Param("id")
	id, err := strconv.ParseUint(StrID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "invalid id",
		})
	}

	task, err := h.taskService.FindByID(uint(id))
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}

func (h *taskHandler) FilterTaskHandler(c *gin.Context) {
	title := c.Query("title")
	description := c.Query("description")
	page, limit := utils.GetPageAndLimit(c)

	tasks, err := h.taskService.FilterTask(title, description, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to filter tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

func (h *taskHandler) PostSubTaskByID(c *gin.Context) {

	var subTaskRequest dtos.AddTaskRequest
	err := c.ShouldBindJSON(&subTaskRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	StrID := c.Param("id")
	id, err := strconv.ParseUint(StrID, 10, 64)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	task, err := h.taskService.CreateSubTask(uint(id), subTaskRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}

func (h *taskHandler) UploadFile(c *gin.Context) {
	var reqtask models.Task
	if err := c.ShouldBindJSON(&reqtask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	//handle file upload
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Save the file
	err = saveUploadedFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, "Ok")

}

func saveUploadedFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create("uploads" + file.Filename)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
