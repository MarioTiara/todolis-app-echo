package storages

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"

	"github.com/marioTiara/todolistapi/internal/app/utils"
)

type localStorage struct {
	UploadsDir string
}

func NewLocalStoarge(uploadsDir string) Storage {
	return &localStorage{uploadsDir}
}

func (s *localStorage) SaveFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if _, err := os.Stat(s.UploadsDir); os.IsNotExist(err) {
		err := os.Mkdir(s.UploadsDir, 0755)
		if err != nil {
			return "", err
		}
	}

	filename := fmt.Sprintf("%s-%s%s", utils.RandomString(10), removeExtension(file.Filename), filepath.Ext(file.Filename))
	dest, err := os.Create(s.UploadsDir + "/" + filename)
	if err != nil {
		return "", err
	}

	defer dest.Close()
	_, err = io.Copy(dest, src)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (s *localStorage) LoadFile(fileName string) (interface{}, error) {
	return fmt.Sprintf("%s/%s", s.UploadsDir, fileName), nil
}

func (s *localStorage) Path() string {
	return s.UploadsDir
}
func (s *localStorage) DeleteFile(fileName string) error {
	err := os.Remove(s.UploadsDir + "/" + fileName)
	return err
}
func removeExtension(fileName string) string {
	return fileName[:len(fileName)-len(path.Ext(fileName))]
}
