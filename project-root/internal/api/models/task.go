package models

import "time"

type Task struct {
	ID          uint      `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	Priority    int       `gorm:"column:priority"`
	Checked     bool      `gorm:"column:checked"`
	IsActive    bool      `gorm:"column:is_active"`
	ParentID    *uint
	Children    []Task  `gorm:"foreignKey:ParentID"`
	Files       []Files `gorm:"foreignKey:TaskID"`
}
