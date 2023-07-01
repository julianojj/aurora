package usecases

import (
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/infra/adapters"
	"go.uber.org/zap"
)

type RemoveFile struct {
	fileRepository domain.FileRepository
	bucket         adapters.Bucket
	logger         *zap.Logger
}

func NewRemoveFile(
	fileRepository domain.FileRepository,
	bucket adapters.Bucket,
	logger *zap.Logger,
) *RemoveFile {
	return &RemoveFile{
		fileRepository,
		bucket,
		logger,
	}
}

func (rf *RemoveFile) Execute(fileID string) error {
	existingFile, err := rf.fileRepository.Find(fileID)
	if err != nil {
		rf.logger.Info(err.Error())
		return err
	}
	if existingFile == nil {
		return nil
	}
	err = rf.bucket.DeleteObject(existingFile.Name)
	if err != nil {
		rf.logger.Info(err.Error())
		return err
	}
	err = rf.fileRepository.Delete(fileID)
	if err != nil {
		rf.logger.Info(err.Error())
		return err
	}
	rf.logger.Info("Removed file")
	return nil
}
