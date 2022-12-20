package api

import (
	"context"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/organizations/models"
)

func (c *OrganizationsAPIClient) GetOrganizationalAccount(ctx context.Context, id string) (*models.OrganizationalAccount, error) {
	return c.dbClient.GetOrganizationalAccount(ctx, id)
}

func (c *OrganizationsAPIClient) GetOrganizationalAccounts(ctx context.Context, limit int32, cursor string) (*models.OrganizationalAccounts, error) {
	return c.dbClient.GetOrganizationalAccounts(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) PutOrganizationalAccount(ctx context.Context, input *models.NewOrganizationalAccount) (*models.OrganizationalAccount, error) {
	accountId, err := identifiers.NewIdentifier(identifiers.ResourceTypeOrganizationalAccount)
	if err != nil {
		return nil, err
	}

	orgAccount := models.OrganizationalAccount{
		OrgAccountId: accountId.String(),
		Name:         input.Name,
	}
	err = c.dbClient.PutOrganizationalAccount(ctx, &orgAccount)
	if err != nil {
		return nil, err
	} else {
		return &orgAccount, nil
	}
}

func (c *OrganizationsAPIClient) DeleteOrganizationalAccount(ctx context.Context, id string) error {
	return c.dbClient.DeleteOrganizationalAccount(ctx, id)
}
