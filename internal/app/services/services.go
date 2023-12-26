package services

import (
	"github.com/marioTiara/todolistapi/internal/app/repository"
	"github.com/marioTiara/todolistapi/internal/app/storages"
)

type Service interface {
	TaskService() TaskService
	FileService() FileService
}

type service struct {
	uow   repository.UnitOfWork
	store storages.Storage
}

func NewServices(uow repository.UnitOfWork, store storages.Storage) Service {
	return &service{uow, store}
}

func (s *service) TaskService() TaskService {
	return NewTaskService(s.uow, s.store)
}

func (s *service) FileService() FileService {
	return NewFileSevice(s.uow, s.store)
}
