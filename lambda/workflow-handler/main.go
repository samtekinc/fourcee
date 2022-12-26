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
		DimensionsTableName:   "tfom-org-service-organizational-dimensions",
		UnitsTableName:        "tfom-org-service-organizational-units",
		AccountsTableName:     "tfom-org-service-organizational-accounts",
		MembershipsTableName:  "tfom-org-service-organizational-unit-memberships",
		GroupsTableName:       "tfom-org-service-module-groups",
		VersionsTableName:     "tfom-org-service-module-versions",
		PropagationsTableName: "tfom-org-service-module-propagations",
		ModulePropagationExecutionRequestsTableName: "tfom-org-service-module-propagation-execution-requests",
		ModuleAccountAssociationsTableName:          "tfom-org-service-module-account-associations",
		PlanExecutionsTableName:                     "tfom-exec-service-plan-execution-requests",
		ApplyExecutionsTableName:                    "tfom-exec-service-apply-execution-requests",
		ResultsBucketName:                           "tfom-exec-service-execution-results",
	}
	dbClient := database.NewOrganizationsDatabaseClient(&dbInput)
	apiClient := api.NewOrganizationsAPIClient(dbClient, "./tmp/", sfnClient, "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-org-service-execute-module-propagation", "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-exec-service-terraform-execution")

	handler := workflow.NewTaskHandler(apiClient)
	lambda.StartWithOptions(handler.RouteTask, lambda.WithContext(context.Background()))
}
