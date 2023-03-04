package domain

import "github.com/julianojj/aurora/internal/core/exceptions"

type File struct {
	FileID   string
	Name     string
	FileType string
	Path     string
}

func NewFile(fileID, name, fileType, path string) (*File, error) {
	file := &File{
		FileID:   fileID,
		Name:     name,
		FileType: fileType,
		Path:     path,
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
	if f.Path == "" {
		return exceptions.NewValidationException("Path cannot be empty")
	}
	return nil
}

func (f *File) IsInvalidType() bool {
	return f.FileType != "image/jpeg" && f.FileType != "image/png"
}
