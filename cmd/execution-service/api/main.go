package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/services/execution/api"
	"github.com/sheacloud/tfom/internal/services/execution/database"
	"github.com/sheacloud/tfom/internal/services/execution/rest"
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

	dbInput := database.ExecutionDatabaseClientInput{
		DynamoDB:                 dynamodbClient,
		S3:                       s3Client,
		PlanExecutionsTableName:  "tfom-exec-service-plan-execution-requests",
		ApplyExecutionsTableName: "tfom-exec-service-apply-execution-requests",
		ResultsBucketName:        "tfom-exec-service-execution-results",
	}
	execDbClient := database.NewExecutionDatabaseClient(&dbInput)
	execApiClient := api.NewExecutionAPIClient(execDbClient, sfnClient, "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-exec-service-terraform-execution")
	execRouter := rest.NewExecutionRouter(execApiClient)

	router := gin.Default()

	execRouter.RegisterRoutes(&router.RouterGroup)

	router.Run(":8080")
}
