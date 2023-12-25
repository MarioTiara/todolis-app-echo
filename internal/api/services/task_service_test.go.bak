package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	datafake "github.com/marioTiara/todolistapp/data-fake"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/utils"
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
	service := NewTaskService(mock_uow, mock_store)

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
	fakeSubTasks := datafake.GenerateSubtaskList(2)
	subtaskQueryModels := []dtos.SubtaskQueryModel{}

	for _, subTask := range fakeSubTasks {
		subTaskQueryModel := utils.ConvertSubTaskToSubtaskQueryModel(subTask)
		subtaskQueryModels = append(subtaskQueryModels, subTaskQueryModel)
	}

	testCases := []struct {
		name          string
		mockSetup     func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
		expectedTasks []dtos.SubtaskQueryModel
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
			expectedTasks: []dtos.SubtaskQueryModel{},
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FindSubTaskByTaskID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(fakeSubTasks, nil)

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
			expectedTasks: subtaskQueryModels,
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
					assert.Equal(t, tc.expectedTasks[i].ID, result[i].ID)
					assert.Equal(t, tc.expectedTasks[i].Title, result[i].Title)
					assert.Equal(t, tc.expectedTasks[i].Description, result[i].Description)
					assert.Equal(t, tc.expectedTasks[i].CreatedAt, result[i].CreatedAt)
					assert.Equal(t, tc.expectedTasks[i].UpdatedAt, result[i].UpdatedAt)
					assert.Equal(t, tc.expectedTasks[i].Priority, result[i].Priority)
					assert.Equal(t, tc.expectedTasks[i].Checked, result[i].Checked)
					assert.Equal(t, tc.expectedTasks[i].IsActive, result[i].IsActive)
					assert.Equal(t, tc.expectedTasks[i].ParentID, result[i].ParentID)
				}
			}
		})
	}
}

func TestCreate(t *testing.T) {
	addTaskRequest := datafake.GenerateAddTaskRequest(2)
	task := utils.ConvertRequestToTaskEntity(addTaskRequest)
	testCases := []struct {
		name          string
		mockSetup     func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
		expectedTasks models.Task
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().Create(gomock.Any()).
					Return(models.Task{}, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
			expectedTasks: models.Task{},
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().Create(gomock.Any()).
					Return(task, nil)

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
			expectedTasks: task,
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
			result, err := taskService.Create(addTaskRequest)
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedTasks, result)

		})
	}

}

func TestCreateSubTask(t *testing.T) {
	subTaskRequest := datafake.GenerateAddSubTaskRequest(1)
	task := utils.ConvertSubTaskRequestToTaskEntity(subTaskRequest)
	testCases := []struct {
		name          string
		mockSetup     func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
		expectedTasks models.Task
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().CreateSubTask(gomock.Any()).
					Return(models.Task{}, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
			expectedTasks: models.Task{},
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().CreateSubTask(gomock.Any()).
					Return(task, nil)

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
			expectedTasks: task,
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
			result, err := taskService.CreateSubTask(subTaskRequest)
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedTasks, result)

		})
	}

}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name          string
		mockSetup     func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().SoftDelete(gomock.Any()).
					Return(errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
		},
		{
			name: "Successful Delete",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().SoftDelete(gomock.Any()).
					Return(nil)

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
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
			err := taskService.Delete(1)
			assert.Equal(t, err, tc.expectedError)

		})
	}

}

func TestUpdate(t *testing.T) {
	updateRequest := datafake.GenerateUpdateTaskRequest()
	task := models.Task{
		ID:          updateRequest.ID,
		Title:       updateRequest.Title,
		Description: updateRequest.Description,
		Priority:    updateRequest.Priority,
		ParentID:    updateRequest.ParentID,
		Checked:     updateRequest.Checked,
	}

	testCases := []struct {
		name           string
		mockSetup      func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError  error
		expectedResult models.Task
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().Update(gomock.Any()).
					Return(models.Task{}, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError:  errors.New("mock error"),
			expectedResult: models.Task{},
		},

		{
			name: "Successful Update ",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().Update(gomock.Any()).
					Return(task, nil)

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError:  nil,
			expectedResult: task,
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
			result, err := taskService.Update(updateRequest)
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, result, tc.expectedResult)

		})
	}

}

func TestFilterTask(t *testing.T) {
	tasks := datafake.GenerateTasksList(3)

	taskQueryModels := []dtos.TaskQueryModel{}
	for _, task := range tasks {
		queryModel := utils.ConvertTaskToQueryModel(task)
		taskQueryModels = append(taskQueryModels, queryModel)
	}

	testCases := []struct {
		name           string
		mockSetup      func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork)
		expectedError  error
		expectedResult []dtos.TaskQueryModel
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FilterByTitleAndDescription(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).
					Return([]models.Task{}, errors.New("mock error"))

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError:  errors.New("mock error"),
			expectedResult: []dtos.TaskQueryModel{},
		},

		{
			name: "Successful Retrieve",
			mockSetup: func(mockTaskRepo *mocks.MockTaskRepository, mockuow *mocks.MockUnitOfWork) {
				mockTaskRepo.EXPECT().FilterByTitleAndDescription(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(tasks, nil)

				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().TaskRepository().Return(mockTaskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError:  nil,
			expectedResult: taskQueryModels,
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
			result, err := taskService.FilterTask(gomock.Any().String(), gomock.Any().String(),
				1, 10, true)
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, result, tc.expectedResult)

		})
	}

}
