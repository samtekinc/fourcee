package config

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/database"
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

func (c *Config) GetDatabaseClient(cfg aws.Config) database.DatabaseClientInterface {
	return database.NewDatabaseClient(&database.DatabaseClientInput{
		DynamoDB:              dynamodb.NewFromConfig(cfg),
		S3:                    s3.NewFromConfig(cfg),
		DimensionsTableName:   fmt.Sprintf("%s-organizational-dimensions", c.Prefix),
		UnitsTableName:        fmt.Sprintf("%s-organizational-units", c.Prefix),
		AccountsTableName:     fmt.Sprintf("%s-organizational-accounts", c.Prefix),
		MembershipsTableName:  fmt.Sprintf("%s-organizational-unit-memberships", c.Prefix),
		GroupsTableName:       fmt.Sprintf("%s-module-groups", c.Prefix),
		VersionsTableName:     fmt.Sprintf("%s-module-versions", c.Prefix),
		PropagationsTableName: fmt.Sprintf("%s-module-propagations", c.Prefix),
		ModulePropagationExecutionRequestsTableName:  fmt.Sprintf("%s-module-propagation-execution-requests", c.Prefix),
		ModulePropagationDriftCheckRequestsTableName: fmt.Sprintf("%s-module-propagation-drift-check-requests", c.Prefix),
		ModuleAssignmentsTableName:                   fmt.Sprintf("%s-module-assignments", c.Prefix),
		ModulePropagationAssignmentsTableName:        fmt.Sprintf("%s-module-propagation-assignments", c.Prefix),
		TerraformExecutionRequestsTableName:          fmt.Sprintf("%s-terraform-execution-requests", c.Prefix),
		TerraformDriftCheckRequestsTableName:         fmt.Sprintf("%s-terraform-drift-check-requests", c.Prefix),
		PlanExecutionsTableName:                      fmt.Sprintf("%s-plan-execution-requests", c.Prefix),
		ApplyExecutionsTableName:                     fmt.Sprintf("%s-apply-execution-requests", c.Prefix),
		ResultsBucketName:                            fmt.Sprintf("%s-execution-results", c.Prefix),
	})
}

func (c *Config) GetApiClient(cfg aws.Config, dbClient database.DatabaseClientInterface) api.APIClientInterface {
	return api.NewAPIClient(&api.APIClientInput{
		DBClient:                       dbClient,
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
