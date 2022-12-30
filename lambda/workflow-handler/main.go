package main

import (
	"context"

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

	dbInput := database.OrganizationsDatabaseClientInput{
		DynamoDB:              dynamodbClient,
		S3:                    s3Client,
		DimensionsTableName:   "tfom-organizational-dimensions",
		UnitsTableName:        "tfom-organizational-units",
		AccountsTableName:     "tfom-organizational-accounts",
		MembershipsTableName:  "tfom-organizational-unit-memberships",
		GroupsTableName:       "tfom-module-groups",
		VersionsTableName:     "tfom-module-versions",
		PropagationsTableName: "tfom-module-propagations",
		ModulePropagationExecutionRequestsTableName: "tfom-module-propagation-execution-requests",
		ModuleAccountAssociationsTableName:          "tfom-module-account-associations",
		PlanExecutionsTableName:                     "tfom-plan-execution-requests",
		ApplyExecutionsTableName:                    "tfom-apply-execution-requests",
		ResultsBucketName:                           "tfom-execution-results",
	}
	dbClient := database.NewOrganizationsDatabaseClient(&dbInput)
	apiClient := api.NewOrganizationsAPIClient(dbClient, "./tmp/", sfnClient, "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-execute-module-propagation", "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-terraform-execution")

	handler := workflow.NewTaskHandler(apiClient, "tfom-backend-states", "us-east-1")
	lambda.StartWithOptions(handler.RouteTask, lambda.WithContext(context.Background()))
}
