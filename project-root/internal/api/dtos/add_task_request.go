package dtos

type AddTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	FileName    string `json:"file_name"`
	Children    []AddTaskRequest
}
