package adapters

import "github.com/julianojj/aurora/internal/core/domain"

type FakeBucket struct {
	files []*domain.File
}

func NewFakeBucket() *FakeBucket {
	return &FakeBucket{
		files: make([]*domain.File, 0),
	}
}

func (f *FakeBucket) CreateBucket() error {
	return nil
}

func (f *FakeBucket) PutObject(file *domain.File) error {
	f.files = append(f.files, file)
	return nil
}
