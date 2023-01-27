package client

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyOrgDimensionFilters(tx *gorm.DB, filters *models.OrgDimensionFilters) *gorm.DB {
	if filters != nil {
		if filters.NameContains != nil {
			tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
		}
	}
	return tx
}

func applyOrgDimensionPreloads(tx *gorm.DB) *gorm.DB {
	return tx
}

func (c *APIClient) GetOrgDimensionsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var orgDimensions []*models.OrgDimension
	tx := applyOrgDimensionPreloads(c.db)
	err := tx.Find(&orgDimensions, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: orgDimensions[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetOrgDimension(ctx context.Context, id uint) (*models.OrgDimension, error) {
	var orgDimension models.OrgDimension
	tx := applyOrgDimensionPreloads(c.db)
	err := tx.First(&orgDimension, id).Error
	if err != nil {
		return nil, err
	}
	return &orgDimension, nil
}

func (c *APIClient) GetOrgDimensionBatched(ctx context.Context, id uint) (*models.OrgDimension, error) {
	thunk := c.orgDimensionsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.OrgDimension), nil
}

func (c *APIClient) GetOrgDimensions(ctx context.Context, filters *models.OrgDimensionFilters, limit *int, offset *int) ([]*models.OrgDimension, error) {
	var orgDimensions []*models.OrgDimension
	tx := applyPagination(c.db, limit, offset)
	tx = applyOrgDimensionFilters(tx, filters)
	tx = applyOrgDimensionPreloads(tx)
	err := tx.Find(&orgDimensions).Error
	if err != nil {
		return nil, err
	}
	return orgDimensions, nil
}

func (c *APIClient) CreateOrgDimension(ctx context.Context, input *models.NewOrgDimension) (*models.OrgDimension, error) {
	orgDimension := models.OrgDimension{
		Name: input.Name,
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		// create org dimension
		err := tx.Create(&orgDimension).Error
		if err != nil {
			return err
		}

		// create root org unit
		orgUnit := models.OrgUnit{
			Name:            "Root",
			OrgDimensionID:  orgDimension.ID,
			ParentOrgUnitID: nil,
		}
		err = tx.Create(&orgUnit).Error
		if err != nil {
			return err
		}

		// update org dimension with root org unit id
		err = tx.Model(&orgDimension).Updates(models.OrgDimension{RootOrgUnitID: orgUnit.ID}).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &orgDimension, nil
}

func (c *APIClient) DeleteOrgDimension(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.OrgDimension{}, id).Error
}
