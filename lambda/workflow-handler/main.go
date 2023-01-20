package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	tfomConfig "github.com/sheacloud/tfom/internal/config"
	"github.com/sheacloud/tfom/internal/workflow"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("unable to create logger, " + err.Error())
	}
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	conf := tfomConfig.ConfigFromEnv()
	dbClient := conf.GetDatabaseClient(cfg)
	apiClient := conf.GetApiClient(cfg, dbClient)

	handler := workflow.NewTaskHandler(apiClient, &conf)
	lambda.StartWithOptions(handler.RouteTask, lambda.WithContext(context.Background()))
}
