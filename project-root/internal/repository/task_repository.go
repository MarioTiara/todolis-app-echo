package repository

import (
	"time"

	"github.com/marioTiara/todolistapp/internal/api/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	SoftDelete(id uint) error
	Update(task models.Task) (models.Task, error)
	FindAll() ([]models.Task, error)
	FindByID(ID uint, preload bool) (models.Task, error)
	FindSubTaskByTaskID(title, description string, parentID uint, page, limit int) ([]models.Task, error)
	Create(task models.Task) (models.Task, error)
	CreateSubTask(task models.Task) (models.Task, error)
	FilterByTitleAndDescription(title, description string, page, limit int, preload bool) ([]models.Task, error)
}

type task_repository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &task_repository{db: db}
}

func (r *task_repository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *task_repository) FindByID(ID uint, preload bool) (models.Task, error) {
	var parentTaskWithChildren models.Task
	var err error
	if preload {

		err = r.db.Preload("Children", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active= ?", true)
		}).Where("is_active= ?", true).First(&parentTaskWithChildren, ID).Error

	} else {
		err = r.db.Where("is_active= ?", true).First(&parentTaskWithChildren, ID).Error
	}

	return parentTaskWithChildren, err
}

func (r *task_repository) FindSubTaskByTaskID(title, description string, parentID uint, page, limit int) ([]models.Task, error) {
	var Subtasks []models.Task
	query := r.db.Model(&models.Task{}).Order("priority asc").Where("is_active = ?", true).Where("parent_id= ?", parentID)

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	offset := (page - 1) * limit

	if err := query.Offset(offset).Limit(limit).Find(&Subtasks).Error; err != nil {
		return nil, err
	}

	return Subtasks, nil

}

func (r *task_repository) Create(task models.Task) (models.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *task_repository) FilterByTitleAndDescription(title, description string, page, limit int, preload bool) ([]models.Task, error) {
	var tasks []models.Task
	query := r.db.Model(&models.Task{}).Order("priority asc").Where("is_active = ?", true)

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	offset := (page - 1) * limit

	if preload {
		query = query.Preload("Children")
	}
	if !preload {
		query = query.Where("parent_id IS NULL")
	}
	if err := query.Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *task_repository) CreateSubTask(task models.Task) (models.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *task_repository) SoftDelete(id uint) error {
	var taskTageted models.Task

	r.db.Find(&taskTageted, id)
	err := r.db.Model(&taskTageted).Update("is_active", false).Error
	return err
}

func (r *task_repository) Update(task models.Task) (models.Task, error) {
	oldData, err := r.FindByID(task.ID, false)
	if err != nil {
		return models.Task{}, err
	}
	oldData.Title = task.Title
	oldData.Description = task.Description
	oldData.UpdatedAt = time.Now().UTC()
	err = r.db.Updates(oldData).Error

	return oldData, err
}
