package api

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
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
	GetOrgUnitsForDimension(ctx context.Context, dimensionID uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	GetOrgUnitsForParent(ctx context.Context, parentOrgUnitID uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	GetDownstreamOrgUnits(ctx context.Context, orgUnitID uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	GetUpstreamOrgUnits(ctx context.Context, orgUnitID uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
	CreateOrgUnit(ctx context.Context, input *models.NewOrgUnit) (*models.OrgUnit, error)
	DeleteOrgUnit(ctx context.Context, id uint) error
	UpdateOrgUnit(ctx context.Context, id uint, update *models.OrgUnitUpdate) (*models.OrgUnit, error)
	GetOrgAccountsForOrgUnit(ctx context.Context, orgUnitID uint, filters *models.OrgAccountFilters, limit *int, offset *int) ([]*models.OrgAccount, error)
	AddAccountToOrgUnit(ctx context.Context, orgUnitID uint, orgAccountID uint) error
	RemoveAccountFromOrgUnit(ctx context.Context, orgUnitID uint, orgAccountID uint) error

	GetOrgAccount(ctx context.Context, id uint) (*models.OrgAccount, error)
	GetOrgAccountBatched(ctx context.Context, id uint) (*models.OrgAccount, error)
	GetOrgAccounts(ctx context.Context, filters *models.OrgAccountFilters, limit *int, offset *int) ([]*models.OrgAccount, error)
	GetOrgUnitsForOrgAccount(ctx context.Context, orgAccountID uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error)
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
	GetModuleVersionsForModuleGroup(ctx context.Context, moduleGroupID uint, filters *models.ModuleVersionFilters, limit *int, offset *int) ([]*models.ModuleVersion, error)
	CreateModuleVersion(ctx context.Context, input *models.NewModuleVersion) (*models.ModuleVersion, error)
	DeleteModuleVersion(ctx context.Context, id uint) error

	GetModulePropagation(ctx context.Context, id uint) (*models.ModulePropagation, error)
	GetModulePropagationBatched(ctx context.Context, id uint) (*models.ModulePropagation, error)
	GetModulePropagations(ctx context.Context, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForModuleGroup(ctx context.Context, moduleGroupID uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForModuleVersion(ctx context.Context, moduleVersionID uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForOrgUnit(ctx context.Context, orgUnitID uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	GetModulePropagationsForOrgDimension(ctx context.Context, orgDimensionID uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error)
	CreateModulePropagation(ctx context.Context, input *models.NewModulePropagation) (*models.ModulePropagation, error)
	DeleteModulePropagation(ctx context.Context, id uint) error
	UpdateModulePropagation(ctx context.Context, id uint, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error)

	GetModuleAssignment(ctx context.Context, id uint) (*models.ModuleAssignment, error)
	GetModuleAssignmentBatched(ctx context.Context, id uint) (*models.ModuleAssignment, error)
	GetModuleAssignments(ctx context.Context, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForModulePropagation(ctx context.Context, modulePropagationID uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForOrgAccount(ctx context.Context, orgAccountID uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForModuleVersion(ctx context.Context, moduleVersionID uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	GetModuleAssignmentsForModuleGroup(ctx context.Context, moduleGroupID uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error)
	CreateModuleAssignment(ctx context.Context, input *models.NewModuleAssignment) (*models.ModuleAssignment, error)
	UpdateModuleAssignment(ctx context.Context, id uint, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error)
	DeleteModuleAssignment(ctx context.Context, id uint) error

	GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestID uint) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequestBatched(ctx context.Context, modulePropagationExecutionRequestID uint) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequests(ctx context.Context, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequestsForModulePropagation(ctx context.Context, modulePropagationID uint, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error)
	CreateModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error)
	UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestID uint, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error)

	GetModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationDriftCheckRequestID uint) (*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequestBatched(ctx context.Context, modulePropagationDriftCheckRequestID uint) (*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequests(ctx context.Context, filters *models.ModulePropagationDriftCheckRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequestsForModulePropagation(ctx context.Context, modulePropagationID uint, filters *models.ModulePropagationDriftCheckRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationDriftCheckRequest, error)
	CreateModulePropagationDriftCheckRequest(ctx context.Context, input *models.NewModulePropagationDriftCheckRequest) (*models.ModulePropagationDriftCheckRequest, error)
	UpdateModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationDriftCheckRequestID uint, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error)

	GetAwsIamPolicy(ctx context.Context, id uint) (*models.AwsIamPolicy, error)
	GetAwsIamPolicyBatched(ctx context.Context, id uint) (*models.AwsIamPolicy, error)
	GetAwsIamPolicies(ctx context.Context, filters *models.AwsIamPolicyFilters, limit *int, offset *int) ([]*models.AwsIamPolicy, error)
	CreateAwsIamPolicy(ctx context.Context, input *models.NewAwsIamPolicy) (*models.AwsIamPolicy, error)
	UpdateAwsIamPolicy(ctx context.Context, id uint, update *models.AwsIamPolicyUpdate) (*models.AwsIamPolicy, error)
	DeleteAwsIamPolicy(ctx context.Context, id uint) error
	GetAwsIamPoliciesForCloudAccessRole(ctx context.Context, cloudAccessRoleId uint, filters *models.AwsIamPolicyFilters, limit *int, offset *int) ([]*models.AwsIamPolicy, error)

	GetCloudAccessRole(ctx context.Context, id uint) (*models.CloudAccessRole, error)
	GetCloudAccessRoleBatched(ctx context.Context, id uint) (*models.CloudAccessRole, error)
	GetCloudAccessRoles(ctx context.Context, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error)
	CreateCloudAccessRole(ctx context.Context, input *models.NewCloudAccessRole) (*models.CloudAccessRole, error)
	UpdateCloudAccessRole(ctx context.Context, id uint, update *models.CloudAccessRoleUpdate) (*models.CloudAccessRole, error)
	DeleteCloudAccessRole(ctx context.Context, id uint) error
	GetCloudAccessRolesForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error)
	GetInheritedCloudAccessRolesForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error)
	GetCloudAccessRolesForOrgAccount(ctx context.Context, orgAccountId uint, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error)

	// // Execution Methods

	GetTerraformExecutionRequest(ctx context.Context, id uint) (*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequestBatched(ctx context.Context, id uint) (*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequests(ctx context.Context, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequestsForModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestID uint, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error)
	GetTerraformExecutionRequestsForModuleAssignment(ctx context.Context, moduleAssignmentID uint, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error)
	CreateTerraformExecutionRequest(ctx context.Context, input *models.NewTerraformExecutionRequest, triggerWorkflow bool) (*models.TerraformExecutionRequest, error)
	UpdateTerraformExecutionRequest(ctx context.Context, id uint, update *models.TerraformExecutionRequestUpdate) (*models.TerraformExecutionRequest, error)

	GetTerraformDriftCheckRequest(ctx context.Context, id uint) (*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequestBatched(ctx context.Context, id uint) (*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequests(ctx context.Context, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequestsForModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationDriftCheckRequestID uint, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error)
	GetTerraformDriftCheckRequestsForModuleAssignment(ctx context.Context, moduleAssignmentID uint, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error)
	CreateTerraformDriftCheckRequest(ctx context.Context, input *models.NewTerraformDriftCheckRequest, triggerWorkflow bool) (*models.TerraformDriftCheckRequest, error)
	UpdateTerraformDriftCheckRequest(ctx context.Context, id uint, update *models.TerraformDriftCheckRequestUpdate) (*models.TerraformDriftCheckRequest, error)

	GetPlanExecutionRequest(ctx context.Context, id uint) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequestBatched(ctx context.Context, id uint) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestID uint) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequestForTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestID uint) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequests(ctx context.Context, filters *models.PlanExecutionRequestFilters, limit *int, offset *int) ([]*models.PlanExecutionRequest, error)
	CreatePlanExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestID uint, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error)
	CreatePlanExecutionRequestForTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestID uint, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error)
	UpdatePlanExecutionRequest(ctx context.Context, id uint, update *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error)

	GetApplyExecutionRequest(ctx context.Context, id uint) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequestBatched(ctx context.Context, id uint) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestID uint) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequests(ctx context.Context, filters *models.ApplyExecutionRequestFilters, limit *int, offset *int) ([]*models.ApplyExecutionRequest, error)
	CreateApplyExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestID uint, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error)
	UpdateApplyExecutionRequest(ctx context.Context, id uint, update *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error)

	// DownloadResultObject(ctx context.Context, objectKey string) ([]byte, error)
	// GetResultObjectWriter(ctx context.Context, objectKey string, withLiveUploads bool) (io.WriteCloser, error)

	// // Misc methods
	SendAlert(ctx context.Context, subject string, message string) error

	GetStateFileVersions(ctx context.Context, stateBucket string, stateKey string, limit *int) ([]*models.StateVersion, error)
	GetStateFileVersion(ctx context.Context, stateBucket string, stateKey string, versionID string) (*models.StateFile, error)
}
