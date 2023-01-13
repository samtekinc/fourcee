package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/database"
	"github.com/sheacloud/tfom/internal/workflow"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)
	sfnClient := sfn.NewFromConfig(cfg)

	dbInput := database.DatabaseClientInput{
		DynamoDB:              dynamodbClient,
		S3:                    s3Client,
		DimensionsTableName:   "tfom-organizational-dimensions",
		UnitsTableName:        "tfom-organizational-units",
		AccountsTableName:     "tfom-organizational-accounts",
		MembershipsTableName:  "tfom-organizational-unit-memberships",
		GroupsTableName:       "tfom-module-groups",
		VersionsTableName:     "tfom-module-versions",
		PropagationsTableName: "tfom-module-propagations",
		ModulePropagationExecutionRequestsTableName:  "tfom-module-propagation-execution-requests",
		ModulePropagationDriftCheckRequestsTableName: "tfom-module-propagation-drift-check-requests",
		ModuleAssignmentsTableName:                   "tfom-module-assignments",
		ModulePropagationAssignmentsTableName:        "tfom-module-propagation-assignments",
		TerraformExecutionRequestsTableName:          "tfom-terraform-execution-requests",
		TerraformDriftCheckRequestsTableName:         "tfom-terraform-drift-check-requests",
		PlanExecutionsTableName:                      "tfom-plan-execution-requests",
		ApplyExecutionsTableName:                     "tfom-apply-execution-requests",
		ResultsBucketName:                            "tfom-execution-results",
	}
	dbClient := database.NewDatabaseClient(&dbInput)
	apiInput := api.APIClientInput{
		DBClient:                       dbClient,
		WorkingDirectory:               "./tmp/",
		SfnClient:                      sfnClient,
		ModulePropagationExecutionArn:  "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-module-propagation-execution",
		ModulePropagationDriftCheckArn: "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-module-propagation-drift-check",
		TerraformCommandWorkflowArn:    "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-terraform-command",
		TerraformExecutionArn:          "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-terraform-execution",
		TerraformDriftCheckArn:         "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-terraform-drift-check",
		RemoteStateBucket:              "tfom-backend-states",
		RemoteStateRegion:              "us-east-1",
		DataLoaderWaitTime:             time.Millisecond * 16,
	}
	apiClient := api.NewAPIClient(&apiInput)

	handler := workflow.NewTaskHandler(apiClient, "tfom-backend-states", "us-east-1")
	lambda.StartWithOptions(handler.RouteTask, lambda.WithContext(context.Background()))
}
