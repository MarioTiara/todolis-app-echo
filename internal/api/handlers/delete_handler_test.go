package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/mocks"
	"github.com/stretchr/testify/assert"
)

type errorResult struct {
	Error string `json:"error"`
}

func TestDeleteTask(t *testing.T) {
	testCases := []struct {
		name               string
		mockSetup          func(mockServices *mocks.MockService, mockTaskServices *mocks.MockTaskService, mockFileServices *mocks.MockFileService)
		setRequest         func(e *echo.Echo) *http.Request
		expectedResponse   int
		expectedResultBody errorResult
	}{
		{
			name: "Error - Invalid Input",
			mockSetup: func(mockServices *mocks.MockService, mockTaskServices *mocks.MockTaskService, mockFileServices *mocks.MockFileService) {

			},
			setRequest: func(e *echo.Echo) *http.Request {

				req := httptest.NewRequest(http.MethodDelete, "/tasks/a", nil)
				return req
			},
			expectedResponse: 400,
			expectedResultBody: errorResult{
				Error: "Invalid Input",
			},
		},
		{
			name: "Error - Failed to delete Task",
			mockSetup: func(mockServices *mocks.MockService, mockTaskServices *mocks.MockTaskService, mockFileServices *mocks.MockFileService) {
				mockTaskServices.EXPECT().Delete(gomock.Any()).AnyTimes().Return(errors.New("mock error"))
				mockServices.EXPECT().TaskService().Return(mockTaskServices)

			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/tasks/123", nil)
				return req
			},
			expectedResponse: 500,
			expectedResultBody: errorResult{
				Error: "Failed to delete Task",
			},
		},
		{
			name: "Error - Failed to delete files",
			mockSetup: func(mockServices *mocks.MockService, mockTaskServices *mocks.MockTaskService, mockFileServices *mocks.MockFileService) {
				mockTaskServices.EXPECT().Delete(gomock.Any()).AnyTimes().Return(nil)
				mockFileServices.EXPECT().DeleteByTaskID(gomock.Any()).AnyTimes().Return(errors.New("mock errors"))
				mockServices.EXPECT().TaskService().Return(mockTaskServices)
				mockServices.EXPECT().FileService().Return(mockFileServices)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/tasks/123", nil)
				return req
			},
			expectedResponse: 500,
			expectedResultBody: errorResult{
				Error: "Failed to delete files",
			},
		},
		{
			name: "Success - No Content ",
			mockSetup: func(mockServices *mocks.MockService, mockTaskServices *mocks.MockTaskService, mockFileServices *mocks.MockFileService) {
				mockTaskServices.EXPECT().Delete(gomock.Any()).AnyTimes().Return(nil)
				mockFileServices.EXPECT().DeleteByTaskID(gomock.Any()).AnyTimes().Return(nil)
				mockServices.EXPECT().TaskService().Return(mockTaskServices)
				mockServices.EXPECT().FileService().Return(mockFileServices)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/tasks/123", nil)
				return req
			},
			expectedResponse: http.StatusNoContent,
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
			mockFileService := mocks.NewMockFileService(ctr)
			tc.mockSetup(mockService, mockTaskService, mockFileService)

			//Instantiate handlers
			handler := handlers.NewHandlers(mockService)
			e.DELETE("/tasks/:id", handler.DeleteTask)
			tc.setRequest(e)
			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedResponse, rec.Code)
			if tc.expectedResponse != http.StatusNoContent {
				var Error errorResult
				json.Unmarshal(rec.Body.Bytes(), &Error)
				assert.Equal(t, tc.expectedResultBody, Error)
			}
		})
	}
}

func TestDeleteFile(t *testing.T) {
	testCases := []struct {
		name               string
		mockSetup          func(mockService *mocks.MockService, mockFileService *mocks.MockFileService)
		setRequest         func(e *echo.Echo) *http.Request
		expectedResponse   int
		expectedResultBody errorResult
	}{
		{
			name: "Eror: Invalid Input",
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/files/a", nil)
				return req
			},
			mockSetup:        func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {},
			expectedResponse: 400,
			expectedResultBody: errorResult{
				Error: "Invalid input",
			},
		},
		{
			name: "Eror: Failed to delete file",
			mockSetup: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {
				mockFileService.EXPECT().DeleteByID(gomock.Any()).Return(errors.New("mock error"))
				mockService.EXPECT().FileService().Return(mockFileService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/files/123", nil)
				return req
			},
			expectedResponse: 500,
			expectedResultBody: errorResult{
				Error: "Failed to delete file",
			},
		},
		{
			name: "Success: No content",
			mockSetup: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {
				mockFileService.EXPECT().DeleteByID(gomock.Any()).Return(nil)
				mockService.EXPECT().FileService().Return(mockFileService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/files/123", nil)
				return req
			},
			expectedResponse: http.StatusNoContent,
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
			mockServcie := mocks.NewMockService(ctr)
			mockFileService := mocks.NewMockFileService(ctr)
			tc.mockSetup(mockServcie, mockFileService)

			//Instantiate handlers
			handler := handlers.NewHandlers(mockServcie)
			e.DELETE("/files/:id", handler.DeleteFile)
			tc.setRequest(e)

			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedResponse, rec.Code)
			if tc.expectedResponse != http.StatusNoContent {
				var Error errorResult
				json.Unmarshal(rec.Body.Bytes(), &Error)
				assert.Equal(t, tc.expectedResultBody, Error)
			}

		})
	}
}
