package services

import (
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/utils"
	"github.com/marioTiara/todolistapp/internal/platform/storages"
	"github.com/marioTiara/todolistapp/internal/repository"
)

type TaskService interface {
	FindAll() ([]dtos.TaskQueryModel, error)
	FindSubTaskByTaskID(title, description string, parentID uint, page, limit int) ([]dtos.SubtaskQueryModel, error)
	FindByID(ID uint, preload bool) (dtos.TaskQueryModel, error)
	CreateSubTask(subTask dtos.AddSubTaskRequest) (models.Task, error)
	Create(task dtos.AddTaskRequest) (models.Task, error)
	FilterTask(title, description string, page, limit int, preload bool) ([]dtos.TaskQueryModel, error)
	Delete(id uint) error
	Update(task dtos.UpdateTaskRequest) (models.Task, error)
}

type task_service struct {
	uow   repository.UnitOfWork
	store storages.Storage
}

func NewTaskService(uow repository.UnitOfWork, store storages.Storage) TaskService {
	return &task_service{uow, store}
}

func (s *task_service) FindAll() ([]dtos.TaskQueryModel, error) {
	s.uow.Begin()
	tasks, err := s.uow.TaskRepository().FindAll()
	s.uow.Commit()
	if err != nil {
		return []dtos.TaskQueryModel{}, err
	}
	taskQueryModels := []dtos.TaskQueryModel{}
	for _, task := range tasks {
		queryModel := utils.ConvertTaskToQueryModel(task)
		taskQueryModels = append(taskQueryModels, queryModel)
	}
	return taskQueryModels, err
}

func (s *task_service) FindByID(ID uint, preload bool) (dtos.TaskQueryModel, error) {
	s.uow.Begin()
	task, err := s.uow.TaskRepository().FindByID(ID, preload)
	s.uow.Commit()
	if err != nil {
		return dtos.TaskQueryModel{}, err
	}
	return utils.ConvertTaskToQueryModel(task), err
}

func (s *task_service) FindSubTaskByTaskID(title, description string, parentID uint, page, limit int) ([]dtos.SubtaskQueryModel, error) {
	s.uow.Begin()
	subtasks, err := s.uow.TaskRepository().FindSubTaskByTaskID(title, description, parentID, page, limit)
	s.uow.Commit()
	subtaskQueryModels := []dtos.SubtaskQueryModel{}
	if err != nil {
		return subtaskQueryModels, err
	}
	for _, subTask := range subtasks {
		subTaskQueryModel := utils.ConvertSubTaskToSubtaskQueryModel(subTask)
		subtaskQueryModels = append(subtaskQueryModels, subTaskQueryModel)
	}
	return subtaskQueryModels, err
}

func (s *task_service) Create(task dtos.AddTaskRequest) (models.Task, error) {
	var parentTask = utils.ConvertRequestToTaskEntity(task)
	s.uow.Begin()
	createdTask, err := s.uow.TaskRepository().Create(parentTask)
	s.uow.Commit()

	if err != nil {
		return createdTask, err
	}
	return createdTask, err
}

func (s *task_service) CreateSubTask(request dtos.AddSubTaskRequest) (models.Task, error) {
	var subTask = utils.ConvertSubTaskRequestToTaskEntity(request)
	s.uow.Begin()
	task, err := s.uow.TaskRepository().CreateSubTask(subTask)
	s.uow.Commit()
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *task_service) FilterTask(title, description string, page, limit int, preload bool) ([]dtos.TaskQueryModel, error) {
	s.uow.Begin()
	tasks, err := s.uow.TaskRepository().FilterByTitleAndDescription(title, description, page, limit, preload)
	s.uow.Commit()

	taskQueryModels := []dtos.TaskQueryModel{}
	if err != nil {
		return taskQueryModels, err
	}
	filterdData := removeSubtaskFromParentList(tasks)
	for _, task := range filterdData {
		queryModel := utils.ConvertTaskToQueryModel(task)
		taskQueryModels = append(taskQueryModels, queryModel)
	}
	return taskQueryModels, err
}

func (s *task_service) Delete(id uint) error {
	s.uow.Begin()
	err := s.uow.TaskRepository().SoftDelete(id)
	s.uow.Commit()
	return err
}

func (s *task_service) Update(task dtos.UpdateTaskRequest) (models.Task, error) {
	newtask := models.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		ParentID:    task.ParentID,
		Checked:     task.Checked,
	}
	s.uow.Begin()
	updatedTask, err := s.uow.TaskRepository().Update(newtask)
	s.uow.Commit()
	if err != nil {
		return models.Task{}, err
	}
	return updatedTask, err
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
