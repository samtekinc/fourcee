package api

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/internal/database"
	"github.com/sheacloud/tfom/pkg/models"
)

type APIClientInterface interface {
	GetOrganizationalDimension(ctx context.Context, dimensionId string) (*models.OrganizationalDimension, error)
	GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error)
	PutOrganizationalDimension(ctx context.Context, input *models.NewOrganizationalDimension) (*models.OrganizationalDimension, error)
	DeleteOrganizationalDimension(ctx context.Context, dimensionId string) error

	GetOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string) (*models.OrganizationalUnit, error)
	GetOrganizationalUnits(ctx context.Context, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByDimension(ctx context.Context, dimensionId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByParent(ctx context.Context, dimensionId string, parentOrgUnitId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByHierarchy(ctx context.Context, dimensionId string, hierarchy string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	PutOrganizationalUnit(ctx context.Context, input *models.NewOrganizationalUnit) (*models.OrganizationalUnit, error)
	DeleteOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string) error
	UpdateOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string, update *models.OrganizationalUnitUpdate) (*models.OrganizationalUnit, error)
	UpdateOrganizationalUnitHierarchies(ctx context.Context, orgDimensionId string) error

	GetOrganizationalAccount(ctx context.Context, orgAccountId string) (*models.OrganizationalAccount, error)
	GetOrganizationalAccounts(ctx context.Context, limit int32, cursor string) (*models.OrganizationalAccounts, error)
	PutOrganizationalAccount(ctx context.Context, input *models.NewOrganizationalAccount) (*models.OrganizationalAccount, error)
	DeleteOrganizationalAccount(ctx context.Context, orgAccountId string) error
	UpdateOrganizationalAccount(ctx context.Context, orgAccountId string, update *models.OrganizationalAccountUpdate) (*models.OrganizationalAccount, error)

	GetOrganizationalUnitMembershipsByAccount(ctx context.Context, accountId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	GetOrganizationalUnitMembershipsByOrgUnit(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	GetOrganizationalUnitMembershipsByDimension(ctx context.Context, dimensionId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	PutOrganizationalUnitMembership(ctx context.Context, input *models.NewOrganizationalUnitMembership) (*models.OrganizationalUnitMembership, error)
	DeleteOrganizationalUnitMembership(ctx context.Context, dimensionId string, accountId string) error

	GetModuleGroup(ctx context.Context, moduleGroupId string) (*models.ModuleGroup, error)
	GetModuleGroups(ctx context.Context, limit int32, cursor string) (*models.ModuleGroups, error)
	PutModuleGroup(ctx context.Context, input *models.NewModuleGroup) (*models.ModuleGroup, error)
	DeleteModuleGroup(ctx context.Context, moduleGroupId string) error

	GetModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) (*models.ModuleVersion, error)
	GetModuleVersions(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleVersions, error)
	PutModuleVersion(ctx context.Context, input *models.NewModuleVersion) (*models.ModuleVersion, error)
	DeleteModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) error

	GetModulePropagation(ctx context.Context, modulePropagationId string) (*models.ModulePropagation, error)
	GetModulePropagations(ctx context.Context, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByOrgUnitId(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByOrgDimensionId(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.ModulePropagations, error)
	PutModulePropagation(ctx context.Context, input *models.NewModulePropagation) (*models.ModulePropagation, error)
	DeleteModulePropagation(ctx context.Context, modulePropagationId string) error
	UpdateModulePropagation(ctx context.Context, modulePropagationId string, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error)

	GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	GetModulePropagationExecutionRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	PutModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error)
	UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error)

	GetModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string) (*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error)
	GetModulePropagationDriftCheckRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error)
	PutModulePropagationDriftCheckRequest(ctx context.Context, input *models.NewModulePropagationDriftCheckRequest) (*models.ModulePropagationDriftCheckRequest, error)
	UpdateModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error)

	GetModuleAssignment(ctx context.Context, moduleAssignmentId string) (*models.ModuleAssignment, error)
	GetModuleAssignments(ctx context.Context, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	PutModuleAssignment(ctx context.Context, input *models.NewModuleAssignment) (*models.ModuleAssignment, error)
	UpdateModuleAssignment(ctx context.Context, moduleAssignmentId string, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error)

	GetModulePropagationAssignment(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModulePropagationAssignment, error)
	GetModulePropagationAssignments(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationAssignments, error)
	GetModulePropagationAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error)
	GetModulePropagationAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error)
	PutModulePropagationAssignment(ctx context.Context, input *models.NewModulePropagationAssignment) (*models.ModulePropagationAssignment, *models.ModuleAssignment, error)

	// Execution Methods

	GetTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId string) (*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformExecutionRequests, error)
	GetTerraformExecutionRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformExecutionRequests, error)
	GetTerraformExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformExecutionRequests, error)
	PutTerraformExecutionRequest(ctx context.Context, input *models.NewTerraformExecutionRequest) (*models.TerraformExecutionRequest, error)
	UpdateTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId string, update *models.TerraformExecutionRequestUpdate) (*models.TerraformExecutionRequest, error)

	GetTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestId string) (*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error)
	GetTerraformDriftCheckRequestsByModulePropagationDriftCheckRequestId(ctx context.Context, modulePropagationDriftCheckRequestId string, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error)
	GetTerraformDriftCheckRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error)
	PutTerraformDriftCheckRequest(ctx context.Context, input *models.NewTerraformDriftCheckRequest) (*models.TerraformDriftCheckRequest, error)
	UpdateTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestId string, update *models.TerraformDriftCheckRequestUpdate) (*models.TerraformDriftCheckRequest, error)

	GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string, withOutputs bool) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string, withOutputs bool) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string, withOutputs bool) (*models.PlanExecutionRequests, error)
	PutPlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error)
	UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error)

	GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, withOutputs bool) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string, withOutputs bool) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string, withOutputs bool) (*models.ApplyExecutionRequests, error)
	PutApplyExecutionRequest(ctx context.Context, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error)
	UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, input *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error)

	UploadTerraformPlanInitResults(ctx context.Context, planExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error)
	UploadTerraformPlanResults(ctx context.Context, planExecutionRequestId string, planResults *models.TerraformPlanOutput) (string, error)
	UploadTerraformApplyInitResults(ctx context.Context, applyExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error)
	UploadTerraformApplyResults(ctx context.Context, applyExecutionRequestId string, applyResults *models.TerraformApplyOutput) (string, error)

	DownloadTerraformPlanInitResults(ctx context.Context, initResultsObjectKey string) (*models.TerraformInitOutput, error)
	DownloadTerraformPlanResults(ctx context.Context, planResultsObjectKey string) (*models.TerraformPlanOutput, error)
	DownloadTerraformApplyInitResults(ctx context.Context, initResultsObjectKey string) (*models.TerraformInitOutput, error)
	DownloadTerraformApplyResults(ctx context.Context, applyResultsObjectKey string) (*models.TerraformApplyOutput, error)
}

type APIClient struct {
	dbClient                       database.DatabaseClientInterface
	workingDirectory               string
	sfnClient                      awsclients.StepFunctionsInterface
	modulePropagationExecutionArn  string
	modulePropagationDriftCheckArn string
	terraformCommandWorkflowArn    string
	terraformExecutionArn          string
	terraformDriftCheckArn         string
	remoteStateBucket              string
	remoteStateRegion              string

	applyExecutionRequestsLoader              *dataloader.Loader
	moduleAssignmentsLoader                   *dataloader.Loader
	moduleGroupsLoader                        *dataloader.Loader
	modulePropagationAssignmentsLoader        *dataloader.Loader
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
	DBClient                       database.DatabaseClientInterface
	WorkingDirectory               string
	SfnClient                      awsclients.StepFunctionsInterface
	ModulePropagationExecutionArn  string
	ModulePropagationDriftCheckArn string
	TerraformCommandWorkflowArn    string
	TerraformExecutionArn          string
	TerraformDriftCheckArn         string
	RemoteStateBucket              string
	RemoteStateRegion              string
	DataLoaderWaitTime             time.Duration
}

func NewAPIClient(input *APIClientInput) *APIClient {
	apiClient := &APIClient{
		dbClient:                       input.DBClient,
		workingDirectory:               input.WorkingDirectory,
		sfnClient:                      input.SfnClient,
		modulePropagationExecutionArn:  input.ModulePropagationExecutionArn,
		modulePropagationDriftCheckArn: input.ModulePropagationDriftCheckArn,
		terraformCommandWorkflowArn:    input.TerraformCommandWorkflowArn,
		terraformExecutionArn:          input.TerraformExecutionArn,
		terraformDriftCheckArn:         input.TerraformDriftCheckArn,
		remoteStateBucket:              input.RemoteStateBucket,
		remoteStateRegion:              input.RemoteStateRegion,
	}

	dataLoaderOptions := []dataloader.Option{
		dataloader.WithClearCacheOnBatch(), // don't cache responses long-term, only within a single batch request
		dataloader.WithWait(input.DataLoaderWaitTime),
		dataloader.WithBatchCapacity(100), // limit of BatchGetItems in DynamoDB
	}

	apiClient.applyExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetApplyExecutionRequestsByIds, dataLoaderOptions...)
	apiClient.moduleAssignmentsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleAssignmentsByIds, dataLoaderOptions...)
	apiClient.moduleGroupsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleGroupsByIds, dataLoaderOptions...)
	apiClient.modulePropagationAssignmentsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationAssignmentsByIds, dataLoaderOptions...)
	apiClient.modulePropagationDriftCheckRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationDriftCheckRequestsByIds, dataLoaderOptions...)
	apiClient.modulePropagationExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationExecutionRequestsByIds, dataLoaderOptions...)
	apiClient.modulePropagationsLoader = dataloader.NewBatchedLoader(apiClient.GetModulePropagationsByIds, dataLoaderOptions...)
	apiClient.moduleVersionsLoader = dataloader.NewBatchedLoader(apiClient.GetModuleVersionsByIds, dataLoaderOptions...)
	apiClient.orgAccountsLoader = dataloader.NewBatchedLoader(apiClient.GetOrganizationalAccountsByIds, dataLoaderOptions...)
	apiClient.orgDimensionsLoader = dataloader.NewBatchedLoader(apiClient.GetOrganizationalDimensionsByIds, dataLoaderOptions...)
	apiClient.orgUnitsLoader = dataloader.NewBatchedLoader(apiClient.GetOrganizationalUnitsByIds, dataLoaderOptions...)
	apiClient.planExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetPlanExecutionRequestsByIds, dataLoaderOptions...)
	apiClient.terraformDriftCheckRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetTerraformDriftCheckRequestsByIds, dataLoaderOptions...)
	apiClient.terraformExecutionRequestsLoader = dataloader.NewBatchedLoader(apiClient.GetTerraformExecutionRequestsByIds, dataLoaderOptions...)

	return apiClient
}
