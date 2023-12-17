package utils

import (
	"testing"

	datafake "github.com/marioTiara/todolistapp/data-fake"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/stretchr/testify/assert"
)

func TestConvertTaskToQueryModel(t *testing.T) {
	//Arrange
	task := datafake.GenerateSubtask()
	task.Children = append(task.Children, datafake.GenerateSubtask())
	task.Files = datafake.GenerateFilesList(1)

	expectedResult := dtos.TaskQueryModel{}
	expectedResult.ID = task.ID
	expectedResult.Title = task.Title
	expectedResult.Description = task.Description
	expectedResult.Priority = uint(task.Priority)
	expectedResult.CreatedAt = task.CreatedAt
	expectedResult.Checked = task.Checked
	expectedResult.IsActive = task.IsActive
	for _, file := range task.Files {
		queryFile := ConverFileToFileQueryModel(file)
		expectedResult.Files = append(expectedResult.Files, queryFile)
	}

	for _, subTask := range task.Children {
		subtaskQuery := ConvertSubTaskToSubtaskQueryModel(subTask)
		expectedResult.SubTasks = append(expectedResult.SubTasks, subtaskQuery)
	}

	//Act
	actual := ConvertTaskToQueryModel(task)

	//Assert
	assert.Equal(t, expectedResult, actual)
}

func TestConverFileToFileQueryModel(t *testing.T) {
	//Arrange
	file := datafake.GenerateFile()
	queryModel := dtos.FileQueryModel{
		ID:           file.ID,
		FileName:     file.FileName,
		FileSize:     file.FileSize,
		FileURL:      file.FileURL,
		UploadedTime: file.CreatedAt,
	}

	//Act
	actual := ConverFileToFileQueryModel(file)

	//Assert
	assert.Equal(t, queryModel, actual)

}

func TestConvertSubTaskToSubtaskQueryModel(t *testing.T) {
	//Arrange
	subTask := datafake.GenerateSubtask()
	subTask.Files = append(subTask.Files, datafake.GenerateFile())

	subTaskQueryModel := dtos.SubtaskQueryModel{
		ID:          subTask.ID,
		Title:       subTask.Title,
		Description: subTask.Description,
		CreatedAt:   subTask.CreatedAt,
		UpdatedAt:   subTask.UpdatedAt,
		Priority:    uint(subTask.Priority),
		Checked:     subTask.Checked,
		IsActive:    subTask.IsActive,
		ParentID:    *subTask.ParentID,
	}
	for _, file := range subTask.Files {
		fileQuery := ConverFileToFileQueryModel(file)
		subTaskQueryModel.Files = append(subTaskQueryModel.Files, fileQuery)
	}

	//Actual
	actual := ConvertSubTaskToSubtaskQueryModel(subTask)

	//Assert
	assert.Equal(t, subTaskQueryModel, actual)

}

func TestConvertRequestToTaskEntity(t *testing.T) {
	//Arrange
	request := datafake.GenerateAddTaskRequest(2)
	newtask := models.Task{Title: request.Title, Description: request.Description}
	for _, child := range request.Childrens {
		newtask.Children = append(newtask.Children, models.Task{Title: child.Title, Description: child.Description})
	}

	//Act
	actual := ConvertRequestToTaskEntity(request)

	//Assert
	assert.Equal(t, newtask, actual)
}

func TestConvertSubTaskRequestToTaskEntity(t *testing.T) {
	//Arrange
	request := datafake.GenerateAddSubTaskRequest(1)
	var subTask = models.Task{Title: request.Title, Description: request.Description, ParentID: &request.ParentID}

	//Actual
	actual := ConvertSubTaskRequestToTaskEntity(request)

	//Assert
	assert.Equal(t, subTask, actual)
}
