package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/database"
	"github.com/sheacloud/tfom/internal/graph/generated"
	"github.com/sheacloud/tfom/internal/graph/resolver"
)

// Defining the Graphql handler
func graphqlHandler(apiClient api.OrganizationsAPIClientInterface) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(apiClient)}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

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
		TerraformWorkflowRequestsTableName:          "tfom-terraform-workflow-requests",
		PlanExecutionsTableName:                     "tfom-plan-execution-requests",
		ApplyExecutionsTableName:                    "tfom-apply-execution-requests",
		ResultsBucketName:                           "tfom-execution-results",
	}
	dbClient := database.NewOrganizationsDatabaseClient(&dbInput)
	apiClient := api.NewOrganizationsAPIClient(dbClient, "./tmp/", sfnClient, "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-execute-module-propagation", "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-terraform-execution")

	// orgRouter := rest.NewOrganizationsRouter(apiClient)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.POST("/query", graphqlHandler(apiClient))
	router.GET("/", playgroundHandler())
	// orgRouter.RegisterRoutes(&router.RouterGroup)

	router.Run(":8080")
}
