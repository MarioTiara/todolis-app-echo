package repository

import (
	"github.com/marioTiara/todolistapp/internal/api/models"
	"gorm.io/gorm"
)

type FileRepository interface {
	Create(file models.Files) (models.Files, error)
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
