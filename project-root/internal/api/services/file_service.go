package services

import (
	"fmt"
	"mime/multipart"

	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/platform/storages"
	"github.com/marioTiara/todolistapp/internal/repository"
)

type FileService interface {
	SaveFile(taskID uint, file *multipart.FileHeader) (*models.Files, error)
	GetByTaskID(taskID uint) (*[]models.Files, error)
	GetByID(fileID uint) (*models.Files, error)
	DeleteByID(fileID uint) error
	DeleteByTaskID(taskID uint) error
	Download(fileName string) string
}

type files_service struct {
	uow   repository.UnitOfWork
	store storages.Storage
}

func NewFileSevice(uow repository.UnitOfWork, store storages.Storage) FileService {
	return &files_service{uow, store}
}

func (s *files_service) SaveFile(taskID uint, file *multipart.FileHeader) (*models.Files, error) {
	fileName, err := s.store.SaveFile(file)
	fileDetail := models.Files{}
	if err != nil {
		return &fileDetail, err
	}
	fileDetail.FileName = fileName
	fileDetail.TaskID = uint(taskID)

	s.uow.Begin()
	savedFile, _ := s.uow.FileRepository().Create(fileDetail)
	s.uow.Commit()

	return &savedFile, err
}

func (s *files_service) GetByTaskID(taskID uint) (*[]models.Files, error) {
	s.uow.Begin()
	files, err := s.uow.FileRepository().GetByTaskID(taskID)
	s.uow.Commit()
	return &files, err
}

func (s *files_service) GetByID(fileID uint) (*models.Files, error) {
	s.uow.Begin()
	file, err := s.uow.FileRepository().GetByID(fileID)
	s.uow.Commit()
	return &file, err
}

func (s *files_service) DeleteByID(fileID uint) error {
	s.uow.Begin()
	err := s.uow.FileRepository().DeleteByID(fileID)
	s.uow.Commit()
	return err
}

func (s *files_service) DeleteByTaskID(taskID uint) error {
	s.uow.Begin()
	files, err := s.uow.FileRepository().GetByTaskID(taskID)
	s.uow.Commit()
	if err != nil {
		return nil
	}
	s.uow.Begin()
	for _, file := range files {
		s.store.DeleteFile(file.FileName)
		s.uow.FileRepository().DeleteByID(file.ID)
	}
	s.uow.Commit()
	return err
}

func (s *files_service) Download(fileName string) string {
	dir, _ := s.store.LoadFile(fileName)
	str := fmt.Sprint(dir)
	return str
}
