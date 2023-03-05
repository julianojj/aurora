package adapters

import "github.com/julianojj/aurora/internal/core/domain"

type FakeBucket struct {
}

func NewFakeBucket() *FakeBucket {
	return &FakeBucket{}
}

func (f *FakeBucket) CreateBucket() error {
	return nil
}

func (f *FakeBucket) PutObject(file *domain.File) error {
	return nil
}
