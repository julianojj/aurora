package usecases

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/infra/adapters"
)

type UploadFile struct {
	fileRepository domain.FileRepository
	bucket         adapters.Bucket
}

type UploadFileInput struct {
	Name     string    `json:"name"`
	Mimetype string    `json:"mimetype"`
	Size     int64     `json:"size"`
	Reader   io.Reader `json:"reader"`
}

func NewUploadFile(fileRepository domain.FileRepository, bucket adapters.Bucket) *UploadFile {
	return &UploadFile{
		fileRepository,
		bucket,
	}
}

func (u *UploadFile) Execute(input UploadFileInput) error {
	fileID := uuid.NewString()
	bucketName := os.Getenv("MINIO_BUCKET_NAME")
	ext := path.Ext(input.Name)
	newName := fmt.Sprintf("%s%s", fileID, ext)
	file, err := domain.NewFile(
		fileID,
		input.Name,
		input.Mimetype,
		fmt.Sprintf("/%s/%s", bucketName, newName),
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
	err = u.bucket.PutObject(file)
	if err != nil {
		return err
	}
	return nil
}
