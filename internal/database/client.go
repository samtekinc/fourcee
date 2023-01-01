package database

import (
	"context"

	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/pkg/models"
)

type OrganizationsDatabaseClientInterface interface {
	GetOrganizationalDimension(ctx context.Context, orgDimensionId string) (*models.OrganizationalDimension, error)
	GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error)
	PutOrganizationalDimension(ctx context.Context, input *models.OrganizationalDimension) error
	DeleteOrganizationalDimension(ctx context.Context, orgDimensionId string) error

	GetOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string) (*models.OrganizationalUnit, error)
	GetOrganizationalUnits(ctx context.Context, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByDimension(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByParent(ctx context.Context, orgDimensionId string, parentOrgUnitId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByHierarchy(ctx context.Context, orgDimensionId string, hierarchy string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	PutOrganizationalUnit(ctx context.Context, input *models.OrganizationalUnit) error
	DeleteOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string) error
	UpdateOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string, update *OrganizationalUnitUpdate) (*models.OrganizationalUnit, error)

	GetOrganizationalAccount(ctx context.Context, orgAccountId string) (*models.OrganizationalAccount, error)
	GetOrganizationalAccounts(ctx context.Context, limit int32, cursor string) (*models.OrganizationalAccounts, error)
	PutOrganizationalAccount(ctx context.Context, input *models.OrganizationalAccount) error
	DeleteOrganizationalAccount(ctx context.Context, orgAccountId string) error

	GetOrganizationalUnitMembershipsByAccount(ctx context.Context, accountId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	GetOrganizationalUnitMembershipsByOrgUnit(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	GetOrganizationalUnitMembershipsByDimension(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	PutOrganizationalUnitMembership(ctx context.Context, input *models.OrganizationalUnitMembership) error
	DeleteOrganizationalUnitMembership(ctx context.Context, orgDimensionId string, accountId string) error

	GetModuleGroup(ctx context.Context, moduleGroupId string) (*models.ModuleGroup, error)
	GetModuleGroups(ctx context.Context, limit int32, cursor string) (*models.ModuleGroups, error)
	PutModuleGroup(ctx context.Context, input *models.ModuleGroup) error
	DeleteModuleGroup(ctx context.Context, moduleGroupId string) error

	GetModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) (*models.ModuleVersion, error)
	GetModuleVersions(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleVersions, error)
	PutModuleVersion(ctx context.Context, input *models.ModuleVersion) error
	DeleteModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) error

	GetModulePropagation(ctx context.Context, modulePropagationId string) (*models.ModulePropagation, error)
	GetModulePropagations(ctx context.Context, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByOrgUnitId(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.ModulePropagations, error)
	GetModulePropagationsByOrgDimensionId(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.ModulePropagations, error)
	PutModulePropagation(ctx context.Context, input *models.ModulePropagation) error
	DeleteModulePropagation(ctx context.Context, modulePropagationId string) error
	UpdateModulePropagation(ctx context.Context, modulePropagationId string, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error)

	GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	GetModulePropagationExecutionRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	PutModulePropagationExecutionRequest(ctx context.Context, input *models.ModulePropagationExecutionRequest) error
	UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error)

	GetModuleAccountAssociation(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModuleAccountAssociation, error)
	GetModuleAccountAssociations(ctx context.Context, limit int32, cursor string) (*models.ModuleAccountAssociations, error)
	GetModuleAccountAssociationsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAccountAssociations, error)
	GetModuleAccountAssociationsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAccountAssociations, error)
	PutModuleAccountAssociation(ctx context.Context, input *models.ModuleAccountAssociation) error
	UpdateModuleAccountAssociation(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModuleAccountAssociationUpdate) (*models.ModuleAccountAssociation, error)

	// Execution Related Methods

	GetTerraformWorkflowRequest(ctx context.Context, terraformExecutionRequestId string) (*models.TerraformWorkflowRequest, error)
	GetTerraformWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformWorkflowRequests, error)
	GetTerraformWorkflowRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformWorkflowRequests, error)
	GetTerraformWorkflowRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.TerraformWorkflowRequests, error)
	PutTerraformWorkflowRequest(ctx context.Context, input *models.TerraformWorkflowRequest) error
	UpdateTerraformWorkflowRequest(ctx context.Context, terraformExecutionRequestId string, update *models.TerraformWorkflowRequestUpdate) (*models.TerraformWorkflowRequest, error)

	GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	PutPlanExecutionRequest(ctx context.Context, input *models.PlanExecutionRequest) error
	UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error)

	GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	PutApplyExecutionRequest(ctx context.Context, input *models.ApplyExecutionRequest) error
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

type OrganizationsDatabaseClient struct {
	dynamodb awsclients.DynamoDBInterface
	s3       awsclients.S3Interface

	dimensionsTableName                         string
	unitsTableName                              string
	accountsTableName                           string
	membershipsTableName                        string
	groupsTableName                             string
	versionsTableName                           string
	propagationsTableName                       string
	modulePropagationExecutionRequestsTableName string
	moduleAccountAssociationsTableName          string

	terraformWorkflowRequestsTableName string
	planExecutionsTableName            string
	applyExecutionsTableName           string
	resultsBucketName                  string
}

type OrganizationsDatabaseClientInput struct {
	DynamoDB awsclients.DynamoDBInterface
	S3       awsclients.S3Interface

	DimensionsTableName                         string
	UnitsTableName                              string
	AccountsTableName                           string
	MembershipsTableName                        string
	GroupsTableName                             string
	VersionsTableName                           string
	PropagationsTableName                       string
	ModulePropagationExecutionRequestsTableName string
	ModuleAccountAssociationsTableName          string

	TerraformWorkflowRequestsTableName string
	PlanExecutionsTableName            string
	ApplyExecutionsTableName           string
	ResultsBucketName                  string
}

func NewOrganizationsDatabaseClient(input *OrganizationsDatabaseClientInput) *OrganizationsDatabaseClient {
	return &OrganizationsDatabaseClient{
		dynamodb:              input.DynamoDB,
		s3:                    input.S3,
		dimensionsTableName:   input.DimensionsTableName,
		unitsTableName:        input.UnitsTableName,
		accountsTableName:     input.AccountsTableName,
		membershipsTableName:  input.MembershipsTableName,
		groupsTableName:       input.GroupsTableName,
		versionsTableName:     input.VersionsTableName,
		propagationsTableName: input.PropagationsTableName,
		modulePropagationExecutionRequestsTableName: input.ModulePropagationExecutionRequestsTableName,
		moduleAccountAssociationsTableName:          input.ModuleAccountAssociationsTableName,
		terraformWorkflowRequestsTableName:          input.TerraformWorkflowRequestsTableName,
		planExecutionsTableName:                     input.PlanExecutionsTableName,
		applyExecutionsTableName:                    input.ApplyExecutionsTableName,
		resultsBucketName:                           input.ResultsBucketName,
	}
}
