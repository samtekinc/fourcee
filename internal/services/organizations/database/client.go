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
}

type OrganizationsDatabaseClient struct {
	dynamodb             awsclients.DynamoDBInterface
	dimensionsTableName  string
	unitsTableName       string
	accountsTableName    string
	membershipsTableName string
}

func NewOrganizationsDatabaseClient(dynamodb awsclients.DynamoDBInterface, dimensionsTableName string, unitsTableName string, accountsTableName string, membershipsTableName string) *OrganizationsDatabaseClient {
	return &OrganizationsDatabaseClient{
		dynamodb:             dynamodb,
		dimensionsTableName:  dimensionsTableName,
		unitsTableName:       unitsTableName,
		accountsTableName:    accountsTableName,
		membershipsTableName: membershipsTableName,
	}
}
