package database

import (
	"context"

	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/pkg/execution/models"
)

type ExecutionDatabaseClientInterface interface {
	GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByGroupingKey(ctx context.Context, groupingKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	PutPlanExecutionRequest(ctx context.Context, input *models.PlanExecutionRequest) error
	UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error)

	GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByGroupingKey(ctx context.Context, groupingKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	PutApplyExecutionRequest(ctx context.Context, input *models.ApplyExecutionRequest) error
	UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, input *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error)

	UploadTerraformPlanInitResults(ctx context.Context, planExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error)
	UploadTerraformPlanResults(ctx context.Context, planExecutionRequestId string, planResults *models.TerraformPlanOutput) (string, error)
	UploadTerraformApplyInitResults(ctx context.Context, applyExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error)
	UploadTerraformApplyResults(ctx context.Context, applyExecutionRequestId string, applyResults *models.TerraformApplyOutput) (string, error)

	DownloadTerraformPlanInitResults(ctx context.Context, initResultsObjectKey string) (*models.TerraformInitOutput, error)
	DownloadTerraformPlanResults(ctx context.Context, planResultsObjectKey string) (*models.TerraformPlanOutput, error)
	DownloadTerraformApplyInitResults(ctx context.Context, initResultsObjectKey string) (*models.TerraformInitOutput, error)
	DownloadTerraformApplyResults(ctx context.Context, applyResultsObjectKey string) (*models.TerraformApplyOutput, error)
}

type ExecutionDatabaseClient struct {
	dynamodb                 awsclients.DynamoDBInterface
	s3                       awsclients.S3Interface
	planExecutionsTableName  string
	applyExecutionsTableName string
	resultsBucketName        string
}

type ExecutionDatabaseClientInput struct {
	DynamoDB                 awsclients.DynamoDBInterface
	S3                       awsclients.S3Interface
	PlanExecutionsTableName  string
	ApplyExecutionsTableName string
	ResultsBucketName        string
}

func NewExecutionDatabaseClient(input *ExecutionDatabaseClientInput) *ExecutionDatabaseClient {
	return &ExecutionDatabaseClient{
		dynamodb:                 input.DynamoDB,
		s3:                       input.S3,
		planExecutionsTableName:  input.PlanExecutionsTableName,
		applyExecutionsTableName: input.ApplyExecutionsTableName,
		resultsBucketName:        input.ResultsBucketName,
	}
}
