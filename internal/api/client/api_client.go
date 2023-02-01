package client

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/internal/config"
	"github.com/sheacloud/tfom/pkg/models"
	"go.temporal.io/sdk/client"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type APIClient struct {
	db             *gorm.DB
	temporalClient client.Client

	workingDirectory               string
	snsClient                      awsclients.SNSInterface
	modulePropagationExecutionArn  string
	modulePropagationDriftCheckArn string
	terraformCommandWorkflowArn    string
	terraformExecutionArn          string
	terraformDriftCheckArn         string
	remoteStateBucket              string
	remoteStateRegion              string
	alertsTopic                    string

	applyExecutionRequestsLoader              *dataloader.Loader
	moduleAssignmentsLoader                   *dataloader.Loader
	moduleGroupsLoader                        *dataloader.Loader
	modulePropagationDriftCheckRequestsLoader *dataloader.Loader
	modulePropagationExecutionRequestsLoader  *dataloader.Loader
	modulePropagationsLoader                  *dataloader.Loader
	moduleVersionsLoader                      *dataloader.Loader
	orgAccountsLoader                         *dataloader.Loader
	orgDimensionsLoader                       *dataloader.Loader
	orgUnitsLoader                            *dataloader.Loader
	planExecutionRequestsLoader               *dataloader.Loader
	terraformDriftCheckRequestsLoader         *dataloader.Loader
	terraformExecutionRequestsLoader          *dataloader.Loader
}

type APIClientInput struct {
	DB                             *gorm.DB
	TemporalClient                 client.Client
	WorkingDirectory               string
	SnsClient                      awsclients.SNSInterface
	ModulePropagationExecutionArn  string
	ModulePropagationDriftCheckArn string
	TerraformCommandWorkflowArn    string
	TerraformExecutionArn          string
	TerraformDriftCheckArn         string
	RemoteStateBucket              string
	RemoteStateRegion              string
	DataLoaderWaitTime             time.Duration
	AlertsTopic                    string
}

func NewAPIClient(input *APIClientInput) *APIClient {
	apiClient := &APIClient{
		db:                             input.DB,
		temporalClient:                 input.TemporalClient,
		workingDirectory:               input.WorkingDirectory,
		snsClient:                      input.SnsClient,
		modulePropagationExecutionArn:  input.ModulePropagationExecutionArn,
		modulePropagationDriftCheckArn: input.ModulePropagationDriftCheckArn,
		terraformCommandWorkflowArn:    input.TerraformCommandWorkflowArn,
		terraformExecutionArn:          input.TerraformExecutionArn,
		terraformDriftCheckArn:         input.TerraformDriftCheckArn,
		remoteStateBucket:              input.RemoteStateBucket,
		remoteStateRegion:              input.RemoteStateRegion,
		alertsTopic:                    input.AlertsTopic,
	}

	dataLoaderOptions := []dataloader.Option{
		dataloader.WithClearCacheOnBatch(), // don't cache responses long-term, only within a single batch request
		dataloader.WithWait(input.DataLoaderWaitTime),
		dataloader.WithBatchCapacity(100), // limit of BatchGetItems in DynamoDB
	}

	apiClient.applyExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetApplyExecutionRequestsByIDs, dataLoaderOptions...)
	apiClient.moduleAssignmentsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleAssignmentsByIDs, dataLoaderOptions...)
	apiClient.moduleGroupsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleGroupsByIDs, dataLoaderOptions...)
	apiClient.modulePropagationDriftCheckRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationDriftCheckRequestsByIDs, dataLoaderOptions...)
	apiClient.modulePropagationExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationExecutionRequestsByIDs, dataLoaderOptions...)
	apiClient.modulePropagationsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationsByIDs, dataLoaderOptions...)
	apiClient.moduleVersionsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleVersionsByIDs, dataLoaderOptions...)
	apiClient.orgAccountsLoader = dataloader.NewBatchedLoader(apiClient.GetOrgAccountsByIDs, dataLoaderOptions...)
	apiClient.orgDimensionsLoader = dataloader.NewBatchedLoader(apiClient.GetOrgDimensionsByIDs, dataLoaderOptions...)
	apiClient.orgUnitsLoader = dataloader.NewBatchedLoader(apiClient.GetOrgUnitsByIDs, dataLoaderOptions...)
	apiClient.planExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetPlanExecutionRequestsByIDs, dataLoaderOptions...)
	apiClient.terraformDriftCheckRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetTerraformDriftCheckRequestsByIDs, dataLoaderOptions...)
	apiClient.terraformExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetTerraformExecutionRequestsByIDs, dataLoaderOptions...)

	return apiClient
}

func APIClientFromConfig(conf *config.Config, cfg aws.Config) (*APIClient, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  conf.DBConnectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.OrgAccount{}, &models.OrgDimension{}, &models.OrgUnit{}, &models.ModuleGroup{}, &models.ModuleVersion{}, &models.ModulePropagation{}, &models.ModuleAssignment{},
		&models.ModulePropagationExecutionRequest{}, &models.ModulePropagationDriftCheckRequest{}, &models.TerraformExecutionRequest{}, &models.TerraformDriftCheckRequest{}, &models.PlanExecutionRequest{}, &models.ApplyExecutionRequest{})
	if err != nil {
		panic("unable to migrate database, " + err.Error())
	}

	tc, err := client.Dial(client.Options{})
	if err != nil {
		return nil, err
	}

	return NewAPIClient(&APIClientInput{
		DB:                             db,
		TemporalClient:                 tc,
		SnsClient:                      sns.NewFromConfig(cfg),
		WorkingDirectory:               conf.WorkingDirectory,
		ModulePropagationExecutionArn:  fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-module-propagation-execution", conf.Region, conf.AccountId, conf.Prefix),
		ModulePropagationDriftCheckArn: fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-module-propagation-drift-check", conf.Region, conf.AccountId, conf.Prefix),
		TerraformCommandWorkflowArn:    fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-terraform-command", conf.Region, conf.AccountId, conf.Prefix),
		TerraformExecutionArn:          fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-terraform-execution", conf.Region, conf.AccountId, conf.Prefix),
		TerraformDriftCheckArn:         fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s-terraform-drift-check", conf.Region, conf.AccountId, conf.Prefix),
		RemoteStateBucket:              conf.StateBucket,
		RemoteStateRegion:              conf.StateRegion,
		AlertsTopic:                    conf.AlertsTopic,
		DataLoaderWaitTime:             time.Millisecond * 16,
	}), nil
}
