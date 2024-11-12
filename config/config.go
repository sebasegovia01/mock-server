package config

import (
	"fmt"
	"mock-server/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWS struct {
		AccessKeyID     string
		SecretAccessKey string
		Region          string
	}
	S3 struct {
		BucketName string
		FileKey    string
	}
	Server struct {
		Port string
	}
	Env string
}

func LoadConfig() (*Config, error) {
	// Try to load variables from .env
	if err := godotenv.Load(); err != nil {
		logger.InfoLogger.Println(".env file not found, using environment variables")
	} else {
		logger.InfoLogger.Println("Variables loaded from .env file")
	}

	config := &Config{}

	// Load AWS configuration
	config.AWS.AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	config.AWS.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	config.AWS.Region = os.Getenv("AWS_REGION")

	// Load S3 configuration
	config.S3.BucketName = os.Getenv("S3_BUCKET_NAME")
	config.S3.FileKey = os.Getenv("S3_FILE_KEY")

	// Load other configurations
	config.Env = os.Getenv("ENV")

	// Validate required configuration
	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) validate() error {
	if c.AWS.AccessKeyID == "" || c.AWS.SecretAccessKey == "" || c.AWS.Region == "" {
		return fmt.Errorf("incomplete AWS configuration")
	}
	if c.S3.BucketName == "" || c.S3.FileKey == "" {
		return fmt.Errorf("incomplete S3 configuration")
	}
	return nil
}
