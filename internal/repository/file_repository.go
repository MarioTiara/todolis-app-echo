package repository

import (
	"github.com/marioTiara/todolistapp/internal/api/models"
	"gorm.io/gorm"
)

type FileRepository interface {
	Create(file models.Files) (models.Files, error)
	GetByTaskID(taskID uint) ([]models.Files, error)
	GetByID(id uint) (models.Files, error)
	DeleteByID(id uint) error
	DeleteByTaskID(taskID uint) error
}

type file_repository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &file_repository{db: db}
}

func (r *file_repository) Create(file models.Files) (models.Files, error) {
	er := r.db.Create(&file).Error
	return file, er
}

func (r *file_repository) GetByTaskID(taskID uint) ([]models.Files, error) {
	var files []models.Files
	query := r.db.Model(&models.Files{}).Where("task_id = ?", taskID)
	if err := query.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (r *file_repository) GetByID(id uint) (models.Files, error) {
	var file models.Files
	err := r.db.First(&file, id).Error
	return file, err
}

func (r *file_repository) DeleteByID(id uint) error {
	err := r.db.Delete(&models.Files{}, id).Error
	return err
}

func (r *file_repository) DeleteByTaskID(taskID uint) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&models.Files{}).Error
	return err
}
