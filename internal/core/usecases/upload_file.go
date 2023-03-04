package usecases

import (
	"fmt"
	"io"
	"path"

	"github.com/google/uuid"
	"github.com/julianojj/aurora/internal/core/domain"
)

type UploadFile struct {
	fileRepository domain.FileRepository
}

type UploadFileInput struct {
	Name     string    `json:"name"`
	Mimetype string    `json:"mimetype"`
	Size     int64     `json:"size"`
	Reader   io.Reader `json:"reader"`
}

func NewUploadFile(fileRepository domain.FileRepository) *UploadFile {
	return &UploadFile{
		fileRepository,
	}
}

func (u *UploadFile) Execute(input UploadFileInput) error {
	fileID := uuid.NewString()
	bucketName := "aurora"
	ext := path.Ext(input.Name)
	newName := fmt.Sprintf("%s%s", fileID, ext)
	file, err := domain.NewFile(
		fileID,
		input.Name,
		input.Mimetype,
		fmt.Sprintf("%s/%s", bucketName, newName),
		input.Size,
		input.Reader,
	)
	if err != nil {
		return err
	}
	err = u.fileRepository.Save(file)
	if err != nil {
		return nil
	}
	return nil
}
