package models

import "time"

type Task struct {
	ID          uint      `gorm:"column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"last_updated"`
	Priority    int       `gorm:"column:priority" json:"priority"`
	Checked     bool      `gorm:"column:checked" json:"checked"`
	IsActive    bool      `gorm:"column:is_active" json:"is_active"`
	ParentID    *uint     `json:"parent_id"`
	Children    []Task    `gorm:"foreignKey:ParentID" json:"sub_tasks"`
	Files       []Files   `gorm:"foreignKey:TaskID" json:"files"`
}
