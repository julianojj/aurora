package usecases

import (
	"bytes"
	"testing"

	"github.com/julianojj/aurora/internal/infra/adapters"
	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestShouldGetUploads(t *testing.T) {
	fileRepository := repository.NewFileRepositoryMemory()
	bucket := adapters.NewFakeBucket()
	logger, _ := zap.NewProduction()
	uploadFile := NewUploadFile(fileRepository, bucket, logger)
	getUploads := NewGetUploads(fileRepository)
	file := bytes.NewReader([]byte("test"))
	input := UploadFileInput{
		Name:     "test",
		Size:     123,
		Mimetype: "image/jpeg",
		Reader:   file,
	}
	err := uploadFile.Execute(input)
	assert.NoError(t, err)
	uploads, err := getUploads.Execute()
	assert.NoError(t, err)
	assert.Len(t, uploads, 1)
}
