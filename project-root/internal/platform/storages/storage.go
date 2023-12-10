package storages

import "mime/multipart"

type Storage interface {
	SaveFile(file *multipart.FileHeader) (string, error)
	LoadFile(fileName string) (interface{}, error)
	DeleteFile(fileName string) error
}
