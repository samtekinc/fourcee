package awsclients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Interface interface {
	GetObjectInterface
	PutObjectInterface
}

type GetObjectInterface interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, options ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type PutObjectInterface interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, options ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}
