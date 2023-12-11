package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewTaskService(t *testing.T) {
	//Arrange
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	mock_uow := mocks.NewMockUnitOfWork(ctr)
	mock_store := mocks.NewMockStorage(ctr)

	//Act
	service := NewFileSevice(mock_uow, mock_store)

	//Assert
	assert.NotNil(t, service)
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name          string
		mockSetup     func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
		expectedTasks []models.Task
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FindAll().Return([]models.Task{}, errors.New("mocker error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mocker error"),
			expectedTasks: []models.Task{},
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FindAll().Return([]models.Task{
					{ID: 1, Title: "Task 1"},
					{ID: 2, Title: "Task 2"},
				}, nil)
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
			expectedTasks: []models.Task{
				{ID: 1, Title: "Task 1"},
				{ID: 2, Title: "Task 2"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_repo := mocks.NewMockTaskRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)
			tc.mockSetup(mock_repo, mock_uow)

			//Create Task Service Instance
			taskService := NewTaskService(mock_uow, mock_store)

			//Act
			result, err := taskService.FindAll()
			assert.Equal(t, len(result), len(tc.expectedTasks))
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestFindByID(t *testing.T) {
	testCases := []struct {
		name          string
		mockSetup     func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
		expectedTasks dtos.TaskQueryModel
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FindByID(gomock.Any(), gomock.Any()).
					Return(models.Task{}, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
			expectedTasks: dtos.TaskQueryModel{},
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FindByID(gomock.Any(), gomock.Any()).
					Return(models.Task{ID: 1, Title: "Task 1"},
						nil)
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
			expectedTasks: dtos.TaskQueryModel{
				ID: 1, Title: "Task 1",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_repo := mocks.NewMockTaskRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_repo, mock_uow)

			//Create Task Service Instance
			taskService := NewTaskService(mock_uow, mock_store)

			//Act
			result, err := taskService.FindByID(1, true)
			assert.Equal(t, err, tc.expectedError)

			if tc.expectedError == nil {
				assert.Equal(t, result.ID, tc.expectedTasks.ID)
				assert.Equal(t, result.Title, tc.expectedTasks.Title)
			}
		})
	}

}

func TestFindSubTaskByTaskID(t *testing.T) {

	testCases := []struct {
		name          string
		mockSetup     func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
		expectedTasks []dtos.TaskQueryModel
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FindSubTaskByTaskID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return([]models.Task{}, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
			expectedTasks: []dtos.TaskQueryModel{},
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FindSubTaskByTaskID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return([]models.Task{
						{ID: 1, Title: "Task 1"},
						{ID: 2, Title: "Task 2"},
					}, nil)

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
			expectedTasks: []dtos.TaskQueryModel{
				{ID: 1, Title: "Task 1"},
				{ID: 2, Title: "Task 2"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_repo := mocks.NewMockTaskRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_repo, mock_uow)

			//Create Task Service Instance
			taskService := NewTaskService(mock_uow, mock_store)

			//Act
			result, err := taskService.FindSubTaskByTaskID("title", "desc", 1, 1, 10)
			assert.Equal(t, err, tc.expectedError)

			if tc.expectedError == nil {
				for i := range result {
					assert.Equal(t, tc.expectedTasks[i].ID, result[1].ID)
					assert.Equal(t, tc.expectedTasks[i].Title, result[1].Title)
				}
			}
		})
	}
}
