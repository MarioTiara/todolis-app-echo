package models

type Files struct {
	ID       uint   `gorm:"column:id"`
	FileName string `gorm:"file_name"`
	TaskID   uint   `gorm:"task_id"`
	Task     Task
}
