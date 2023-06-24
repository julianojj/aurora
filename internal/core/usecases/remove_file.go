package usecases

import "github.com/julianojj/aurora/internal/core/domain"

type RemoveFile struct {
	fileRepository domain.FileRepository
}

func NewRemoveFile(fileRepository domain.FileRepository) *RemoveFile {
	return &RemoveFile{
		fileRepository,
	}
}

func (rf *RemoveFile) Execute(fileID string) error {
	err := rf.fileRepository.Delete(fileID)
	if err != nil {
		return err
	}
	return nil
}
