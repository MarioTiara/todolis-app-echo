package dtos

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskQueryModel_JSONEncodingDecoding(t *testing.T) {
	// Create a sample TaskQueryModel instance
	now := time.Now().UTC()
	taskQueryModel := TaskQueryModel{
		ID:          1,
		Title:       "Sample Task",
		Description: "Task Description",
		CreatedAt:   now,
		UpdatedAt:   now,
		Priority:    2,
		Checked:     false,
		IsActive:    true,
		Files: []FileQueryModel{
			{
				ID:           1,
				FileName:     "example.txt",
				FileSize:     1024,
				FileURL:      "/files/example.txt",
				UploadedTime: now,
			},
		},
		SubTasks: []SubtaskQueryModel{
			{
				ID:          1,
				Title:       "Sample Subtask",
				Description: "Subtask Description",
				CreatedAt:   now,
				UpdatedAt:   now,
				Priority:    2,
				Checked:     false,
				IsActive:    true,
				Files: []FileQueryModel{
					{
						ID:           2,
						FileName:     "subtask_file.txt",
						FileSize:     512,
						FileURL:      "/files/subtask_file.txt",
						UploadedTime: now,
					},
				},
				ParentID: 1,
			},
		},
	}

	// Marshal the TaskQueryModel to JSON
	jsonData, err := json.Marshal(taskQueryModel)
	assert.NoError(t, err, "Error encoding JSON")

	// Unmarshal the JSON data into a new TaskQueryModel instance
	var decodedTaskQueryModel TaskQueryModel
	err = json.Unmarshal(jsonData, &decodedTaskQueryModel)
	assert.NoError(t, err, "Error decoding JSON")

	// Assertions
	assert.Equal(t, taskQueryModel.ID, decodedTaskQueryModel.ID, "ID should match")
	assert.Equal(t, taskQueryModel.Title, decodedTaskQueryModel.Title, "Title should match")
	assert.Equal(t, taskQueryModel.Description, decodedTaskQueryModel.Description, "Description should match")
	assert.Equal(t, taskQueryModel.CreatedAt, decodedTaskQueryModel.CreatedAt, "CreatedAt should match")
	assert.Equal(t, taskQueryModel.UpdatedAt, decodedTaskQueryModel.UpdatedAt, "UpdatedAt should match")
	assert.Equal(t, taskQueryModel.Priority, decodedTaskQueryModel.Priority, "Priority should match")
	assert.Equal(t, taskQueryModel.Checked, decodedTaskQueryModel.Checked, "Checked should match")
	assert.Equal(t, taskQueryModel.IsActive, decodedTaskQueryModel.IsActive, "IsActive should match")
	assert.Len(t, decodedTaskQueryModel.Files, 1, "Files length should match")
	assert.Len(t, decodedTaskQueryModel.SubTasks, 1, "SubTasks length should match")

	// Check properties of the first file
	assert.Equal(t, taskQueryModel.Files[0].ID, decodedTaskQueryModel.Files[0].ID, "File ID should match")
	assert.Equal(t, taskQueryModel.Files[0].FileName, decodedTaskQueryModel.Files[0].FileName, "File FileName should match")
	assert.Equal(t, taskQueryModel.Files[0].FileSize, decodedTaskQueryModel.Files[0].FileSize, "File FileSize should match")
	assert.Equal(t, taskQueryModel.Files[0].FileURL, decodedTaskQueryModel.Files[0].FileURL, "File FileURL should match")
	assert.Equal(t, taskQueryModel.Files[0].UploadedTime, decodedTaskQueryModel.Files[0].UploadedTime, "File UploadedTime should match")

	// Check properties of the first subtask
	assert.Equal(t, taskQueryModel.SubTasks[0].ID, decodedTaskQueryModel.SubTasks[0].ID, "Subtask ID should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Title, decodedTaskQueryModel.SubTasks[0].Title, "Subtask Title should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Description, decodedTaskQueryModel.SubTasks[0].Description, "Subtask Description should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].CreatedAt, decodedTaskQueryModel.SubTasks[0].CreatedAt, "Subtask CreatedAt should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].UpdatedAt, decodedTaskQueryModel.SubTasks[0].UpdatedAt, "Subtask UpdatedAt should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Priority, decodedTaskQueryModel.SubTasks[0].Priority, "Subtask Priority should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Checked, decodedTaskQueryModel.SubTasks[0].Checked, "Subtask Checked should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].IsActive, decodedTaskQueryModel.SubTasks[0].IsActive, "Subtask IsActive should match")
	assert.Len(t, decodedTaskQueryModel.SubTasks[0].Files, 1, "Subtask Files length should match")

	// Check properties of the first file in the subtask
	assert.Equal(t, taskQueryModel.SubTasks[0].Files[0].ID, decodedTaskQueryModel.SubTasks[0].Files[0].ID, "Subtask File ID should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Files[0].FileName, decodedTaskQueryModel.SubTasks[0].Files[0].FileName, "Subtask File FileName should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Files[0].FileSize, decodedTaskQueryModel.SubTasks[0].Files[0].FileSize, "Subtask File FileSize should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Files[0].FileURL, decodedTaskQueryModel.SubTasks[0].Files[0].FileURL, "Subtask File FileURL should match")
	assert.Equal(t, taskQueryModel.SubTasks[0].Files[0].UploadedTime, decodedTaskQueryModel.SubTasks[0].Files[0].UploadedTime, "Subtask File UploadedTime should match")
}

// Add more test cases as needed...
