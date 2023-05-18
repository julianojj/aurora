package domain

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ImagePath = "image/aurora.jpg"
const ImageType = "image/jpeg"

func TestNotShouldCreateFileIfEmptyFileID(t *testing.T) {
	_, err := NewFile(
		"",
		"test",
		ImageType,
		ImagePath,
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File ID cannot be empty")
}

func TestNotShouldCreateFileIfEmptyName(t *testing.T) {
	_, err := NewFile(
		"1",
		"",
		ImageType,
		ImagePath,
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "Name cannot be empty")
}

func TestNotShouldCreateFileIfEmptyFileType(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		"",
		ImagePath,
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File type cannot be empty")
}

func TestNotShouldCreateFileIfEmptyPath(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		ImageType,
		"",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "Path cannot be empty")
}

func TestNotShouldCreateFileIfInvalidFileType(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		"image/test",
		ImagePath,
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File type not supported")
}

func TestNotShouldCreateFileIfNegativeSize(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		ImageType,
		ImagePath,
		0,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File size must be greater than 0")
}

func TestNotShouldCreateFileIfEmptyReader(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		ImageType,
		ImagePath,
		1,
		nil,
	)
	assert.EqualError(t, err, "File reader cannot be empty")
}

func TestShouldCreateFile(t *testing.T) {
	file, _ := NewFile(
		"1",
		"test",
		ImageType,
		ImagePath,
		1,
		bytes.NewReader([]byte("")),
	)
	assert.Equal(t, "test", file.Name)
}
