package awsclients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSInterface interface {
	PublishInterface
}

type PublishInterface interface {
	Publish(ctx context.Context, params *sns.PublishInput, options ...func(*sns.Options)) (*sns.PublishOutput, error)
}
