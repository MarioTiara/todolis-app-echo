package dtos

type AddSubTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id"`
}
