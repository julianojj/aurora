package repository

import (
	"github.com/julianojj/aurora/internal/core/domain"
)

type FileRepositoryMemory struct {
	Files []*domain.File
}

func NewFileRepositoryMemory() *FileRepositoryMemory {
	return &FileRepositoryMemory{
		make([]*domain.File, 0),
	}
}

func (frm *FileRepositoryMemory) Save(file *domain.File) error {
	frm.Files = append(frm.Files, file)
	return nil
}

func (frm *FileRepositoryMemory) FindAll() ([]*domain.File, error) {
	return frm.Files, nil
}

func (frm *FileRepositoryMemory) Find(fileID string) (*domain.File, error) {
	for _, file := range frm.Files {
		if file.FileID == fileID {
			return file, nil
		}
	}
	return nil, nil
}

func (frm *FileRepositoryMemory) Delete(fileID string) error {
	var files []*domain.File
	for _, file := range frm.Files {
		if file.FileID != fileID {
			files = append(files, file)
		}
	}
	frm.Files = files
	return nil
}
