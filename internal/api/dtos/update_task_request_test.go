// update_task_request_test.go

package dtos

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTaskRequest_JSONEncodingDecoding(t *testing.T) {
	// Create a sample UpdateTaskRequest instance
	updateTaskRequest := UpdateTaskRequest{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated Task Description",
		Priority:    3,
		Checked:     true,
		ParentID:    nil,
	}

	// Marshal the UpdateTaskRequest to JSON
	jsonData, err := json.Marshal(updateTaskRequest)
	assert.NoError(t, err, "Error encoding JSON")

	// Unmarshal the JSON data into a new UpdateTaskRequest instance
	var decodedUpdateTaskRequest UpdateTaskRequest
	err = json.Unmarshal(jsonData, &decodedUpdateTaskRequest)
	assert.NoError(t, err, "Error decoding JSON")

	// Assertions
	assert.Equal(t, updateTaskRequest.ID, decodedUpdateTaskRequest.ID, "ID should match")
	assert.Equal(t, updateTaskRequest.Title, decodedUpdateTaskRequest.Title, "Title should match")
	assert.Equal(t, updateTaskRequest.Description, decodedUpdateTaskRequest.Description, "Description should match")
	assert.Equal(t, updateTaskRequest.Priority, decodedUpdateTaskRequest.Priority, "Priority should match")
	assert.Equal(t, updateTaskRequest.Checked, decodedUpdateTaskRequest.Checked, "Checked should match")
	assert.Nil(t, decodedUpdateTaskRequest.ParentID, "ParentID should be nil")
}

// Add more test cases as needed...
