package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/samtekinc/fourcee/internal/api/client"
	tfomConfig "github.com/samtekinc/fourcee/internal/config"
	"github.com/samtekinc/fourcee/internal/temporal/activities"
	"github.com/samtekinc/fourcee/internal/temporal/constants"
	"github.com/samtekinc/fourcee/internal/temporal/workflows"
	temporalClient "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	w := worker.New(c, constants.TFOMTaskQueue, worker.Options{
		MaxConcurrentActivityExecutionSize: 1,
	})
	workflows.RegisterWorkflows(w)
	a := activities.NewActivities(apiClient, &conf)
	w.RegisterActivity(a)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
