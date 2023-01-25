package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/sheacloud/tfom/internal/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	WorkingDirectory string
	Prefix           string
	StateBucket      string
	StateRegion      string
	ResultsBucket    string
	AccountId        string
	Region           string
	AlertsTopic      string
}

func ConfigFromEnv() Config {
	return Config{
		WorkingDirectory: os.Getenv("TFOM_WORKING_DIRECTORY"),
		Prefix:           os.Getenv("TFOM_PREFIX"),
		StateBucket:      os.Getenv("TFOM_STATE_BUCKET"),
		StateRegion:      os.Getenv("TFOM_STATE_REGION"),
		ResultsBucket:    os.Getenv("TFOM_RESULTS_BUCKET"),
		AccountId:        os.Getenv("TFOM_ACCOUNT_ID"),
		Region:           os.Getenv("TFOM_REGION"),
		AlertsTopic:      os.Getenv("TFOM_ALERTS_TOPIC"),
	}
}

func (c *Config) GetDatabase(context context.Context) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}

func (c *Config) GetApiClient(cfg aws.Config, db *gorm.DB) api.APIClientInterface {
	return api.NewAPIClient(&api.APIClientInput{
		DB:                             db,
		SfnClient:                      sfn.NewFromConfig(cfg),
		SnsClient:                      sns.NewFromConfig(cfg),
		WorkingDirectory:               c.WorkingDirectory,
		ModulePropagationExecutionArn:  fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-module-propagation-execution", c.Region, c.AccountId, c.Prefix),
		ModulePropagationDriftCheckArn: fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-module-propagation-drift-check", c.Region, c.AccountId, c.Prefix),
		TerraformCommandWorkflowArn:    fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-terraform-command", c.Region, c.AccountId, c.Prefix),
		TerraformExecutionArn:          fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-terraform-execution", c.Region, c.AccountId, c.Prefix),
		TerraformDriftCheckArn:         fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-terraform-drift-check", c.Region, c.AccountId, c.Prefix),
		RemoteStateBucket:              c.StateBucket,
		RemoteStateRegion:              c.StateRegion,
		AlertsTopic:                    c.AlertsTopic,
		DataLoaderWaitTime:             time.Millisecond * 16,
	})
}
