package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	datafake "github.com/marioTiara/todolistapi/data-fake"
	"github.com/marioTiara/todolistapi/internal/app/models"
	"github.com/marioTiara/todolistapi/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	inputUpdate := datafake.GenerateUpdateTaskRequest()
	type successBody struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	testCases := []struct {
		name          string
		setupMock     func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService)
		setRequest    func(e *echo.Echo) *http.Request
		expectedCode  int
		expectedError errorResult
		successBody   successBody
	}{
		{
			name: "error:failed to update",
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().Update(gomock.Any()).Return(models.Task{}, errors.New("mock error"))
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				inputJson, err := json.Marshal(inputUpdate)
				assert.NoError(t, err)
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(string(inputJson)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			expectedCode: 500,
			expectedError: errorResult{
				Error: "failed to update",
			},
		},
		{
			name: "success update",
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().Update(gomock.Any()).Return(models.Task{}, nil)
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				inputJson, err := json.Marshal(inputUpdate)
				assert.NoError(t, err)
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(string(inputJson)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			expectedCode: http.StatusOK,
			successBody: successBody{
				Status:  "success",
				Message: fmt.Sprintf("%d updated", inputUpdate.ID),
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

			//Instantiate handler
			handler := NewHandlers(mockService)
			e.PUT("/tasks", handler.Update)
			req := tc.setRequest(e)
			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedCode, rec.Code)
			if tc.expectedCode == 200 {
				var result successBody
				json.Unmarshal(rec.Body.Bytes(), &result)
				assert.Equal(t, tc.successBody, result)
			} else {
				var err errorResult
				json.Unmarshal(rec.Body.Bytes(), &err)
				assert.Equal(t, tc.expectedError, err)
			}

		})
	}
}
