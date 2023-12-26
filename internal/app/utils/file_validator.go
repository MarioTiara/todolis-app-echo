package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

var allowedExtensions = []string{".text", ".pdf", ".txt"}

func IsValidFileExtension(file *multipart.FileHeader) bool {

	ext := strings.ToLower(filepath.Ext(file.Filename))
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}

	return false
}
