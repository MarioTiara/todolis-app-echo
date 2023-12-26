package dtos

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSubTaskRequest_JSONBinding(t *testing.T) {
	// Create a sample AddSubTaskRequest
	requestBody := `{"title": "Sample Task", "description": "Task Description", "priority": 1, "parent_id": 123}`

	// Create a new HTTP request with the sample JSON body
	req, err := http.NewRequest("POST", "/add-subtask", strings.NewReader(requestBody))
	assert.NoError(t, err, "Failed to create HTTP request")

	// Parse the JSON request body into the AddSubTaskRequest struct
	var addSubTaskReq AddSubTaskRequest
	err = json.NewDecoder(req.Body).Decode(&addSubTaskReq)

	// Assertions
	assert.NoError(t, err, "Error decoding JSON")
	assert.Equal(t, "Sample Task", addSubTaskReq.Title, "Title should match")
	assert.Equal(t, "Task Description", addSubTaskReq.Description, "Description should match")
	assert.Equal(t, 1, addSubTaskReq.Priority, "Priority should match")
	assert.Equal(t, uint(123), addSubTaskReq.ParentID, "ParentID should match")
}
