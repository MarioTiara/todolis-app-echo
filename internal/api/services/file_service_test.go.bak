package services

import (
	"errors"
	"mime/multipart"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	datafake "github.com/marioTiara/todolistapp/data-fake"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/utils"
	"github.com/marioTiara/todolistapp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock_uow := mocks.NewMockUnitOfWork(ctrl)
	mock_store := mocks.NewMockStorage(ctrl)

	//Act
	service := NewFileSevice(mock_uow, mock_store)

	//Assert
	assert.NotNil(t, service)
}

func TestGetByTaskID(t *testing.T) {
	files := datafake.GenerateFilesList(3)
	testCases := []struct {
		name          string
		mockSetup     func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
		expectedTasks []models.Files
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork) {
				mockFileskRepo.EXPECT().GetByTaskID(gomock.Any()).
					Return(nil, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
			expectedTasks: nil,
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork) {
				mockFileskRepo.EXPECT().GetByTaskID(gomock.Any()).
					Return(files, nil)
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: nil,
			expectedTasks: files,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_repo := mocks.NewMockFileRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_repo, mock_uow)

			//Create Task Service Instance
			taskService := NewFileSevice(mock_uow, mock_store)

			//Act
			files, err := taskService.GetByTaskID(1)
			assert.Equal(t, err, tc.expectedError)
			if tc.expectedError == nil {
				for i := 0; i < len(files); i++ {
					assert.Equal(t, tc.expectedTasks[i], files[i])
				}
			}

		})
	}
}

func TestGyByID(t *testing.T) {
	file := datafake.GenerateFile()
	testCases := []struct {
		name           string
		mockSetup      func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork)
		expectedError  error
		expectedResult models.Files
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork) {
				mockFileskRepo.EXPECT().GetByID(gomock.Any()).
					Return(models.Files{}, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError:  errors.New("mock error"),
			expectedResult: models.Files{},
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork) {
				mockFileskRepo.EXPECT().GetByID(gomock.Any()).
					Return(file, nil)
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError:  nil,
			expectedResult: file,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_repo := mocks.NewMockFileRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_repo, mock_uow)

			//Create Task Service Instance
			taskService := NewFileSevice(mock_uow, mock_store)

			//Act
			result, err := taskService.GetByID(1)
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, result)

		})
	}
}

func TestDeleteByID(t *testing.T) {
	testCases := []struct {
		name          string
		mockSetup     func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork)
		expectedError error
	}{
		{
			name: "error during retrieval",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork) {
				mockFileskRepo.EXPECT().DeleteByID(gomock.Any()).
					Return(errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork) {
				mockFileskRepo.EXPECT().DeleteByID(gomock.Any()).
					Return(nil)
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
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
			mock_repo := mocks.NewMockFileRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_repo, mock_uow)

			//Create Task Service Instance
			taskService := NewFileSevice(mock_uow, mock_store)

			//Act
			err := taskService.DeleteByID(1)
			assert.Equal(t, err, tc.expectedError)

		})
	}
}

func TestDeleteByTaskID(t *testing.T) {
	files := datafake.GenerateFilesList(2)
	testCases := []struct {
		name          string
		mockSetup     func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork, store *mocks.MockStorage)
		expectedError error
	}{
		{
			name: "error during loadFile from DB",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork, store *mocks.MockStorage) {
				mockFileskRepo.EXPECT().GetByTaskID(gomock.Any()).
					Return([]models.Files{}, errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(1)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)
			},
			expectedError: errors.New("mock error"),
		},
		{
			name: "error during delete file from storage",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork, store *mocks.MockStorage) {
				mockFileskRepo.EXPECT().GetByTaskID(gomock.Any()).
					Return(files, nil)
				store.EXPECT().DeleteFile(gomock.Any()).Return(errors.New("mock error"))
				mockuow.EXPECT().Begin().Times(2)
				mockuow.EXPECT().FileRepository().Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)

			},
			expectedError: errors.New("mock error"),
		},
		{
			name: "error during delete file from database",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork, store *mocks.MockStorage) {
				mockFileskRepo.EXPECT().GetByTaskID(gomock.Any()).
					Return(files, nil)
				mockFileskRepo.EXPECT().DeleteByID(gomock.Any()).Times(1).Return(errors.New("mock error"))
				store.EXPECT().DeleteFile(gomock.Any()).Times(1).Return(nil)
				mockuow.EXPECT().Begin().Times(2)
				mockuow.EXPECT().FileRepository().Times(2).Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(1)

			},
			expectedError: errors.New("mock error"),
		},
		{
			name: "success delete data from db and storage",
			mockSetup: func(mockFileskRepo *mocks.MockFileRepository, mockuow *mocks.MockUnitOfWork, store *mocks.MockStorage) {
				mockFileskRepo.EXPECT().GetByTaskID(gomock.Any()).
					Return(files, nil)
				mockFileskRepo.EXPECT().DeleteByID(gomock.Any()).Times(len(files)).Return(nil)
				store.EXPECT().DeleteFile(gomock.Any()).Times(len(files)).Return(nil)
				mockuow.EXPECT().Begin().Times(2)
				mockuow.EXPECT().FileRepository().Times(len(files) + 1).Return(mockFileskRepo)
				mockuow.EXPECT().Commit().Times(2)

			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_repo := mocks.NewMockFileRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_repo, mock_uow, mock_store)

			//Create Task Service Instance
			taskService := NewFileSevice(mock_uow, mock_store)

			//Act
			err := taskService.DeleteByTaskID(1)
			assert.Equal(t, err, tc.expectedError)

		})
	}
}

func TestDownload(t *testing.T) {
	dir := "uploads/image.jpg"
	testCases := []struct {
		name           string
		mockSetup      func(mockStore *mocks.MockStorage)
		expectedError  error
		expectedResult string
	}{
		{
			name: "error during loadFile",
			mockSetup: func(mockStore *mocks.MockStorage) {
				mockStore.EXPECT().LoadFile(gomock.Any()).Return(nil, errors.New("mock error"))
			},
			expectedError:  errors.New("mock error"),
			expectedResult: "",
		},
		{
			name: "Successful retrieval",
			mockSetup: func(mockStore *mocks.MockStorage) {
				mockStore.EXPECT().LoadFile(gomock.Any()).Return(dir, nil)
			},
			expectedError:  nil,
			expectedResult: dir,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_store)

			//Create Task Service Instance
			taskService := NewFileSevice(mock_uow, mock_store)

			//Act
			result, err := taskService.Download(gomock.Any().String())
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, result, tc.expectedResult)

		})
	}
}

func TestSaveFile(t *testing.T) {
	dir := "../../../static"
	fileName := "test_file.png"
	wrong_file := "wrong_name.png"

	_, expectedFileInforErr := os.Stat(dir + "/" + wrong_file)
	file := datafake.GenerateFile()
	modelfileQuery := utils.ConverFileToFileQueryModel(file)

	testCases := []struct {
		name           string
		mockSetup      func(mockStore *mocks.MockStorage, mockFileRepo *mocks.MockFileRepository, mock_uow *mocks.MockUnitOfWork)
		expectedError  error
		expectedResult dtos.FileQueryModel
	}{
		{
			name: "error during save file in storage",
			mockSetup: func(mockStore *mocks.MockStorage, mockFileRepo *mocks.MockFileRepository, mock_uow *mocks.MockUnitOfWork) {
				mockStore.EXPECT().SaveFile(gomock.Any()).Return("", errors.New("mock error"))
			},
			expectedError:  errors.New("mock error"),
			expectedResult: dtos.FileQueryModel{},
		},
		{
			name: "err file is not found ",
			mockSetup: func(mockStore *mocks.MockStorage, mockFileRepo *mocks.MockFileRepository, mock_uow *mocks.MockUnitOfWork) {
				mockStore.EXPECT().SaveFile(gomock.Any()).Return("wrong_name.png", nil)
				mockStore.EXPECT().Path().Return(dir)
			},
			expectedError:  expectedFileInforErr,
			expectedResult: dtos.FileQueryModel{},
		},
		{
			name: "Failed save data to database",
			mockSetup: func(mockStore *mocks.MockStorage, mockFileRepo *mocks.MockFileRepository, mock_uow *mocks.MockUnitOfWork) {
				mockStore.EXPECT().SaveFile(gomock.Any()).Return(fileName, nil)
				mockStore.EXPECT().Path().Return(dir)
				mockFileRepo.EXPECT().Create(gomock.Any()).Return(models.Files{}, errors.New("mock error"))
				mock_uow.EXPECT().Begin().Times(1)
				mock_uow.EXPECT().Commit().Times(1)
				mock_uow.EXPECT().FileRepository().Return(mockFileRepo)
			},
			expectedError:  errors.New("mock error"),
			expectedResult: dtos.FileQueryModel{},
		},
		{
			name: "Success save data to database",
			mockSetup: func(mockStore *mocks.MockStorage, mockFileRepo *mocks.MockFileRepository, mock_uow *mocks.MockUnitOfWork) {
				mockStore.EXPECT().SaveFile(gomock.Any()).Return(fileName, nil)
				mockStore.EXPECT().Path().Return(dir)
				mockFileRepo.EXPECT().Create(gomock.Any()).Return(file, nil)
				mock_uow.EXPECT().Begin().Times(1)
				mock_uow.EXPECT().Commit().Times(1)
				mock_uow.EXPECT().FileRepository().Return(mockFileRepo)
			},
			expectedError:  nil,
			expectedResult: modelfileQuery,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Create mockinstance
			mock_fileRepo := mocks.NewMockFileRepository(ctrl)
			mock_uow := mocks.NewMockUnitOfWork(ctrl)
			mock_store := mocks.NewMockStorage(ctrl)

			tc.mockSetup(mock_store, mock_fileRepo, mock_uow)

			//Create Task Service Instance
			taskService := NewFileSevice(mock_uow, mock_store)

			//Act
			result, err := taskService.SaveFile(1, &multipart.FileHeader{})
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedResult, result)

		})
	}
}
