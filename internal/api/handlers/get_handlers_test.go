package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	datafake "github.com/marioTiara/todolistapp/data-fake"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetTaskByIDHandler(t *testing.T) {
	type successBodyResult struct {
		Status string              `json:"status"`
		Data   dtos.TaskQueryModel `json:"data"`
	}

	taskQueryModel := datafake.GenerateTaskQueryModel()
	testCases := []struct {
		name              string
		mockSetup         func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService)
		setRequest        func(e *echo.Echo) *http.Request
		expectedResponse  int
		expectedError     errorResult
		successDataresult successBodyResult
	}{
		{
			name: "Error: Invalid Input",
			mockSetup: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {

			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/tasks/a", nil)
				return req
			},
			expectedResponse: 400,
			expectedError: errorResult{
				Error: "Invalid input",
			},
		},
		// {
		// 	name: "No Content",
		// 	mockSetup: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
		// 		mockTaskService.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(dtos.TaskQueryModel{}, gorm.ErrRecordNotFound)
		// 		mockService.EXPECT().TaskService().Return(mockTaskService)
		// 	},
		// 	setRequest: func(e *echo.Echo) *http.Request {
		// 		q := make(url.Values)
		// 		q.Set("preloadSubTasks", "true")
		// 		req := httptest.NewRequest(http.MethodGet, "/tasks/123?"+q.Encode(), nil)
		// 		return req
		// 	},
		// 	expectedResponse: http.StatusNoContent,
		// },
		{
			name: "error: Failed to load task",
			mockSetup: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(dtos.TaskQueryModel{}, errors.New("mock error"))
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				q := make(url.Values)
				q.Set("preloadSubTasks", "true")
				req := httptest.NewRequest(http.MethodGet, "/tasks/123?"+q.Encode(), nil)
				return req
			},
			expectedResponse: 500,
			expectedError: errorResult{
				Error: "Failed to load task",
			},
		},
		{
			name: "success",
			mockSetup: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(taskQueryModel, nil)
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				q := make(url.Values)
				q.Set("preloadSubTasks", "true")
				req := httptest.NewRequest(http.MethodGet, "/tasks/123?"+q.Encode(), nil)
				return req
			},
			expectedResponse: 200,
			successDataresult: successBodyResult{
				Status: "success",
				Data:   taskQueryModel,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			e := echo.New()
			req := tc.setRequest(e)
			rec := httptest.NewRecorder()

			//Instantiate mocks
			mockService := mocks.NewMockService(ctr)
			mockTaskService := mocks.NewMockTaskService(ctr)
			tc.mockSetup(mockService, mockTaskService)

			//Instantiate Handlers
			handler := handlers.NewHandlers(mockService)
			e.GET("/tasks/:id", handler.GetTaskByIDHandler)
			tc.setRequest(e)

			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedResponse, rec.Code)
			if tc.expectedResponse != 200 {
				var Error errorResult
				json.Unmarshal(rec.Body.Bytes(), &Error)
				assert.Equal(t, tc.expectedError, Error)
			} else {
				var result successBodyResult
				json.Unmarshal(rec.Body.Bytes(), &result)
				assert.Equal(t, tc.successDataresult, result)
			}
		})
	}
}

func TestGetAllList(t *testing.T) {
	tasks := []dtos.TaskQueryModel{}
	for i := 0; i < 3; i++ {
		tasks = append(tasks, datafake.GenerateTaskQueryModel())
	}

	type successBodyResult struct {
		Status string                `json:"status"`
		Data   []dtos.TaskQueryModel `json:"data"`
	}

	testCases := []struct {
		name            string
		setupMock       func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService)
		setRequest      func(e *echo.Echo) *http.Request
		expectedCode    int
		expectedEror    errorResult
		expectedSuccess successBodyResult
	}{
		{
			name: "error: failed to load taks",
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().FilterTask(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(tasks, errors.New("mock error"))
				mockService.EXPECT().TaskService().Return(mockTaskService)

			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
				return req
			},
			expectedCode: 500,
			expectedEror: errorResult{
				Error: "Failed to load tasks",
			},
		},
		{
			name: "success",
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().FilterTask(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(tasks, nil)
				mockService.EXPECT().TaskService().Return(mockTaskService)

			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
				return req
			},
			expectedCode: 200,
			expectedSuccess: successBodyResult{
				Status: "success",
				Data:   tasks,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			e := echo.New()
			rec := httptest.NewRecorder()

			//Instantiate mocks
			mockService := mocks.NewMockService(ctr)
			mockTaskService := mocks.NewMockTaskService(ctr)
			tc.setupMock(mockService, mockTaskService)

			//Instantiate Hanlders
			handler := handlers.NewHandlers(mockService)
			e.GET("/tasks", handler.GetAllList)
			req := tc.setRequest(e)

			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedCode, rec.Code)
			if tc.expectedCode != 200 {
				var Error errorResult
				json.Unmarshal(rec.Body.Bytes(), &Error)
				assert.Equal(t, tc.expectedEror, Error)
			} else {
				var result successBodyResult
				json.Unmarshal(rec.Body.Bytes(), &result)
				assert.Equal(t, tc.expectedSuccess, result)
			}

		})
	}
}
func TestGetAllSubListByParentID(t *testing.T) {
	subtasks := []dtos.SubtaskQueryModel{}
	for i := 0; i < 3; i++ {
		subtasks = append(subtasks, datafake.GenerateSubTaskQueryModel())
	}
	type successResult struct {
		Status string                   `json:"status"`
		Data   []dtos.SubtaskQueryModel `json:"data"`
	}
	testCases := []struct {
		name                string
		ExpectedCode        int
		setupMock           func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService)
		setRequest          func(e *echo.Echo) *http.Request
		ErrorBodyResponse   errorResult
		successBodyResponse successResult
	}{
		{
			name:         "error; Invalid Input",
			ExpectedCode: 400,
			setupMock:    func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/subtask/a", nil)
				return req
			},
			ErrorBodyResponse: errorResult{
				Error: "Invalid input",
			},
		},
		{
			name:         "error: Failed to load subtasks",
			ExpectedCode: 500,
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().FindSubTaskByTaskID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(subtasks, errors.New("mock error"))
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/subtask/123", nil)
				return req
			},
			ErrorBodyResponse: errorResult{
				Error: "Failed to load subtasks",
			},
		},
		{
			name:         "Success",
			ExpectedCode: 200,
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().FindSubTaskByTaskID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(subtasks, nil)
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/subtask/123", nil)
				return req
			},
			successBodyResponse: successResult{
				Status: "success",
				Data:   subtasks,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			e := echo.New()
			req := tc.setRequest(e)
			rec := httptest.NewRecorder()

			//Instantiate mocsk
			mockServices := mocks.NewMockService(ctr)
			mockTaskService := mocks.NewMockTaskService(ctr)
			tc.setupMock(mockServices, mockTaskService)

			//Instantiate handlers and define route
			handler := handlers.NewHandlers(mockServices)
			e.GET("/subtask/:parentID", handler.GetAllSubListByParentID)
			tc.setRequest(e)

			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.ExpectedCode, rec.Code)
			if tc.ExpectedCode != 200 {
				var Error errorResult
				json.Unmarshal(rec.Body.Bytes(), &Error)
				assert.Equal(t, tc.ErrorBodyResponse, Error)
			} else {
				var bodyResult successResult
				json.Unmarshal(rec.Body.Bytes(), &bodyResult)
				assert.Equal(t, tc.successBodyResponse, bodyResult)
			}

		})
	}

}

func TestDownloadFile(t *testing.T) {
	const fileName = "test_file.png"
	const filePath = "../../../static/test_file.png"
	testCases := []struct {
		name          string
		setupMock     func(mockService *mocks.MockService, mockFileService *mocks.MockFileService)
		setRequest    func(e *echo.Echo) *http.Request
		expectedCode  int
		expectedError errorResult
	}{
		{
			name: "error: no content",
			setupMock: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {
				mockFileService.EXPECT().Download(gomock.Any()).Return("", errors.New("mocks error"))
				mockService.EXPECT().FileService().Return(mockFileService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/downlods", nil)
				return req
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name: "success",
			setupMock: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {
				mockFileService.EXPECT().Download(gomock.Any()).Return(filePath, nil)
				mockService.EXPECT().FileService().Return(mockFileService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				q := make(url.Values)
				q.Set("fileName", fileName)
				req := httptest.NewRequest(http.MethodGet, "/downlods?"+q.Encode(), nil)
				return req
			},
			expectedCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			e := echo.New()
			req := tc.setRequest(e)
			rec := httptest.NewRecorder()

			//Instantiate mocks
			mockService := mocks.NewMockService(ctr)
			mockFileService := mocks.NewMockFileService(ctr)
			tc.setupMock(mockService, mockFileService)

			//instantiate handlers
			handler := handlers.NewHandlers(mockService)
			e.GET("/downlods", handler.DownloadFile)
			tc.setRequest(e)

			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}

}
