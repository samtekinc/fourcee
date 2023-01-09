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
	UpdateOrganizationalAccount(ctx context.Context, orgAccountId string, update *models.OrganizationalAccountUpdate) (*models.OrganizationalAccount, error)

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

	GetModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string) (*models.ModulePropagationDriftCheckRequest, error)
	GetModulePropagationDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error)
	GetModulePropagationDriftCheckRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error)
	PutModulePropagationDriftCheckRequest(ctx context.Context, input *models.ModulePropagationDriftCheckRequest) error
	UpdateModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error)

	GetModuleAssignment(ctx context.Context, moduleAssignmentId string) (*models.ModuleAssignment, error)
	GetModuleAssignments(ctx context.Context, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	GetModuleAssignmentsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleAssignments, error)
	PutModuleAssignment(ctx context.Context, input *models.ModuleAssignment) error
	UpdateModuleAssignment(ctx context.Context, moduleAssignmentId string, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error)

	GetModulePropagationAssignment(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModulePropagationAssignment, error)
	GetModulePropagationAssignments(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationAssignments, error)
	GetModulePropagationAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error)
	GetModulePropagationAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error)
	PutModulePropagationAssignment(ctx context.Context, input *models.ModuleAssignment) (*models.ModulePropagationAssignment, *models.ModuleAssignment, error) // yes, this should be a ModuleAssignment - we perform a db transaction to create both a module assignment and a module propagation assignment

	// Execution Related Methods

	GetTerraformExecutionWorkflowRequest(ctx context.Context, terraformExecutionRequestId string) (*models.TerraformExecutionWorkflowRequest, error)
	GetTerraformExecutionWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error)
	GetTerraformExecutionWorkflowRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error)
	GetTerraformExecutionWorkflowRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error)
	PutTerraformExecutionWorkflowRequest(ctx context.Context, input *models.TerraformExecutionWorkflowRequest) error
	UpdateTerraformExecutionWorkflowRequest(ctx context.Context, terraformExecutionRequestId string, update *models.TerraformExecutionWorkflowRequestUpdate) (*models.TerraformExecutionWorkflowRequest, error)

	GetTerraformDriftCheckWorkflowRequest(ctx context.Context, terraformDriftCheckRequestId string) (*models.TerraformDriftCheckWorkflowRequest, error)
	GetTerraformDriftCheckWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error)
	GetTerraformDriftCheckWorkflowRequestsByModulePropagationDriftCheckRequestId(ctx context.Context, modulePropagationDriftCheckRequestId string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error)
	GetTerraformDriftCheckWorkflowRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error)
	PutTerraformDriftCheckWorkflowRequest(ctx context.Context, input *models.TerraformDriftCheckWorkflowRequest) error
	UpdateTerraformDriftCheckWorkflowRequest(ctx context.Context, terraformDriftCheckRequestId string, update *models.TerraformDriftCheckWorkflowRequestUpdate) (*models.TerraformDriftCheckWorkflowRequest, error)

	GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error)
	GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	GetPlanExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.PlanExecutionRequests, error)
	PutPlanExecutionRequest(ctx context.Context, input *models.PlanExecutionRequest) error
	UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error)

	GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error)
	GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
	GetApplyExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.ApplyExecutionRequests, error)
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

	dimensionsTableName                          string
	unitsTableName                               string
	accountsTableName                            string
	membershipsTableName                         string
	groupsTableName                              string
	versionsTableName                            string
	propagationsTableName                        string
	modulePropagationExecutionRequestsTableName  string
	modulePropagationDriftCheckRequestsTableName string
	moduleAssignmentsTableName                   string
	modulePropagationAssignmentsTableName        string

	terraformExecutionWorkflowRequestsTableName  string
	terraformDriftCheckWorkflowRequestsTableName string
	planExecutionsTableName                      string
	applyExecutionsTableName                     string
	resultsBucketName                            string
}

type OrganizationsDatabaseClientInput struct {
	DynamoDB awsclients.DynamoDBInterface
	S3       awsclients.S3Interface

	DimensionsTableName                          string
	UnitsTableName                               string
	AccountsTableName                            string
	MembershipsTableName                         string
	GroupsTableName                              string
	VersionsTableName                            string
	PropagationsTableName                        string
	ModulePropagationExecutionRequestsTableName  string
	ModulePropagationDriftCheckRequestsTableName string
	ModuleAssignmentsTableName                   string
	ModulePropagationAssignmentsTableName        string

	TerraformExecutionWorkflowRequestsTableName  string
	TerraformDriftCheckWorkflowRequestsTableName string
	PlanExecutionsTableName                      string
	ApplyExecutionsTableName                     string
	ResultsBucketName                            string
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
		modulePropagationExecutionRequestsTableName:  input.ModulePropagationExecutionRequestsTableName,
		modulePropagationDriftCheckRequestsTableName: input.ModulePropagationDriftCheckRequestsTableName,
		moduleAssignmentsTableName:                   input.ModuleAssignmentsTableName,
		modulePropagationAssignmentsTableName:        input.ModulePropagationAssignmentsTableName,
		terraformExecutionWorkflowRequestsTableName:  input.TerraformExecutionWorkflowRequestsTableName,
		terraformDriftCheckWorkflowRequestsTableName: input.TerraformDriftCheckWorkflowRequestsTableName,
		planExecutionsTableName:                      input.PlanExecutionsTableName,
		applyExecutionsTableName:                     input.ApplyExecutionsTableName,
		resultsBucketName:                            input.ResultsBucketName,
	}
}
