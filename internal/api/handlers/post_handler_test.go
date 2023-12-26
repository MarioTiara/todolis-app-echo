package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	datafake "github.com/marioTiara/todolistapp/data-fake"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/utils"
	"github.com/marioTiara/todolistapp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPostTaskHandler(t *testing.T) {
	type successBodyResult struct {
		Status string      `json:"status"`
		Data   models.Task `json:"data"`
	}
	inputTask := datafake.GenerateAddTaskRequest(1)
	createdtask := utils.ConvertRequestToTaskEntity(inputTask)
	testCases := []struct {
		name           string
		setupMock      func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService)
		setRequest     func(e *echo.Echo) *http.Request
		expectedCode   int
		expectedError  errorResult
		expectedResult successBodyResult
	}{
		{
			name: "error: failed create new task",
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().Create(gomock.Any()).Return(createdtask, errors.New("mock error"))
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				inputJson, err := json.Marshal(inputTask)
				assert.NoError(t, err)
				req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(string(inputJson)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			expectedCode: 500,
			expectedError: errorResult{
				Error: "Failed to create task",
			},
		},
		{
			name: "succcess",
			setupMock: func(mockService *mocks.MockService, mockTaskService *mocks.MockTaskService) {
				mockTaskService.EXPECT().Create(gomock.Any()).Return(createdtask, nil)
				mockService.EXPECT().TaskService().Return(mockTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				inputJson, err := json.Marshal(inputTask)
				assert.NoError(t, err)
				req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(string(inputJson)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			expectedCode: 201,
			expectedResult: successBodyResult{
				Status: "success",
				Data:   createdtask,
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
			handler := handlers.NewHandlers(mockService)
			e.POST("/tasks", handler.PostTaskHandler)
			req := tc.setRequest(e)
			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedCode, rec.Code)
			if tc.expectedCode == 201 {
				var result successBodyResult
				json.Unmarshal(rec.Body.Bytes(), &result)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}

func TestPostSubTaskHandler(t *testing.T) {
	type successBody struct {
		Status string `json:"status"`
		ID     uint   `json:"ID"`
	}
	inputSbTask := datafake.GenerateAddSubTaskRequest(1)
	createdTask := utils.ConvertSubTaskRequestToTaskEntity(inputSbTask)
	testCases := []struct {
		name                string
		setupMock           func(mockService *mocks.MockService, mocksTaskService *mocks.MockTaskService)
		setRequest          func(e *echo.Echo) *http.Request
		expectedCode        int
		expectedError       errorResult
		expectedSuccessBody successBody
	}{
		{
			name: "error: Failed to create task",
			setupMock: func(mockService *mocks.MockService, mocksTaskService *mocks.MockTaskService) {
				mocksTaskService.EXPECT().CreateSubTask(gomock.Any()).Return(createdTask, errors.New("mock error"))
				mockService.EXPECT().TaskService().Return(mocksTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				inputJson, err := json.Marshal(inputSbTask)
				assert.NoError(t, err)
				req := httptest.NewRequest(http.MethodPost, "/subtasks", strings.NewReader(string(inputJson)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			expectedCode: 500,
			expectedError: errorResult{
				Error: "Failed to create task",
			},
		},
		{
			name: "success",
			setupMock: func(mockService *mocks.MockService, mocksTaskService *mocks.MockTaskService) {
				mocksTaskService.EXPECT().CreateSubTask(gomock.Any()).Return(createdTask, nil)
				mockService.EXPECT().TaskService().Return(mocksTaskService)
			},
			setRequest: func(e *echo.Echo) *http.Request {
				inputJson, err := json.Marshal(inputSbTask)
				assert.NoError(t, err)
				req := httptest.NewRequest(http.MethodPost, "/subtasks", strings.NewReader(string(inputJson)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			expectedCode: 201,
			expectedSuccessBody: successBody{
				Status: "success",
				ID:     createdTask.ID,
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
			handler := handlers.NewHandlers(mockService)
			e.POST("/subtasks", handler.PostSubTaskHandler)
			req := tc.setRequest(e)
			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedCode, rec.Code)
			if tc.expectedCode == 201 {
				var result successBody
				json.Unmarshal(rec.Body.Bytes(), &result)
				assert.Equal(t, tc.expectedSuccessBody, result)
			} else {
				var err errorResult
				json.Unmarshal(rec.Body.Bytes(), &err)
				assert.Equal(t, err, tc.expectedError)
			}
		})
	}
}

func TestUploadTaskFileHandler(t *testing.T) {
	fakeFileDetails := datafake.GenerateFileQueryModel()

	type successBody struct {
		Status string                `json:"status"`
		Data   []dtos.FileQueryModel `json:"data"`
	}
	testCases := []struct {
		name          string
		setupMock     func(mockService *mocks.MockService, mockFileService *mocks.MockFileService)
		setRequest    func(e *echo.Echo) *http.Request
		expectedCode  int
		expectedError errorResult
		successBody   successBody
	}{
		{
			name:      "error: Invalid Parameter",
			setupMock: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {},
			setRequest: func(e *echo.Echo) *http.Request {
				f := make(url.Values)
				f.Set("taskID", "a")
				req := httptest.NewRequest(http.MethodPost, "/upload", strings.NewReader(f.Encode()))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
				return req
			},
			expectedCode: 400,
			expectedError: errorResult{
				Error: "Invalid Parameter",
			},
		},
		{
			name:      "BadRequest: invalid extension",
			setupMock: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {},
			setRequest: func(e *echo.Echo) *http.Request {
				buffer, writer := createFileBuffer(t, "../../../static/test_file.png")

				req := httptest.NewRequest(http.MethodPost, "/upload", buffer)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				return req
			},
			expectedCode: http.StatusBadRequest,
			expectedError: errorResult{
				Error: "file extension should in .text or .pdf",
			},
		},
		{
			name: "error: failed save file",
			setupMock: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {
				mockFileService.EXPECT().SaveFile(gomock.Any(), gomock.Any()).Return(fakeFileDetails, errors.New("mock error"))
				mockService.EXPECT().FileService().Return(mockFileService)

			},
			setRequest: func(e *echo.Echo) *http.Request {
				buffer, writer := createFileBuffer(t, "../../../static/test_file.txt")

				req := httptest.NewRequest(http.MethodPost, "/upload", buffer)
				req.Header.Set("Content-Type", writer.FormDataContentType())

				return req
			},
			expectedCode: 500,
			expectedError: errorResult{
				Error: "failed to save file",
			},
		},
		{
			name: "success",
			setupMock: func(mockService *mocks.MockService, mockFileService *mocks.MockFileService) {
				mockFileService.EXPECT().SaveFile(gomock.Any(), gomock.Any()).Return(fakeFileDetails, nil)
				mockService.EXPECT().FileService().Return(mockFileService)

			},
			setRequest: func(e *echo.Echo) *http.Request {
				buffer, writer := createFileBuffer(t, "../../../static/test_file.txt")

				req := httptest.NewRequest(http.MethodPost, "/upload", buffer)
				req.Header.Set("Content-Type", writer.FormDataContentType())

				return req
			},
			expectedCode: 201,
			successBody: successBody{
				Status: "success",
				Data: []dtos.FileQueryModel{
					fakeFileDetails,
				},
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
			mockTaskService := mocks.NewMockFileService(ctr)
			tc.setupMock(mockService, mockTaskService)

			//Instantiate handler
			handler := handlers.NewHandlers(mockService)
			e.POST("/upload", handler.UploadTaskFilesHandler)
			req := tc.setRequest(e)
			//Act
			e.ServeHTTP(rec, req)

			//Assert
			assert.Equal(t, tc.expectedCode, rec.Code)
			if tc.expectedCode != 201 {
				var err errorResult
				json.Unmarshal(rec.Body.Bytes(), &err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				var result successBody
				json.Unmarshal(rec.Body.Bytes(), &result)
				assert.Equal(t, tc.successBody, result)
			}
		})
	}
}

func createFileBuffer(t *testing.T, filePath string) (*bytes.Buffer, *multipart.Writer) {
	var buffer bytes.Buffer

	//Create a multipart writer
	writer := multipart.NewWriter(&buffer)
	//Add text field
	writer.WriteField("taskID", "1")

	// Add file field
	fileField, err := writer.CreateFormFile("files", filePath)
	if err != nil {
		t.Fatal("Error creating form file:", err)
	}

	// Open and copy the file content to the form file field
	fileContent := []byte("This is the content of the file.")

	_, err = fileField.Write(fileContent)
	if err != nil {
		t.Fatal("Error writing file content:", err)
	}

	// Close the multipart writer
	writer.Close()
	return &buffer, writer
}
