package adapters

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/julianojj/aurora/internal/core/domain"
)

type S3 struct {
	client     *s3.S3
	bucketName string
}

func NewS3(bucketName string) *S3 {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		},
		Profile: *aws.String(os.Getenv("PROFILE")),
	})
	if err != nil {
		panic(err)
	}
	client := s3.New(sess)
	return &S3{
		client,
		bucketName,
	}
}

func (s *S3) CreateBucket() error {
	_, err := s.client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(s.bucketName),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *S3) PutObject(file *domain.File) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(file.FileID),
		Body:   aws.ReadSeekCloser(file.Reader),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *S3) DeleteObject(fileID string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return err
	}
	return nil
}
