package dtos

type AddTaskRequest struct {
	Title       string           `json:"title" binding:"required"`
	Description string           `json:"description"`
	Priority    int              `json:"priority"`
	Childrens   []AddTaskRequest `json:"childrens"`
}

type NewTaskRequest struct {
	Title       string           `json:"title" binding:"required"`
	Description string           `json:"description"`
	Priority    int              `json:"priority"`
	Childrens   []NewTaskRequest `json:"childrens"`
}
