package dtos

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTaskRequest_JSONBinding(t *testing.T) {
	// Create a sample AddTaskRequest with children
	requestBody := `{
		"title": "Parent Task",
		"description": "Parent Task Description",
		"priority": 1,
		"childrens": [
			{"title": "Child Task 1", "description": "Child Task 1 Description", "priority": 2, "childrens": []},
			{"title": "Child Task 2", "description": "Child Task 2 Description", "priority": 3, "childrens": []}
		]
	}`

	// Create a new HTTP request with the sample JSON body
	req, err := http.NewRequest("POST", "/add-task", strings.NewReader(requestBody))
	assert.NoError(t, err, "Failed to create HTTP request")

	// Parse the JSON request body into the AddTaskRequest struct
	var addTaskReq AddTaskRequest
	err = json.NewDecoder(req.Body).Decode(&addTaskReq)

	// Assertions
	assert.NoError(t, err, "Error decoding JSON")
	assert.Equal(t, "Parent Task", addTaskReq.Title, "Title should match")
	assert.Equal(t, "Parent Task Description", addTaskReq.Description, "Description should match")
	assert.Equal(t, 1, addTaskReq.Priority, "Priority should match")

	// Check children
	assert.Len(t, addTaskReq.Childrens, 2, "Expected two children")

	// Check the properties of the first child
	assert.Equal(t, "Child Task 1", addTaskReq.Childrens[0].Title, "Child Title should match")
	assert.Equal(t, "Child Task 1 Description", addTaskReq.Childrens[0].Description, "Child Description should match")
	assert.Equal(t, 2, addTaskReq.Childrens[0].Priority, "Child Priority should match")

	// Check the properties of the second child
	assert.Equal(t, "Child Task 2", addTaskReq.Childrens[1].Title, "Child Title should match")
	assert.Equal(t, "Child Task 2 Description", addTaskReq.Childrens[1].Description, "Child Description should match")
	assert.Equal(t, 3, addTaskReq.Childrens[1].Priority, "Child Priority should match")
}
