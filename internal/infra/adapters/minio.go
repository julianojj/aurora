package adapters

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	client     *minio.Client
	ctx        context.Context
	bucketName string
}

func NewMinio(bucketName string) *Minio {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessID := os.Getenv("MINIO_ROOT_USER")
	accessKey := os.Getenv("MINIO_ROOT_PASSWORD")
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessID, accessKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	return &Minio{
		client:     client,
		ctx:        context.Background(),
		bucketName: bucketName,
	}
}

func (m *Minio) CreateBucket() error {
	existingBucket, err := m.client.BucketExists(m.ctx, m.bucketName)
	if err != nil {
		return err
	}
	if existingBucket {
		return nil
	}
	err = m.client.MakeBucket(m.ctx, m.bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Created bucket %s\n", m.bucketName)
	return nil
}

func (m *Minio) PutObject(file *domain.File) error {
	_, err := m.client.PutObject(m.ctx, m.bucketName, file.Name, file.Reader, file.Size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Minio) GetObject(fileID string) ([]byte, error) {
	object, err := m.client.GetObject(m.ctx, m.bucketName, fileID, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	stat, err := object.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stat.Size)
	n, err := object.Read(data)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return data[:n], nil
}
