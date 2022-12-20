package api

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/internal/services/organizations/database"
	"github.com/sheacloud/tfom/pkg/organizations/models"
)

func (c *OrganizationsAPIClient) GetOrganizationalUnit(ctx context.Context, id string) (*models.OrganizationalUnit, error) {
	return c.dbClient.GetOrganizationalUnit(ctx, id)
}

func (c *OrganizationsAPIClient) GetOrganizationalUnits(ctx context.Context, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	return c.dbClient.GetOrganizationalUnits(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetOrganizationalUnitsByDimension(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	return c.dbClient.GetOrganizationalUnitsByDimension(ctx, orgDimensionId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetOrganizationalUnitsByParent(ctx context.Context, orgDimensionId string, parentOrgUnitId string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	return c.dbClient.GetOrganizationalUnitsByParent(ctx, orgDimensionId, parentOrgUnitId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetOrganizationalUnitsByHierarchy(ctx context.Context, orgDimensionId string, hierarchy string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	return c.dbClient.GetOrganizationalUnitsByHierarchy(ctx, orgDimensionId, hierarchy, limit, cursor)
}

func (c *OrganizationsAPIClient) PutOrganizationalUnit(ctx context.Context, input *models.NewOrganizationalUnit) (*models.OrganizationalUnit, error) {
	parentOrgUnit, err := c.dbClient.GetOrganizationalUnit(ctx, input.ParentOrgUnitId)
	if err != nil {
		return nil, err
	}

	if parentOrgUnit.OrgDimensionId != input.OrgDimensionId {
		return nil, errors.New("parent org unit and new org unit must be in the same dimension") // FIXME: use custom error type
	}

	id, err := identifiers.NewIdentifier(identifiers.ResourceTypeOrganizationalUnit)
	if err != nil {
		return nil, err
	}

	hierarchy := parentOrgUnit.Hierarchy + input.ParentOrgUnitId + "/"

	orgUnit := models.OrganizationalUnit{
		OrgUnitId:       id.String(),
		Name:            input.Name,
		OrgDimensionId:  input.OrgDimensionId,
		Hierarchy:       hierarchy,
		ParentOrgUnitId: input.ParentOrgUnitId,
	}
	err = c.dbClient.PutOrganizationalUnit(ctx, &orgUnit)
	if err != nil {
		return nil, err
	} else {
		return &orgUnit, nil
	}
}

func (c *OrganizationsAPIClient) DeleteOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string) error {
	children, err := c.dbClient.GetOrganizationalUnitsByParent(ctx, orgDimensionId, orgUnitId, 1, "")
	if err != nil {
		return err
	}
	if len(children.Items) > 0 {
		return errors.New("cannot delete org unit with children") // FIXME: use custom error type
	}

	return c.dbClient.DeleteOrganizationalUnit(ctx, orgUnitId)
}

func (c *OrganizationsAPIClient) UpdateOrganizationalUnit(ctx context.Context, id string, update *models.OrganizationalUnitUpdate) (*models.OrganizationalUnit, error) {
	originalOrgUnit, err := c.dbClient.GetOrganizationalUnit(ctx, id)
	if err != nil {
		return nil, err
	}

	orgUnitUpdates := database.OrganizationalUnitUpdate{}
	if update.Name != nil {
		orgUnitUpdates.Name = update.Name
	}
	if update.ParentOrgUnitId != nil {
		if originalOrgUnit.Hierarchy == "/" {
			return nil, errors.New("cannot move root org unit") // FIXME: use custom error type
		}
		parentOrgUnit, err := c.dbClient.GetOrganizationalUnit(ctx, *update.ParentOrgUnitId)
		if err != nil {
			return nil, err
		}

		orgUnitUpdates.ParentOrgUnitId = update.ParentOrgUnitId
		orgUnitUpdates.Hierarchy = aws.String(fmt.Sprintf("%s%s/", parentOrgUnit.Hierarchy, parentOrgUnit.OrgUnitId))
	}

	newOrgUnit, err := c.dbClient.UpdateOrganizationalUnit(ctx, originalOrgUnit.OrgUnitId, &orgUnitUpdates)
	if err != nil {
		return nil, err
	}

	// TODO: rewrite this to be atomic via transactions, or a workflow with retry logic
	// IDEA: trigger a hierarchy recalculation workflow. Lock all Org use during the workflow, recalculate all hierarchies, then unlock
	// don't allow use of the Org chart during this time
	if update.ParentOrgUnitId != nil {
		oldHierarchy := originalOrgUnit.Hierarchy + originalOrgUnit.OrgUnitId
		newHierarchy := newOrgUnit.Hierarchy + newOrgUnit.OrgUnitId
		// update all downstream OUs hierarchies
		nextCursor := ""
		for {
			// get all OUs undernearth the old hierarchy
			ous, err := c.dbClient.GetOrganizationalUnitsByHierarchy(ctx, originalOrgUnit.OrgDimensionId, oldHierarchy, 100, nextCursor)
			if err != nil {
				return nil, err
			}
			for _, ou := range ous.Items {
				// update the OU hierarchy
				childUpdate := database.OrganizationalUnitUpdate{
					Hierarchy: aws.String(strings.Replace(ou.Hierarchy, oldHierarchy, newHierarchy, 1)),
				}
				_, err := c.dbClient.UpdateOrganizationalUnit(ctx, ou.OrgUnitId, &childUpdate)
				if err != nil {
					return nil, err
				}
			}
			if ous.NextCursor == "" {
				break
			} else {
				nextCursor = ous.NextCursor
			}
		}
	}

	return newOrgUnit, nil
}

func (c *OrganizationsAPIClient) UpdateOrganizationalUnitHierarchies(ctx context.Context, orgDimensionId string) error {
	dimension, err := c.dbClient.GetOrganizationalDimension(ctx, orgDimensionId)
	if err != nil {
		return err
	}

	rootOu, err := c.dbClient.GetOrganizationalUnit(ctx, dimension.RootOrgUnitId)
	if err != nil {
		return err
	}

	return updateHierarchy(ctx, c.dbClient, rootOu, nil)
}

func updateHierarchy(ctx context.Context, dbClient database.OrganizationsDatabaseClientInterface, ou *models.OrganizationalUnit, parentOu *models.OrganizationalUnit) error {
	ouUpdate := database.OrganizationalUnitUpdate{}
	if parentOu != nil {
		ouUpdate.Hierarchy = aws.String(parentOu.Hierarchy + parentOu.OrgUnitId + "/")
	} else {
		ouUpdate.Hierarchy = aws.String("/")
	}

	if ouUpdate.Hierarchy != &ou.Hierarchy {
		var err error
		ou, err = dbClient.UpdateOrganizationalUnit(ctx, ou.OrgUnitId, &ouUpdate)
		if err != nil {
			return err
		}
	}

	nextCursor := ""
	for {
		children, err := dbClient.GetOrganizationalUnitsByParent(ctx, ou.OrgDimensionId, ou.OrgUnitId, 100, nextCursor)
		if err != nil {
			return err
		}
		for _, child := range children.Items {
			err := updateHierarchy(ctx, dbClient, &child, ou)
			if err != nil {
				return err
			}
		}
		if children.NextCursor == "" {
			break
		} else {
			nextCursor = children.NextCursor
		}
	}

	return nil
}
