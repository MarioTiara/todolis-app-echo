package dtos

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileQueryModel_JSONEncodingDecoding(t *testing.T) {
	// Create a sample FileQueryModel instance
	now := time.Now().UTC()
	fileQueryModel := FileQueryModel{
		ID:           1,
		FileName:     "example.txt",
		FileSize:     1024,
		FileURL:      "/files/example.txt",
		UploadedTime: now,
	}

	// Marshal the FileQueryModel to JSON
	jsonData, err := json.Marshal(fileQueryModel)
	assert.NoError(t, err, "Error encoding JSON")

	// Unmarshal the JSON data into a new FileQueryModel instance
	var decodedFileQueryModel FileQueryModel
	err = json.Unmarshal(jsonData, &decodedFileQueryModel)
	assert.NoError(t, err, "Error decoding JSON")

	// Assertions
	assert.Equal(t, fileQueryModel.ID, decodedFileQueryModel.ID, "ID should match")
	assert.Equal(t, fileQueryModel.FileName, decodedFileQueryModel.FileName, "FileName should match")
	assert.Equal(t, fileQueryModel.FileSize, decodedFileQueryModel.FileSize, "FileSize should match")
	assert.Equal(t, fileQueryModel.FileURL, decodedFileQueryModel.FileURL, "FileURL should match")
	assert.Equal(t, fileQueryModel.UploadedTime, decodedFileQueryModel.UploadedTime, "UploadedTime should match")
}
