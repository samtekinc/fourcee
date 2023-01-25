package api

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
)

type APIClientInterface interface {
	GetOrgDimension(ctx context.Context, id uint) (*models.OrgDimension, error)
	GetOrgDimensionBatched(ctx context.Context, id uint) (*models.OrgDimension, error)
	GetOrgDimensions(ctx context.Context, filters *models.OrgDimensionFilters, limit *int, offset *int) ([]*models.OrgDimension, error)
	CreateOrgDimension(ctx context.Context, input *models.NewOrgDimension) (*models.OrgDimension, error)
	DeleteOrgDimension(ctx context.Context, id uint) error

	GetOrgUnit(ctx context.Context, id uint) (*models.OrgUnit, error)
	GetOrgUnitBatched(ctx context.Context, id uint) (*models.OrgUnit, error)
	GetOrgUnits(ctx context.Context, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	GetOrgUnitsForDimension(ctx context.Context, dimensionId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	GetOrgUnitsForParent(ctx context.Context, parentOrgUnitId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	GetDownstreamOrgUnits(ctx context.Context, orgUnitId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	GetUpstreamOrgUnits(ctx context.Context, orgUnitId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	CreateOrgUnit(ctx context.Context, input *models.NewOrgUnit) (*models.OrgUnit, error)
	DeleteOrgUnit(ctx context.Context, id uint) error
	UpdateOrgUnit(ctx context.Context, id uint, update *models.OrgUnitUpdate) (*models.OrgUnit, error)
	GetOrgAccountsForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.OrgAccountFilters, limit *int, offset *int) ([]*models.OrgAccount, error)
	AddAccountToOrgUnit(ctx context.Context, orgUnitId uint, orgAccountId uint) error
	RemoveAccountFromOrgUnit(ctx context.Context, orgUnitId uint, orgAccountId uint) error

	GetOrgAccount(ctx context.Context, id uint) (*models.OrgAccount, error)
	GetOrgAccountBatched(ctx context.Context, id uint) (*models.OrgAccount, error)
	GetOrgAccounts(ctx context.Context, filters *models.OrgAccountFilters, limit *int, offset *int) ([]*models.OrgAccount, error)
	GetOrgUnitsForOrgAccount(ctx context.Context, orgAccountId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	CreateOrgAccount(ctx context.Context, input *models.NewOrgAccount) (*models.OrgAccount, error)
	DeleteOrgAccount(ctx context.Context, id uint) error
	UpdateOrgAccount(ctx context.Context, id uint, update *models.OrgAccountUpdate) (*models.OrgAccount, error)

	GetModuleGroup(ctx context.Context, id uint) (*models.ModuleGroup, error)
	GetModuleGroupBatched(ctx context.Context, id uint) (*models.ModuleGroup, error)
	GetModuleGroups(ctx context.Context, filters *models.ModuleGroupFilters, limit *int, offset *int) ([]*models.ModuleGroup, error)
	CreateModuleGroup(ctx context.Context, input *models.NewModuleGroup) (*models.ModuleGroup, error)
	DeleteModuleGroup(ctx context.Context, id uint) error

	GetModuleVersion(ctx context.Context, id uint) (*models.ModuleVersion, error)
	GetModuleVersionBatched(ctx context.Context, id uint) (*models.ModuleVersion, error)
	GetModuleVersions(ctx context.Context, filters *models.ModuleVersionFilters, limit *int, offset *int) ([]*models.ModuleVersion, error)
	GetModuleVersionsForModuleGroup(ctx context.Context, moduleGroupId uint, filters *models.ModuleVersionFilters, limit *int, offset *int) ([]*models.ModuleVersion, error)
	CreateModuleVersion(ctx context.Context, input *models.NewModuleVersion) (*models.ModuleVersion, error)
	DeleteModuleVersion(ctx context.Context, id uint) error

	GetModulePropagation(ctx context.Context, id uint) (*models.ModulePropagation, error)
	GetModulePropagationBatched(ctx context.Context, id uint) (*models.ModulePropagation, error)
	GetModulePropagations(ctx context.Context, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForModuleGroup(ctx context.Context, moduleGroupId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForModuleVersion(ctx context.Context, moduleVersionId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForOrgDimension(ctx context.Context, orgDimensionId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	CreateModulePropagation(ctx context.Context, input *models.NewModulePropagation) (*models.ModulePropagation, error)
	DeleteModulePropagation(ctx context.Context, id uint) error
	UpdateModulePropagation(ctx context.Context, id uint, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error)

	GetModuleAssignment(ctx context.Context, id uint) (*models.ModuleAssignment, error)
	GetModuleAssignmentBatched(ctx context.Context, id uint) (*models.ModuleAssignment, error)
	GetModuleAssignments(ctx context.Context, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForModulePropagation(ctx context.Context, modulePropagationId uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForOrgAccount(ctx context.Context, orgAccountID uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForModuleVersion(ctx context.Context, moduleVersionId uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForModuleGroup(ctx context.Context, moduleGroupId uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	CreateModuleAssignment(ctx context.Context, input *models.NewModuleAssignment) (*models.ModuleAssignment, error)
	UpdateModuleAssignment(ctx context.Context, id uint, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error)

	GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestId uint) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequestBatched(ctx context.Context, modulePropagationExecutionRequestId uint) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequests(ctx context.Context, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequestsForModulePropagation(ctx context.Context, modulePropagationId uint, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error)
	CreateModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error)
	UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestId uint, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error)

	GetModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationDriftCheckRequestId uint) (*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequestBatched(ctx context.Context, modulePropagationDriftCheckRequestId uint) (*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequests(ctx context.Context, filters *models.ModulePropagationDriftCheckRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequestsForModulePropagation(ctx context.Context, modulePropagationId uint, filters *models.ModulePropagationDriftCheckRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationDriftCheckRequest, error)
	CreateModulePropagationDriftCheckRequest(ctx context.Context, input *models.NewModulePropagationDriftCheckRequest) (*models.ModulePropagationDriftCheckRequest, error)
	UpdateModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationDriftCheckRequestId uint, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error)

	// // Execution Methods

	GetTerraformExecutionRequest(ctx context.Context, id uint) (*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequestBatched(ctx context.Context, id uint) (*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequests(ctx context.Context, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequestsForModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestID uint, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequestsForModuleAssignment(ctx context.Context, moduleAssignmentID uint, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error)
	CreateTerraformExecutionRequest(ctx context.Context, input *models.NewTerraformExecutionRequest) (*models.TerraformExecutionRequest, error)
	UpdateTerraformExecutionRequest(ctx context.Context, id uint, update *models.TerraformExecutionRequestUpdate) (*models.TerraformExecutionRequest, error)

	GetTerraformDriftCheckRequest(ctx context.Context, id uint) (*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequestBatched(ctx context.Context, id uint) (*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequests(ctx context.Context, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequestsForModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationDriftCheckRequestID uint, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequestsForModuleAssignment(ctx context.Context, moduleAssignmentID uint, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error)
	CreateTerraformDriftCheckRequest(ctx context.Context, input *models.NewTerraformDriftCheckRequest) (*models.TerraformDriftCheckRequest, error)
	UpdateTerraformDriftCheckRequest(ctx context.Context, id uint, update *models.TerraformDriftCheckRequestUpdate) (*models.TerraformDriftCheckRequest, error)

	GetPlanExecutionRequest(ctx context.Context, id uint) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequestBatched(ctx context.Context, id uint) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequests(ctx context.Context, filters *models.PlanExecutionRequestFilters, limit *int, offset *int) ([]*models.PlanExecutionRequest, error)
	CreatePlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error)
	UpdatePlanExecutionRequest(ctx context.Context, id uint, update *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error)

	GetApplyExecutionRequest(ctx context.Context, id uint) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequestBatched(ctx context.Context, id uint) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequests(ctx context.Context, filters *models.ApplyExecutionRequestFilters, limit *int, offset *int) ([]*models.ApplyExecutionRequest, error)
	CreateApplyExecutionRequest(ctx context.Context, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error)
	UpdateApplyExecutionRequest(ctx context.Context, id uint, update *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error)

	// DownloadResultObject(ctx context.Context, objectKey string) ([]byte, error)
	// GetResultObjectWriter(ctx context.Context, objectKey string, withLiveUploads bool) (io.WriteCloser, error)

	// // Misc methods
	SendAlert(ctx context.Context, subject string, message string) error
}

type APIClient struct {
	db *gorm.DB

	workingDirectory               string
	sfnClient                      awsclients.StepFunctionsInterface
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
	WorkingDirectory               string
	SfnClient                      awsclients.StepFunctionsInterface
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
		workingDirectory:               input.WorkingDirectory,
		sfnClient:                      input.SfnClient,
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

	apiClient.applyExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetApplyExecutionRequestsByIds, dataLoaderOptions...)
	apiClient.moduleAssignmentsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleAssignmentsByIds, dataLoaderOptions...)
	apiClient.moduleGroupsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleGroupsByIds, dataLoaderOptions...)
	apiClient.modulePropagationDriftCheckRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationDriftCheckRequestsByIds, dataLoaderOptions...)
	apiClient.modulePropagationExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationExecutionRequestsByIds, dataLoaderOptions...)
	apiClient.modulePropagationsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationsByIds, dataLoaderOptions...)
	apiClient.moduleVersionsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleVersionsByIds, dataLoaderOptions...)
	apiClient.orgAccountsLoader = dataloader.NewBatchedLoader(apiClient.GetOrgAccountsByIds, dataLoaderOptions...)
	apiClient.orgDimensionsLoader = dataloader.NewBatchedLoader(apiClient.GetOrgDimensionsByIds, dataLoaderOptions...)
	apiClient.orgUnitsLoader = dataloader.NewBatchedLoader(apiClient.GetOrgUnitsByIds, dataLoaderOptions...)
	apiClient.planExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetPlanExecutionRequestsByIds, dataLoaderOptions...)
	apiClient.terraformDriftCheckRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetTerraformDriftCheckRequestsByIds, dataLoaderOptions...)
	apiClient.terraformExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetTerraformExecutionRequestsByIds, dataLoaderOptions...)

	return apiClient
}
