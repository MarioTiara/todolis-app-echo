package models

import "time"

type Files struct {
	ID        uint      `gorm:"column:id"`
	FileName  string    `gorm:"file_name"`
	FileSize  uint      `gorm:"file_size"`
	FileURL   string    `gorm:"file_url"`
	CreatedAt time.Time `gorm:"created_at"`
	TaskID    uint      `gorm:"task_id"`

	Task Task
}
