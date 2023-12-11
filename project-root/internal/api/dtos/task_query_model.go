package dtos

import "time"

type TaskQueryModel struct {
	ID          string              `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"last_updated"`
	Priority    uint                `json:"priority"`
	Checked     bool                `json:"checked"`
	IsActive    bool                `json:"is_active"`
	Files       []FileQueryModel    `json:"files"`
	SubTasks    []SubtaskQueryModel `josn:"sub_tasks"`
}
