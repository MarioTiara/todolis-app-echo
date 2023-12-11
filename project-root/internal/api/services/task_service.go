package services

import (
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/platform/storages"
	"github.com/marioTiara/todolistapp/internal/repository"
)

type TaskService interface {
	FindAll() ([]dtos.TaskQueryModel, error)
	FindSubTaskByTaskID(title, description string, parentID uint, page, limit int) (*[]models.Task, error)
	FindByID(ID uint, preload bool) (*models.Task, error)
	CreateSubTask(subTask dtos.AddSubTaskRequest) (models.Task, error)
	Create(task dtos.AddTaskRequest) (models.Task, error)
	FilterTask(title, description string, page, limit int, preload bool) ([]models.Task, error)
	Delete(id uint) error
	Update(task dtos.AddTaskRequest, id uint) (models.Task, error)
}

type task_service struct {
	uow   repository.UnitOfWork
	store storages.Storage
}

func NewTaskService(uow repository.UnitOfWork, store storages.Storage) TaskService {
	return &task_service{uow, store}
}

func (s *task_service) FindAll() ([]dtos.TaskQueryModel, error) {
	tasks, err := s.uow.TaskRepository().FindAll()

	return tasks, err
}

func (s *task_service) FindByID(ID uint, preload bool) (*models.Task, error) {
	s.uow.Begin()
	task, err := s.uow.TaskRepository().FindByID(ID, preload)
	s.uow.Commit()
	return &task, err
}

func (s *task_service) FindSubTaskByTaskID(title, description string, parentID uint, page, limit int) (*[]models.Task, error) {
	s.uow.Begin()
	subtasks, err := s.uow.TaskRepository().FindSubTaskByTaskID(title, description, parentID, page, limit)
	s.uow.Commit()
	return &subtasks, err
}

func (s *task_service) Create(task dtos.AddTaskRequest) (models.Task, error) {
	//datetime := time.Now()

	var parentTask = convertRequestToTaskEntity(task)
	for _, task := range task.Children {
		parentTask.Children = append(parentTask.Children, convertRequestToTaskEntity(task))
	}
	s.uow.Begin()
	createdTask, err := s.uow.TaskRepository().Create(parentTask)
	s.uow.Commit()
	return createdTask, err
}

func (s *task_service) CreateSubTask(request dtos.AddSubTaskRequest) (models.Task, error) {
	var subTask = models.Task{Title: request.Title, Description: request.Description, ParentID: &request.ParentID}
	return s.uow.TaskRepository().CreateSubTask(subTask)
}

func (s *task_service) FilterTask(title, description string, page, limit int, preload bool) ([]models.Task, error) {
	s.uow.Begin()
	task, err := s.uow.TaskRepository().FilterByTitleAndDescription(title, description, page, limit, preload)
	s.uow.Commit()
	filterdData := removeSubtaskFromParentList(task)
	return filterdData, err
}

func (s *task_service) Delete(id uint) error {
	s.uow.Begin()
	err := s.uow.TaskRepository().SoftDelete(id)
	s.uow.Commit()
	return err
}

func (s *task_service) Update(task dtos.AddTaskRequest, id uint) (models.Task, error) {
	newtask := models.Task{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
	}
	s.uow.Begin()
	updatedTask, err := s.uow.TaskRepository().Update(newtask)
	s.uow.Commit()
	return updatedTask, err
}

func convertRequestToTaskEntity(request dtos.AddTaskRequest) models.Task {
	newtask := models.Task{Title: request.Title, Description: request.Description}
	return newtask
}

func removeSubtaskFromParentList(tasks []models.Task) []models.Task {
	var filteredData []models.Task
	for _, task := range tasks {
		if task.ParentID == nil {
			filteredData = append(filteredData, task)
		}
	}

	return filteredData
}
