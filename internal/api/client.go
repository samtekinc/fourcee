package api

import (
	"context"

	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/internal/database"
	"github.com/sheacloud/tfom/pkg/models"
)

type OrganizationsAPIClientInterface interface {
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

	GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	GetModulePropagationExecutionRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	PutModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error)
	UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error)

	GetModuleAccountAssociation(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModuleAccountAssociation, error)
	GetModuleAccountAssociations(ctx context.Context, limit int32, cursor string) (*models.ModuleAccountAssociations, error)
	GetModuleAccountAssociationsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAccountAssociations, error)
	PutModuleAccountAssociation(ctx context.Context, input *models.NewModuleAccountAssociation) (*models.ModuleAccountAssociation, error)
	UpdateModuleAccountAssociation(ctx context.Context, modulePropagationId string, orgAccountId string, update *models.ModuleAccountAssociationUpdate) (*models.ModuleAccountAssociation, error)

	// Execution Methods

	GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByGroupingKey(ctx context.Context, groupingKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	PutPlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error)
	UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error)

	GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByGroupingKey(ctx context.Context, groupingKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
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

type OrganizationsAPIClient struct {
	dbClient                              database.OrganizationsDatabaseClientInterface
	workingDirectory                      string
	sfnClient                             awsclients.StepFunctionsInterface
	modulePropagationExecutionWorkflowArn string
	terraformExecutionWorkflowArn         string
}

func NewOrganizationsAPIClient(dbClient database.OrganizationsDatabaseClientInterface, workingDirectory string, sfnClient awsclients.StepFunctionsInterface, modulePropagationExecutionWorkflowArn string, terraformExecutionWorkflowArn string) *OrganizationsAPIClient {
	return &OrganizationsAPIClient{
		dbClient:                              dbClient,
		workingDirectory:                      workingDirectory,
		sfnClient:                             sfnClient,
		modulePropagationExecutionWorkflowArn: modulePropagationExecutionWorkflowArn,
		terraformExecutionWorkflowArn:         terraformExecutionWorkflowArn,
	}
}
