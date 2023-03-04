package domain

import (
	"io"

	"github.com/julianojj/aurora/internal/core/exceptions"
)

type File struct {
	FileID   string
	Name     string
	FileType string
	Size     int64
	Reader   io.Reader
	Path     string
}

func NewFile(fileID, name, fileType, path string, size int64, reader io.Reader) (*File, error) {
	file := &File{
		FileID:   fileID,
		Name:     name,
		FileType: fileType,
		Path:     path,
		Size:     size,
		Reader:   reader,
	}
	err := file.Validate()
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *File) Validate() error {
	if f.FileID == "" {
		return exceptions.NewValidationException("File ID cannot be empty")
	}
	if f.Name == "" {
		return exceptions.NewValidationException("Name cannot be empty")
	}
	if f.FileType == "" {
		return exceptions.NewValidationException("File type cannot be empty")
	}
	if f.IsInvalidType() {
		return exceptions.NewValidationException("File type not supported")
	}
	if f.Size <= 0 {
		return exceptions.NewNotFoundException("File size must be greater than 0")
	}
	if f.Reader == nil {
		return exceptions.NewNotFoundException("File reader cannot be empty")
	}
	if f.Path == "" {
		return exceptions.NewValidationException("Path cannot be empty")
	}
	return nil
}

func (f *File) IsInvalidType() bool {
	return f.FileType != "image/jpeg" && f.FileType != "image/png"
}
