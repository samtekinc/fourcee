package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/services/modules/api"
	"github.com/sheacloud/tfom/internal/services/modules/database"
	"github.com/sheacloud/tfom/internal/services/modules/rest"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)

	orgDbClient := database.NewModulesDatabaseClient(dynamodbClient, "tfom-modules-service-module-groups", "tfom-modules-service-module-versions")
	orgApiClient := api.NewModulesAPIClient(orgDbClient)
	orgRouter := rest.NewModulesRouter(orgApiClient)

	router := gin.Default()

	orgRouter.RegisterRoutes(router.Group("/modules"))

	router.Run(":8080")
}
