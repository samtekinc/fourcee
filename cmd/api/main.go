package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/api"
	tfomConfig "github.com/sheacloud/tfom/internal/config"
	"github.com/sheacloud/tfom/internal/graph/generated"
	"github.com/sheacloud/tfom/internal/graph/resolver"
	"github.com/sheacloud/tfom/pkg/models"
	"go.uber.org/zap"
)

// Defining the Graphql handler
func graphqlHandler(apiClient api.APIClientInterface, config *tfomConfig.Config) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(apiClient, config)}))

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
	// dbClient := conf.GetDatabaseClient(cfg)
	db, err := conf.GetDatabase(ctx)
	if err != nil {
		panic("unable to get database, " + err.Error())
	}

	err = db.AutoMigrate(&models.OrgAccount{}, &models.OrgDimension{}, &models.OrgUnit{}, &models.Metadata{}, &models.ModuleGroup{}, &models.ModuleVersion{}, &models.ModuleVariable{}, &models.ModulePropagation{}, &models.Argument{}, &models.AwsProviderConfiguration{}, &models.GcpProviderConfiguration{}, &models.ModuleAssignment{},
		&models.ModulePropagationExecutionRequest{}, &models.ModulePropagationDriftCheckRequest{}, &models.TerraformExecutionRequest{}, &models.TerraformDriftCheckRequest{}, &models.PlanExecutionRequest{}, &models.ApplyExecutionRequest{})
	if err != nil {
		panic("unable to migrate database, " + err.Error())
	}

	if err != nil {
		panic("unable to get database, " + err.Error())
	}
	apiClient := conf.GetApiClient(cfg, db)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.POST("/query", graphqlHandler(apiClient, &conf))
	router.GET("/", playgroundHandler())

	router.Run(":8080")
}
