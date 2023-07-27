package main

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samtekinc/fourcee/internal/api"
	"github.com/samtekinc/fourcee/internal/api/client"
	tfomConfig "github.com/samtekinc/fourcee/internal/config"
	"github.com/samtekinc/fourcee/internal/graph/generated"
	"github.com/samtekinc/fourcee/internal/graph/resolver"
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

	apiClient, err := client.APIClientFromConfig(&conf, cfg)
	if err != nil {
		log.Fatalln("unable to create API Client", err)
	}

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.POST("/query", graphqlHandler(apiClient, &conf))
	router.GET("/", playgroundHandler())

	router.Run(":8080")
}
