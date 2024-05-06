package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileService interface {
	Upload(file *multipart.FileHeader) (string, error)
}

type fileService struct {
}

func NewFileService() FileService {
	return &fileService{}
}

func (fs *fileService) Upload(file *multipart.FileHeader) (string, error) {
	fileName := filepath.Base(file.Filename)
	extension := filepath.Ext(fileName)

	uniqueName := generateUniqueName() + extension
	destination := filepath.Join(".", "images", uniqueName)

	if err := os.MkdirAll(filepath.Join(".", "images"), os.ModePerm); err != nil {
		return "", err
	}

	if err := saveUploadedFile(file, destination); err != nil {
		return "", err
	}

	return uniqueName, nil
}

func generateUniqueName() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func saveUploadedFile(file *multipart.FileHeader, destination string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	_, err = out.ReadFrom(src)
	return err
}
