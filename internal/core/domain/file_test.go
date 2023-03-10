package domain

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateFileIfEmptyFileID(t *testing.T) {
	file, err := NewFile(
		"",
		"test",
		"image/jpeg",
		"/aurora/test.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File ID cannot be empty")
	assert.Nil(t, file)
}

func TestNotShouldCreateFileIfEmptyName(t *testing.T) {
	file, err := NewFile(
		"1",
		"",
		"image/jpeg",
		"/aurora/test.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "Name cannot be empty")
	assert.Nil(t, file)
}

func TestNotShouldCreateFileIfEmptyFileType(t *testing.T) {
	file, err := NewFile(
		"1",
		"test",
		"",
		"/aurora/test.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File type cannot be empty")
	assert.Nil(t, file)
}

func TestNotShouldCreateFileIfEmptyPath(t *testing.T) {
	file, err := NewFile(
		"1",
		"test",
		"image/jpeg",
		"",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "Path cannot be empty")
	assert.Nil(t, file)
}

func TestNotShouldCreateFileIfInvalidFileType(t *testing.T) {
	file, err := NewFile(
		"1",
		"test",
		"image/test",
		"/aurora/image.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File type not supported")
	assert.Nil(t, file)
}

func TestNotShouldCreateFileIfNegativeSize(t *testing.T) {
	file, err := NewFile(
		"1",
		"test",
		"image/jpeg",
		"/aurora/image.jpg",
		0,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File size must be greater than 0")
	assert.Nil(t, file)
}

func TestNotShouldCreateFileIfEmptyReader(t *testing.T) {
	file, err := NewFile(
		"1",
		"test",
		"image/jpeg",
		"/aurora/image.jpg",
		1,
		nil,
	)
	assert.EqualError(t, err, "File reader cannot be empty")
	assert.Nil(t, file)
}

func TestShouldCreateFile(t *testing.T) {
	file, err := NewFile(
		"1",
		"test",
		"image/jpeg",
		"/aurora/test.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.NoError(t, err)
	assert.Equal(t, "test", file.Name)
}
