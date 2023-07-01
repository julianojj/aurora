package usecases

import (
	"bytes"
	"testing"

	"github.com/julianojj/aurora/internal/infra/adapters"
	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetAssets(t *testing.T) {
	fileRepository := repository.NewFileRepositoryMemory()
	bucket := adapters.NewFakeBucket()
	logger, _ := zap.NewProduction()
	uploadFile := NewUploadFile(fileRepository, bucket)
	getAsset := NewGetAsset(bucket, logger)
	file := bytes.NewReader([]byte("test"))
	input := UploadFileInput{
		Name:     "test",
		Size:     123,
		Mimetype: "image/jpeg",
		Reader:   file,
	}
	uploadFile.Execute(input)
	files, _ := fileRepository.FindAll()
	asset, _ := getAsset.Execute(files[0].FileID)
	assert.Len(t, asset, int(file.Size()))
}
