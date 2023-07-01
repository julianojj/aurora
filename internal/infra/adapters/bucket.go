package adapters

import (
	"github.com/julianojj/aurora/internal/core/domain"
)

type Bucket interface {
	CreateBucket() error
	PutObject(file *domain.File) error
	GetObject(fileID string) ([]byte, error)
	DeleteObject(fileID string) error
}
