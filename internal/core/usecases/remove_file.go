package usecases

import (
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/infra/adapters"
)

type RemoveFile struct {
	fileRepository domain.FileRepository
	bucket         adapters.Bucket
}

func NewRemoveFile(
	fileRepository domain.FileRepository,
	bucket adapters.Bucket,
) *RemoveFile {
	return &RemoveFile{
		fileRepository,
		bucket,
	}
}

func (rf *RemoveFile) Execute(fileID string) error {
	existingFile, err := rf.fileRepository.Find(fileID)
	if err != nil {
		return err
	}
	if existingFile == nil {
		return nil
	}
	err = rf.bucket.DeleteObject(existingFile.Name)
	if err != nil {
		return err
	}
	err = rf.fileRepository.Delete(fileID)
	if err != nil {
		return err
	}
	return nil
}
