package usecases

import (
	"bytes"
	"testing"

	"github.com/julianojj/aurora/internal/infra/adapters"
	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestRemoveFile(t *testing.T) {
	fileRepository := repository.NewFileRepositoryMemory()
	bucket := adapters.NewFakeBucket()
	uploadFile := NewUploadFile(fileRepository, bucket)
	getUploads := NewGetUploads(fileRepository)
	removeFile := NewRemoveFile(fileRepository, bucket)
	file := bytes.NewReader([]byte("test"))
	input := UploadFileInput{
		Name:     "test",
		Size:     123,
		Mimetype: "image/jpeg",
		Reader:   file,
	}
	uploadFile.Execute(input)
	beforeFiles, _ := getUploads.Execute()
	assert.Len(t, beforeFiles, 1)
	removeFile.Execute(beforeFiles[0].ID)
	afterFiles, _ := getUploads.Execute()
	assert.Len(t, afterFiles, 0)
}
