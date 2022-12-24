package awsclients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

type StepFunctionsInterface interface {
	StartExecutionInterface
	DescribeExecutionInterface
}

type StartExecutionInterface interface {
	StartExecution(ctx context.Context, params *sfn.StartExecutionInput, options ...func(*sfn.Options)) (*sfn.StartExecutionOutput, error)
}

type DescribeExecutionInterface interface {
	DescribeExecution(ctx context.Context, params *sfn.DescribeExecutionInput, options ...func(*sfn.Options)) (*sfn.DescribeExecutionOutput, error)
}
