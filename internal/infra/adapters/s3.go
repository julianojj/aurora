package adapters

import (
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/julianojj/aurora/internal/config"
	"github.com/julianojj/aurora/internal/core/domain"
)

type S3 struct {
	client *s3.S3
	config *config.Config
}

func NewS3(
	config *config.Config,
) *S3 {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(
			*aws.String(config.AWS_ROOT_USER),
			*aws.String(config.AWS_ROOT_PASSWORD),
			"",
		),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(config.AWS_S3_ENDPOINT),
	})

	if err != nil {
		panic(err)
	}
	client := s3.New(sess)
	return &S3{
		client,
		config,
	}
}

func (s *S3) CreateBucket() error {
	_, err := s.client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(s.config.AWS_BUCKET_NAME),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *S3) PutObject(file *domain.File) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.config.AWS_BUCKET_NAME),
		Key:    aws.String(file.Name),
		Body:   aws.ReadSeekCloser(file.Reader),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *S3) GetObject(fileID string) ([]byte, error) {
	result, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.config.AWS_BUCKET_NAME),
		Key:    aws.String(fileID),
	})
	if err == nil {
		b, err := io.ReadAll(result.Body)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	aerr := err.(awserr.Error)
	switch aerr.Code() {
	case s3.ErrCodeNoSuchKey:
		return nil, errors.New("file not found")
	default:
		return nil, err
	}
}

func (s *S3) DeleteObject(fileID string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.config.AWS_BUCKET_NAME),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return err
	}
	return nil
}
