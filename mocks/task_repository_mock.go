// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repository/task_repository.go

// Package mock_repository is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/marioTiara/todolistapi/internal/app/models"
)

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTaskRepository) Create(task models.Task) (models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", task)
	ret0, _ := ret[0].(models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTaskRepositoryMockRecorder) Create(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskRepository)(nil).Create), task)
}

// CreateSubTask mocks base method.
func (m *MockTaskRepository) CreateSubTask(task models.Task) (models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubTask", task)
	ret0, _ := ret[0].(models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubTask indicates an expected call of CreateSubTask.
func (mr *MockTaskRepositoryMockRecorder) CreateSubTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubTask", reflect.TypeOf((*MockTaskRepository)(nil).CreateSubTask), task)
}

// FilterByTitleAndDescription mocks base method.
func (m *MockTaskRepository) FilterByTitleAndDescription(title, description string, page, limit int, preload bool) ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterByTitleAndDescription", title, description, page, limit, preload)
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterByTitleAndDescription indicates an expected call of FilterByTitleAndDescription.
func (mr *MockTaskRepositoryMockRecorder) FilterByTitleAndDescription(title, description, page, limit, preload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByTitleAndDescription", reflect.TypeOf((*MockTaskRepository)(nil).FilterByTitleAndDescription), title, description, page, limit, preload)
}

// FindAll mocks base method.
func (m *MockTaskRepository) FindAll() ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTaskRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTaskRepository)(nil).FindAll))
}

// FindByID mocks base method.
func (m *MockTaskRepository) FindByID(ID uint, preload bool) (models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ID, preload)
	ret0, _ := ret[0].(models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockTaskRepositoryMockRecorder) FindByID(ID, preload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockTaskRepository)(nil).FindByID), ID, preload)
}

// FindSubTaskByTaskID mocks base method.
func (m *MockTaskRepository) FindSubTaskByTaskID(title, description string, parentID uint, page, limit int) ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSubTaskByTaskID", title, description, parentID, page, limit)
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSubTaskByTaskID indicates an expected call of FindSubTaskByTaskID.
func (mr *MockTaskRepositoryMockRecorder) FindSubTaskByTaskID(title, description, parentID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSubTaskByTaskID", reflect.TypeOf((*MockTaskRepository)(nil).FindSubTaskByTaskID), title, description, parentID, page, limit)
}

// SoftDelete mocks base method.
func (m *MockTaskRepository) SoftDelete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftDelete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SoftDelete indicates an expected call of SoftDelete.
func (mr *MockTaskRepositoryMockRecorder) SoftDelete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftDelete", reflect.TypeOf((*MockTaskRepository)(nil).SoftDelete), id)
}

// Update mocks base method.
func (m *MockTaskRepository) Update(task models.Task) (models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", task)
	ret0, _ := ret[0].(models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTaskRepositoryMockRecorder) Update(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskRepository)(nil).Update), task)
}
