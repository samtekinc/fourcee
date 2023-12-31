package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/graph-gophers/dataloader"
	"github.com/samtekinc/fourcee/internal/helpers"
	"github.com/samtekinc/fourcee/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func orgUnitFilters(filters *models.OrgUnitFilters) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if filters != nil {
			if filters.NameContains != nil {
				tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
			}
		}
		return tx
	}
}

func orgUnitIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id")
}

func (c *APIClient) GetOrgUnitsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var orgUnits []*models.OrgUnit
	tx := c.db.Scopes()
	err := tx.Find(&orgUnits, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	keyToIndex := map[string]int{}
	for i := range keys {
		keyToIndex[keys[i].String()] = i
	}

	response := make([]*dataloader.Result, len(orgUnits))
	for i := range orgUnits {
		index := keyToIndex[idToString(orgUnits[i].ID)]
		response[index] = &dataloader.Result{Data: orgUnits[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("Org Unit %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetOrgUnit(ctx context.Context, id uint) (*models.OrgUnit, error) {
	var orgUnit models.OrgUnit
	tx := c.db.Scopes()
	err := tx.First(&orgUnit, id).Error
	if err != nil {
		return nil, err
	}
	return &orgUnit, nil
}

func (c *APIClient) GetOrgUnitBatched(ctx context.Context, id uint) (*models.OrgUnit, error) {
	thunk := c.orgUnitsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.OrgUnit), nil
}

func (c *APIClient) GetOrgUnits(ctx context.Context, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error) {
	var orgUnits []*models.OrgUnit
	tx := c.db.Scopes(applyPagination(limit, offset), orgUnitFilters(filters), orgUnitIDOrdering)
	err := tx.Find(&orgUnits).Error
	if err != nil {
		return nil, err
	}
	return orgUnits, nil
}

func (c *APIClient) GetOrgUnitsForDimension(ctx context.Context, dimensionId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error) {
	var orgUnits []*models.OrgUnit
	tx := c.db.Scopes(applyPagination(limit, offset), orgUnitFilters(filters), orgUnitIDOrdering)
	err := tx.Model(&models.OrgDimension{Model: gorm.Model{ID: dimensionId}}).Association("OrgUnitsAssociation").Find(&orgUnits)
	if err != nil {
		return nil, err
	}
	return orgUnits, nil
}

func (c *APIClient) GetOrgUnitsForParent(ctx context.Context, parentOrgUnitId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error) {
	var orgUnits []*models.OrgUnit
	tx := c.db.Scopes(applyPagination(limit, offset), orgUnitFilters(filters), orgUnitIDOrdering)
	err := tx.Model(&models.OrgUnit{Model: gorm.Model{ID: parentOrgUnitId}}).Association("ChildOrgUnitsAssociation").Find(&orgUnits)
	if err != nil {
		return nil, err
	}
	return orgUnits, nil
}

func (c *APIClient) GetDownstreamOrgUnits(ctx context.Context, orgUnitId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error) {
	var orgUnits []*models.OrgUnit
	err := c.db.Where("hierarchy LIKE ?", gorm.Expr("(?) || ':' || (?) || '%'", c.db.Model(&models.OrgUnit{Model: gorm.Model{ID: orgUnitId}}).Select("hierarchy"), orgUnitId)).Scopes(applyPagination(limit, offset), orgUnitFilters(filters), orgUnitIDOrdering).Find(&orgUnits).Error
	if err != nil {
		return nil, err
	}
	return orgUnits, nil
}

func (c *APIClient) GetUpstreamOrgUnits(ctx context.Context, orgUnitId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error) {
	// get the hierarchy of the org unit
	var orgUnit models.OrgUnit
	err := c.db.First(&orgUnit, orgUnitId).Error
	if err != nil {
		return nil, err
	}
	hierarchy := orgUnit.Hierarchy
	if hierarchy == "" {
		return nil, nil
	}
	parentOrgUnitIds := strings.Split(hierarchy, ":")[1:] // remove the first element, which is always an empty string

	var orgUnits []*models.OrgUnit
	tx := c.db.Scopes(applyPagination(limit, offset), orgUnitFilters(filters), orgUnitIDOrdering)
	err = tx.Find(&orgUnits, parentOrgUnitIds).Error
	if err != nil {
		return nil, err
	}
	return orgUnits, nil
}

func (c *APIClient) GetOrgUnitsForOrgAccount(ctx context.Context, orgAccountId uint, filters *models.OrgUnitFilters, limit *int, offset *int) ([]*models.OrgUnit, error) {
	var orgUnits []*models.OrgUnit
	tx := c.db.Scopes(applyPagination(limit, offset), orgUnitFilters(filters), orgUnitIDOrdering)
	err := tx.Model(&models.OrgAccount{Model: gorm.Model{ID: orgAccountId}}).Association("OrgUnitsAssociation").Find(&orgUnits)
	if err != nil {
		return nil, err
	}
	return orgUnits, nil
}

func (c *APIClient) CreateOrgUnit(ctx context.Context, input *models.NewOrgUnit) (*models.OrgUnit, error) {
	orgUnit := models.OrgUnit{
		Name:            input.Name,
		OrgDimensionID:  input.OrgDimensionID,
		ParentOrgUnitID: &input.ParentOrgUnitID,
	}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&orgUnit).Error
		if err != nil {
			return err
		}

		err = tx.Model(&orgUnit).Debug().Update("hierarchy", gorm.Expr("(?) || ':' || ?", tx.Model(&models.OrgUnit{Model: gorm.Model{ID: input.ParentOrgUnitID}}).Select("hierarchy"), input.ParentOrgUnitID)).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &orgUnit, nil
}

func (c *APIClient) DeleteOrgUnit(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.OrgUnit{}, id).Error
}

func (c *APIClient) UpdateOrgUnit(ctx context.Context, id uint, update *models.OrgUnitUpdate) (*models.OrgUnit, error) {
	updates := map[string]interface{}{}
	orgUnit := &models.OrgUnit{Model: gorm.Model{ID: id}}

	// if we update the parent org unit, we need to update the hierarchy of all of it's downstream org units to reflect the new location

	/* EXAMPLE
		Start with this org unit hierarchy
		     1
		    / \
		   2   3
		  / \
		 4   5
		/ \
	   6   7
	  /
	 8

		Move 4 to be a child of 3
		New Org Unit Hierarchy would look like
			 1
		    / \
		   2   3
		  /     \
		 5       4
		        / \
			   6   7
		      /
		     8

	   If we update the parent org unit of 4 to be 3, then we need to update the hierarchy of 4, 6, 7, and 8 to reflect the new location

	   4 is simple - it gets updated to equal 3's hierarchy + 3, i.e. :1:3
	   The downstream org units of 4 (6, 7, and 8) are trickier - we need to update them replace 4's old hierarchy with 4's new hierarchy, while keeping the rest of the hierarchy the same
	   	6 and 7 become (6/7s old hierarchy with 4's bit subbed out)->(REPLACE(:1:2:4, :1:2:, :1:3:)) = :1:3:4
		8 becomes (8's old hierarchy with 4's bit subbed out)->(REPLACE(:1:2:4:6, :1:2, :1:3)) = :1:3:4:6
	*/
	err := c.db.Transaction(func(tx *gorm.DB) error {
		if update.Name != nil {
			updates["name"] = *update.Name
		}
		if update.ParentOrgUnitID != nil {
			updates["parent_org_unit_id"] = update.ParentOrgUnitID
			updates["hierarchy"] = gorm.Expr("(?) || ':' || (?)", tx.Model(&models.OrgUnit{Model: gorm.Model{ID: *update.ParentOrgUnitID}}).Select("hierarchy"), *update.ParentOrgUnitID) // update this org units hierarchy to be the hierarchy of the new parent org unit
			// update the hierarchy of all the children
			err := tx.Debug().Model(&models.OrgUnit{}). // select all org units
									Where("hierarchy LIKE ?", gorm.Expr("(?) || ':' || (?) || '%'", tx.Model(&models.OrgUnit{Model: gorm.Model{ID: id}}).Select("hierarchy"), id)). // where the hierarchy starts with this org unit (i.e. all downstream org units of this org unit)
									Update("hierarchy",                                                                                                                             // update the hierarchy
					gorm.Expr("(REPLACE(hierarchy, (?), ((?) || ':' || ?)) )", // substitute the old parents hierarchy with the new parents hierarchy
						tx.Model(&models.OrgUnit{Model: gorm.Model{ID: id}}).Select("hierarchy"),
						tx.Model(&models.OrgUnit{Model: gorm.Model{ID: *update.ParentOrgUnitID}}).Select("hierarchy"),
						*update.ParentOrgUnitID,
					),
				).
				Error
			if err != nil {
				return err
			}
		}

		err := tx.Model(orgUnit).Updates(updates).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return orgUnit, nil
}

func (c *APIClient) AddAccountToOrgUnit(ctx context.Context, orgUnitId uint, orgAccountId uint) error {
	// TODO: check if the account is already associated to an org unit in the same dimension
	return c.db.Model(&models.OrgUnit{Model: gorm.Model{ID: orgUnitId}}).Association("OrgAccountsAssociation").Append(&models.OrgAccount{Model: gorm.Model{ID: orgAccountId}})
}

func (c *APIClient) RemoveAccountFromOrgUnit(ctx context.Context, orgUnitId uint, orgAccountId uint) error {
	return c.db.Model(&models.OrgUnit{Model: gorm.Model{ID: orgUnitId}}).Association("OrgAccountsAssociation").Delete(&models.OrgAccount{Model: gorm.Model{ID: orgAccountId}})
}
