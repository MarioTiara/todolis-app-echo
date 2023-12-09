package dtos

type AddTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Children    []AddTaskRequest
}
