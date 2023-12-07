package services

import (
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/repository"
)

type TaskService interface {
	FindAll() ([]models.Task, error)
	FindByID(ID uint) (*models.Task, error)
	CreateSubTask(parentID uint, subTask dtos.AddTaskRequest) (models.Task, error)
	Create(task dtos.AddTaskRequest) (models.Task, error)
	FilterTask(title, description string, page, limit int, preload bool) ([]models.Task, error)
}

type task_service struct {
	uow repository.UnitOfWork
}

func NewTaskService(uow repository.UnitOfWork) TaskService {
	return &task_service{uow}
}

func (s *task_service) FindAll() ([]models.Task, error) {
	tasks, err := s.uow.TaskRepository().FindAll()
	return tasks, err
}

func (s *task_service) FindByID(ID uint) (*models.Task, error) {
	s.uow.Begin()
	task, err := s.uow.TaskRepository().FindByID(ID, false)
	s.uow.Commit()
	return &task, err
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

func (s *task_service) CreateSubTask(parentID uint, request dtos.AddTaskRequest) (models.Task, error) {
	var subTask = models.Task{Title: request.Title, Descryption: request.Descryption, ParentID: &parentID}
	return s.uow.TaskRepository().CreateSubTask(subTask)
}

func (s *task_service) FilterTask(title, description string, page, limit int, preload bool) ([]models.Task, error) {
	s.uow.Begin()
	task, err := s.uow.TaskRepository().FilterByTitleAndDescription(title, description, page, limit, preload)
	s.uow.Commit()
	filterdData := removeSubtaskFromParentList(task)
	return filterdData, err
}

func convertRequestToTaskEntity(request dtos.AddTaskRequest) models.Task {
	newtask := models.Task{Title: request.Title, Descryption: request.Descryption}
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
