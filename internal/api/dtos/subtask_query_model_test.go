package dtos

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubtaskQueryModel_JSONEncodingDecoding(t *testing.T) {
	// Create a sample SubtaskQueryModel instance
	now := time.Now().UTC()
	subtaskQueryModel := SubtaskQueryModel{
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
				ID:           1,
				FileName:     "example.txt",
				FileSize:     1024,
				FileURL:      "/files/example.txt",
				UploadedTime: now,
			},
		},
		ParentID: 123,
	}

	// Marshal the SubtaskQueryModel to JSON
	jsonData, err := json.Marshal(subtaskQueryModel)
	assert.NoError(t, err, "Error encoding JSON")

	// Unmarshal the JSON data into a new SubtaskQueryModel instance
	var decodedSubtaskQueryModel SubtaskQueryModel
	err = json.Unmarshal(jsonData, &decodedSubtaskQueryModel)
	assert.NoError(t, err, "Error decoding JSON")

	// Assertions
	assert.Equal(t, subtaskQueryModel.ID, decodedSubtaskQueryModel.ID, "ID should match")
	assert.Equal(t, subtaskQueryModel.Title, decodedSubtaskQueryModel.Title, "Title should match")
	assert.Equal(t, subtaskQueryModel.Description, decodedSubtaskQueryModel.Description, "Description should match")
	assert.Equal(t, subtaskQueryModel.CreatedAt, decodedSubtaskQueryModel.CreatedAt, "CreatedAt should match")
	assert.Equal(t, subtaskQueryModel.UpdatedAt, decodedSubtaskQueryModel.UpdatedAt, "UpdatedAt should match")
	assert.Equal(t, subtaskQueryModel.Priority, decodedSubtaskQueryModel.Priority, "Priority should match")
	assert.Equal(t, subtaskQueryModel.Checked, decodedSubtaskQueryModel.Checked, "Checked should match")
	assert.Equal(t, subtaskQueryModel.IsActive, decodedSubtaskQueryModel.IsActive, "IsActive should match")
	assert.Len(t, decodedSubtaskQueryModel.Files, 1, "Files length should match")

	// Check properties of the first file
	assert.Equal(t, subtaskQueryModel.Files[0].ID, decodedSubtaskQueryModel.Files[0].ID, "File ID should match")
	assert.Equal(t, subtaskQueryModel.Files[0].FileName, decodedSubtaskQueryModel.Files[0].FileName, "File FileName should match")
	assert.Equal(t, subtaskQueryModel.Files[0].FileSize, decodedSubtaskQueryModel.Files[0].FileSize, "File FileSize should match")
	assert.Equal(t, subtaskQueryModel.Files[0].FileURL, decodedSubtaskQueryModel.Files[0].FileURL, "File FileURL should match")
	assert.Equal(t, subtaskQueryModel.Files[0].UploadedTime, decodedSubtaskQueryModel.Files[0].UploadedTime, "File UploadedTime should match")

	assert.Equal(t, subtaskQueryModel.ParentID, decodedSubtaskQueryModel.ParentID, "ParentID should match")
}

// Add more test cases as needed...
