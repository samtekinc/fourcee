package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/sheacloud/tfom/internal/api/client"
	tfomConfig "github.com/sheacloud/tfom/internal/config"
	"github.com/sheacloud/tfom/pkg/models"
	temporalClient "go.temporal.io/sdk/client"
)

func main() {
	c, err := temporalClient.Dial(temporalClient.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal Client", err)
	}
	defer c.Close()

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

	tfExecutionWorkflow, err := apiClient.CreateTerraformExecutionRequest(ctx, &models.NewTerraformExecutionRequest{
		ModuleAssignmentID: 1,
		Destroy:            false,
	})
	if err != nil {
		panic("unable to create terraform execution request, " + err.Error())
	}

	fmt.Println(tfExecutionWorkflow)
}
