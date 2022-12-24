package database

import (
	"context"

	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/pkg/organizations/models"
)

type OrganizationsDatabaseClientInterface interface {
	GetOrganizationalDimension(ctx context.Context, orgDimensionId string) (*models.OrganizationalDimension, error)
	GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error)
	PutOrganizationalDimension(ctx context.Context, input *models.OrganizationalDimension) error
	DeleteOrganizationalDimension(ctx context.Context, orgDimensionId string) error

	GetOrganizationalUnit(ctx context.Context, orgUnitId string) (*models.OrganizationalUnit, error)
	GetOrganizationalUnits(ctx context.Context, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByDimension(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByParent(ctx context.Context, orgDimensionId string, parentOrgUnitId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByHierarchy(ctx context.Context, orgDimensionId string, hierarchy string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	PutOrganizationalUnit(ctx context.Context, input *models.OrganizationalUnit) error
	DeleteOrganizationalUnit(ctx context.Context, orgUnitId string) error
	UpdateOrganizationalUnit(ctx context.Context, orgUnitId string, update *OrganizationalUnitUpdate) (*models.OrganizationalUnit, error)

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

	GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string) (*models.ModulePropagationExecutionRequest, error)
	GetModulePropagationExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	GetModulePropagationExecutionRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error)
	PutModulePropagationExecutionRequest(ctx context.Context, input *models.ModulePropagationExecutionRequest) error
	UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error)
}

type OrganizationsDatabaseClient struct {
	dynamodb                                    awsclients.DynamoDBInterface
	dimensionsTableName                         string
	unitsTableName                              string
	accountsTableName                           string
	membershipsTableName                        string
	groupsTableName                             string
	versionsTableName                           string
	propagationsTableName                       string
	modulePropagationExecutionRequestsTableName string
}

type OrganizationsDatabaseClientInput struct {
	DynamoDB                                    awsclients.DynamoDBInterface
	DimensionsTableName                         string
	UnitsTableName                              string
	AccountsTableName                           string
	MembershipsTableName                        string
	GroupsTableName                             string
	VersionsTableName                           string
	PropagationsTableName                       string
	ModulePropagationExecutionRequestsTableName string
}

func NewOrganizationsDatabaseClient(input *OrganizationsDatabaseClientInput) *OrganizationsDatabaseClient {
	return &OrganizationsDatabaseClient{
		dynamodb:              input.DynamoDB,
		dimensionsTableName:   input.DimensionsTableName,
		unitsTableName:        input.UnitsTableName,
		accountsTableName:     input.AccountsTableName,
		membershipsTableName:  input.MembershipsTableName,
		groupsTableName:       input.GroupsTableName,
		versionsTableName:     input.VersionsTableName,
		propagationsTableName: input.PropagationsTableName,
		modulePropagationExecutionRequestsTableName: input.ModulePropagationExecutionRequestsTableName,
	}
}
