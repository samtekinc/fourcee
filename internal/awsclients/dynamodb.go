package awsclients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBInterface interface {
	GetItemInterface
	BatchGetItemInterface
	DeleteItemInterface
	PutItemInterface
	UpdateItemInterface
	QueryInterface
	ScanInterface
}

type GetItemInterface interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, options ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

type BatchGetItemInterface interface {
	BatchGetItem(ctx context.Context, params *dynamodb.BatchGetItemInput, options ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)
}

type DeleteItemInterface interface {
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, options ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
}

type PutItemInterface interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, options ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type UpdateItemInterface interface {
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, options ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
}

type QueryInterface interface {
	Query(ctx context.Context, params *dynamodb.QueryInput, options ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

type ScanInterface interface {
	Scan(ctx context.Context, params *dynamodb.ScanInput, options ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
}
