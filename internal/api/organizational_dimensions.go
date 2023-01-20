package api

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetOrganizationalDimensionsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetOrganizationalDimensionsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetOrganizationalDimension(ctx context.Context, orgDimensionId string) (*models.OrganizationalDimension, error) {
	return c.dbClient.GetOrganizationalDimension(ctx, orgDimensionId)
}

func (c *APIClient) GetOrganizationalDimensionBatched(ctx context.Context, orgDimensionId string) (*models.OrganizationalDimension, error) {
	thunk := c.orgDimensionsLoader.Load(ctx, dataloader.StringKey(orgDimensionId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.OrganizationalDimension), nil
}

func (c *APIClient) GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error) {
	return c.dbClient.GetOrganizationalDimensions(ctx, limit, cursor)
}

func (c *APIClient) PutOrganizationalDimension(ctx context.Context, input *models.NewOrganizationalDimension) (*models.OrganizationalDimension, error) {
	orgDimensionId, err := identifiers.NewIdentifier(identifiers.ResourceTypeOrganizationalDimension)
	if err != nil {
		return nil, err
	}

	rootOuId, err := identifiers.NewIdentifier(identifiers.ResourceTypeOrganizationalUnit)
	if err != nil {
		return nil, err
	}

	// create the root OU for the dimension
	rootOu := models.OrganizationalUnit{
		OrgUnitId:      rootOuId.String(),
		Name:           fmt.Sprintf("%s Root", input.Name),
		OrgDimensionId: orgDimensionId.String(),
		Hierarchy:      "/",
	}
	err = c.dbClient.PutOrganizationalUnit(ctx, &rootOu)
	if err != nil {
		return nil, err
	}

	orgDimension := models.OrganizationalDimension{
		OrgDimensionId: orgDimensionId.String(),
		Name:           input.Name,
		RootOrgUnitId:  rootOuId.String(),
	}
	err = c.dbClient.PutOrganizationalDimension(ctx, &orgDimension)
	if err != nil {
		return nil, err
	} else {
		return &orgDimension, nil
	}
}

func (c *APIClient) DeleteOrganizationalDimension(ctx context.Context, id string) error {
	return c.dbClient.DeleteOrganizationalDimension(ctx, id)
}
