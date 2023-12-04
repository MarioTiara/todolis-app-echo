package models

type Files struct {
	ID       int    `gorm:"column:id"`
	FileName string `gorm:"file_name"`
	FilePath string `gorm:"file_path"`
	TaskID   int    `gorm:"task_id"`
}
