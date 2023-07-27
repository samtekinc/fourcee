package awsclients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Interface interface {
	GetObjectInterface
	ListObjectVersionsInterface
}

type GetObjectInterface interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, options ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type ListObjectVersionsInterface interface {
	ListObjectVersions(ctx context.Context, params *s3.ListObjectVersionsInput, options ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error)
}
