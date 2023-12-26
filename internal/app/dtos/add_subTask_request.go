package dtos

type AddSubTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	ParentID    uint   `json:"parent_id" binding:"required"`
}
