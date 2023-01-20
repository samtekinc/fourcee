package api

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetOrganizationalAccountsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetOrganizationalAccountsByIds(ctx, keys.Keys())
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: &results[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetOrganizationalAccount(ctx context.Context, id string) (*models.OrganizationalAccount, error) {
	return c.dbClient.GetOrganizationalAccount(ctx, id)
}

func (c *APIClient) GetOrganizationalAccountBatched(ctx context.Context, id string) (*models.OrganizationalAccount, error) {
	thunk := c.orgAccountsLoader.Load(ctx, dataloader.StringKey(id))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.OrganizationalAccount), nil
}

func (c *APIClient) GetOrganizationalAccounts(ctx context.Context, limit int32, cursor string) (*models.OrganizationalAccounts, error) {
	return c.dbClient.GetOrganizationalAccounts(ctx, limit, cursor)
}

func (c *APIClient) PutOrganizationalAccount(ctx context.Context, input *models.NewOrganizationalAccount) (*models.OrganizationalAccount, error) {
	accountId, err := identifiers.NewIdentifier(identifiers.ResourceTypeOrganizationalAccount)
	if err != nil {
		return nil, err
	}

	orgAccount := models.OrganizationalAccount{
		OrgAccountId:    accountId.String(),
		Name:            input.Name,
		CloudPlatform:   input.CloudPlatform,
		CloudIdentifier: input.CloudIdentifier,
		AssumeRoleName:  input.AssumeRoleName,
		Metadata:        MetadataInputsToMetadata(input.Metadata),
	}
	err = c.dbClient.PutOrganizationalAccount(ctx, &orgAccount)
	if err != nil {
		return nil, err
	} else {
		return &orgAccount, nil
	}
}

func (c *APIClient) DeleteOrganizationalAccount(ctx context.Context, id string) error {
	return c.dbClient.DeleteOrganizationalAccount(ctx, id)
}

func (c *APIClient) UpdateOrganizationalAccount(ctx context.Context, orgAccountId string, update *models.OrganizationalAccountUpdate) (*models.OrganizationalAccount, error) {
	return c.dbClient.UpdateOrganizationalAccount(ctx, orgAccountId, update)
}
