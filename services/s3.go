package services

import (
	"context"
	"encoding/json"
	"mock-server/config"
	"mock-server/models"
	"mock-server/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3Data(cfg *config.Config) ([]models.User, error) {
	logger.InfoLogger.Println("Starting AWS configuration loading")
	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AWS.AccessKeyID,
			cfg.AWS.SecretAccessKey,
			"",
		)),
		awsconfig.WithRegion(cfg.AWS.Region),
	)
	if err != nil {
		logger.ErrorLogger.Printf("Error loading AWS configuration: %v", err)
		return nil, err
	}

	logger.InfoLogger.Printf("Connecting to S3 bucket: %s", cfg.S3.BucketName)
	client := s3.NewFromConfig(awsCfg)
	input := &s3.GetObjectInput{
		Bucket: aws.String(cfg.S3.BucketName),
		Key:    aws.String(cfg.S3.FileKey),
	}

	result, err := client.GetObject(context.TODO(), input)
	if err != nil {
		logger.ErrorLogger.Printf("Error getting object from S3: %v", err)
		return nil, err
	}
	defer result.Body.Close()

	logger.InfoLogger.Println("Decoding JSON data")
	var response struct {
		Users []models.User `json:"users"`
	}

	if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
		logger.ErrorLogger.Printf("Error decoding JSON: %v", err)
		return nil, err
	}

	logger.InfoLogger.Printf("Data loaded successfully. Users found: %d", len(response.Users))
	return response.Users, nil
}
