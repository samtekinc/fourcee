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

	dbInput := database.OrganizationsDatabaseClientInput{
		DynamoDB:              dynamodbClient,
		DimensionsTableName:   "tfom-org-service-organizational-dimensions",
		UnitsTableName:        "tfom-org-service-organizational-units",
		AccountsTableName:     "tfom-org-service-organizational-accounts",
		MembershipsTableName:  "tfom-org-service-organizational-unit-memberships",
		GroupsTableName:       "tfom-org-service-module-groups",
		VersionsTableName:     "tfom-org-service-module-versions",
		PropagationsTableName: "tfom-org-service-module-propagations",
	}
	orgDbClient := database.NewOrganizationsDatabaseClient(&dbInput)
	orgApiClient := api.NewOrganizationsAPIClient(orgDbClient, "./tmp/")
	orgRouter := rest.NewOrganizationsRouter(orgApiClient)

	router := gin.Default()

	orgRouter.RegisterRoutes(&router.RouterGroup)

	router.Run(":8080")
}
