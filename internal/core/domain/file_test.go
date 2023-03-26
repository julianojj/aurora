package domain

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateFileIfEmptyFileID(t *testing.T) {
	_, err := NewFile(
		"",
		"test",
		"image/jpeg",
		"/aurora/test.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File ID cannot be empty")
}

func TestNotShouldCreateFileIfEmptyName(t *testing.T) {
	_, err := NewFile(
		"1",
		"",
		"image/jpeg",
		"/aurora/test.jpg",
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
		"/aurora/test.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File type cannot be empty")
}

func TestNotShouldCreateFileIfEmptyPath(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		"image/jpeg",
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
		"/aurora/image.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File type not supported")
}

func TestNotShouldCreateFileIfNegativeSize(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		"image/jpeg",
		"/aurora/image.jpg",
		0,
		bytes.NewReader([]byte("")),
	)
	assert.EqualError(t, err, "File size must be greater than 0")
}

func TestNotShouldCreateFileIfEmptyReader(t *testing.T) {
	_, err := NewFile(
		"1",
		"test",
		"image/jpeg",
		"/aurora/image.jpg",
		1,
		nil,
	)
	assert.EqualError(t, err, "File reader cannot be empty")
}

func TestShouldCreateFile(t *testing.T) {
	file, _ := NewFile(
		"1",
		"test",
		"image/jpeg",
		"/aurora/test.jpg",
		1,
		bytes.NewReader([]byte("")),
	)
	assert.Equal(t, "test", file.Name)
}
