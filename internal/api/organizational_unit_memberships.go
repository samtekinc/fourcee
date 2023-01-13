package api

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetOrganizationalUnitMembershipsByAccount(ctx context.Context, accountId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	return c.dbClient.GetOrganizationalUnitMembershipsByAccount(ctx, accountId, limit, cursor)
}

func (c *APIClient) GetOrganizationalUnitMembershipsByOrgUnit(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	return c.dbClient.GetOrganizationalUnitMembershipsByOrgUnit(ctx, orgUnitId, limit, cursor)
}

func (c *APIClient) GetOrganizationalUnitMembershipsByDimension(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	return c.dbClient.GetOrganizationalUnitMembershipsByDimension(ctx, orgDimensionId, limit, cursor)
}

func (c *APIClient) PutOrganizationalUnitMembership(ctx context.Context, input *models.NewOrganizationalUnitMembership) (*models.OrganizationalUnitMembership, error) {
	orgUnitMembership := models.OrganizationalUnitMembership{
		OrgAccountId:   input.OrgAccountId,
		OrgDimensionId: input.OrgDimensionId,
		OrgUnitId:      input.OrgUnitId,
	}
	err := c.dbClient.PutOrganizationalUnitMembership(ctx, &orgUnitMembership)
	if err != nil {
		return nil, err
	} else {
		return &orgUnitMembership, nil
	}
}

func (c *APIClient) DeleteOrganizationalUnitMembership(ctx context.Context, orgDimensionId string, accountId string) error {
	return c.dbClient.DeleteOrganizationalUnitMembership(ctx, orgDimensionId, accountId)
}
