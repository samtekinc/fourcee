package api

import (
	"context"
	"fmt"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetOrganizationalDimension(ctx context.Context, orgDimensionId string) (*models.OrganizationalDimension, error) {
	return c.dbClient.GetOrganizationalDimension(ctx, orgDimensionId)
}

func (c *OrganizationsAPIClient) GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error) {
	return c.dbClient.GetOrganizationalDimensions(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) PutOrganizationalDimension(ctx context.Context, input *models.NewOrganizationalDimension) (*models.OrganizationalDimension, error) {
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

func (c *OrganizationsAPIClient) DeleteOrganizationalDimension(ctx context.Context, id string) error {
	return c.dbClient.DeleteOrganizationalDimension(ctx, id)
}
