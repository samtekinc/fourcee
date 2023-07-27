package client

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/samtekinc/fourcee/internal/helpers"
	"github.com/samtekinc/fourcee/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func orgDimensionFilters(filters *models.OrgDimensionFilters) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if filters != nil {
			if filters.NameContains != nil {
				tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
			}
		}
		return tx
	}
}

func orgDimensionIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id")
}

func (c *APIClient) GetOrgDimensionsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var orgDimensions []*models.OrgDimension
	tx := c.db.Scopes()
	err := tx.Find(&orgDimensions, keys.Keys()).Error
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

	response := make([]*dataloader.Result, len(orgDimensions))
	for i := range orgDimensions {
		index := keyToIndex[idToString(orgDimensions[i].ID)]
		response[index] = &dataloader.Result{Data: orgDimensions[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("Org Dimension %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetOrgDimension(ctx context.Context, id uint) (*models.OrgDimension, error) {
	var orgDimension models.OrgDimension
	tx := c.db.Scopes()
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
	tx := c.db.Scopes(applyPagination(limit, offset), orgDimensionFilters(filters), orgDimensionIDOrdering)
	err := tx.Find(&orgDimensions).Error
	if err != nil {
		return nil, err
	}
	return orgDimensions, nil
}

func (c *APIClient) CreateOrgDimension(ctx context.Context, input *models.NewOrgDimension) (*models.OrgDimension, error) {
	orgDimension := models.OrgDimension{
		Name: input.Name,
		OrgUnitsAssociation: []models.OrgUnit{
			{
				Name:            "Root",
				ParentOrgUnitID: nil,
			},
		},
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		// create org dimension
		err := tx.Create(&orgDimension).Error
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
