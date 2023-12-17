package dtos

import "time"

type FileQueryModel struct {
	ID           uint      `json:"id"`
	FileName     string    `json:"file_name"`
	FileSize     uint      `json:"file_size"`
	FileURL      string    `json:"file_url"`
	UploadedTime time.Time `json:"uploaded_time"`
}
