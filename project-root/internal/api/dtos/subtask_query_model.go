package dtos

import "time"

type SubtaskQueryModel struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"last_updated"`
	Priority    uint             `json:"priority"`
	Checked     bool             `json:"checked"`
	IsActive    bool             `json:"is_active"`
	Files       []FileQueryModel `json:"files"`
	ParentID    uint             `json:"parent_id"`
}
