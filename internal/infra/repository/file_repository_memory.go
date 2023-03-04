package repository

import "github.com/julianojj/aurora/internal/core/domain"

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
