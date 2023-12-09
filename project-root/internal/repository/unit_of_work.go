package repository

import "gorm.io/gorm"

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error

	TaskRepository() TaskRepository
	FileRepository() FileRepository
}

type unitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) Begin() error {
	u.tx = u.db.Begin()
	return u.tx.Error
}

func (u *unitOfWork) Commit() error {
	return u.tx.Commit().Error
}

func (u *unitOfWork) Rollback() error {
	return u.tx.Rollback().Error
}

func (u *unitOfWork) TaskRepository() TaskRepository {
	return NewTaskRepository(u.tx)
}

func (u *unitOfWork) FileRepository() FileRepository {
	return NewFileRepository(u.tx)
}
