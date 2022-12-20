package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/services/organizations/api"
	"github.com/sheacloud/tfom/internal/services/organizations/database"
	"github.com/sheacloud/tfom/internal/services/organizations/rest"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)

	orgDbClient := database.NewOrganizationsDatabaseClient(dynamodbClient, "tfom-org-service-organizational-dimensions", "tfom-org-service-organizational-units", "tfom-org-service-organizational-accounts", "tfom-org-service-organizational-unit-memberships")
	orgApiClient := api.NewOrganizationsAPIClient(orgDbClient)
	orgRouter := rest.NewOrganizationsRouter(orgApiClient)

	router := gin.Default()

	orgRouter.RegisterRoutes(router.Group("/organizations"))

	router.Run(":8080")
}
