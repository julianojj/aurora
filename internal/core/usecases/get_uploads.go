package usecases

import "github.com/julianojj/aurora/internal/core/domain"

type GetUploads struct {
	fileRepository domain.FileRepository
}

type GetUploadsOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func NewGetUploads(
	fileRepository domain.FileRepository,
) *GetUploads {
	return &GetUploads{
		fileRepository,
	}
}

func (gu *GetUploads) Execute() ([]*GetUploadsOutput, error) {
	uploads, err := gu.fileRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var output []*GetUploadsOutput
	for _, upload := range uploads {
		output = append(output, &GetUploadsOutput{
			ID:   upload.FileID,
			Name: upload.Name,
			Size: upload.Size,
		})
	}
	return output, nil
}
