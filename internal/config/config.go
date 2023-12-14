package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AWS_S3_ENDPOINT   string
	AWS_REGION        string
	AWS_ROOT_USER     string
	AWS_ROOT_PASSWORD string
	AWS_BUCKET_NAME   string
}

func LoadConfig() *Config {
	AWS_S3_ENDPOINT := os.Getenv("AWS_S3_ENDPOINT")
	AWS_ROOT_USER := os.Getenv("AWS_ROOT_USER")
	AWS_ROOT_PASSWORD := os.Getenv("AWS_ROOT_PASSWORD")
	AWS_BUCKET_NAME := os.Getenv("AWS_BUCKET_NAME")
	AWS_REGION := os.Getenv("AWS_REGION")

	return &Config{
		AWS_S3_ENDPOINT,
		AWS_REGION,
		AWS_ROOT_USER,
		AWS_ROOT_PASSWORD,
		AWS_BUCKET_NAME,
	}
}
