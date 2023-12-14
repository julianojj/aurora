package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("Load config", func(t *testing.T) {
		config := LoadConfig()
		assert.Equal(t, "http://localhost:4566", config.AWS_S3_ENDPOINT)
		assert.Equal(t, "us-east-1", config.AWS_REGION)
		assert.Equal(t, "access", config.AWS_ROOT_USER)
		assert.Equal(t, "secretkey", config.AWS_ROOT_PASSWORD)
		assert.Equal(t, "aurora", config.AWS_BUCKET_NAME)
	})
}
