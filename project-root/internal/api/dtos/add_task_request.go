package dtos

type AddTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Descryption string `json:"description"`
	FileName    string `json:"file_name"`
	Children    []AddTaskRequest
}
