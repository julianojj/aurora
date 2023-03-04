package usecases

import (
	"bytes"
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestShouldUploadFile(t *testing.T) {
	fileRepository := repository.NewFileRepositoryMemory()
	uploadFile := NewUploadFile(fileRepository)
	file := bytes.NewReader([]byte("test"))
	input := UploadFileInput{
		Name:     "test",
		Size:     123,
		Mimetype: "image/jpeg",
		Reader:   file,
	}
	err := uploadFile.Execute(input)
	assert.NoError(t, err)
	files, _ := fileRepository.FindAll()
	assert.Len(t, files, 1)
	assert.Equal(t, "test", files[0].Name)
}
