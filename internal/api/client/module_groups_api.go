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

func moduleGroupFilters(filters *models.ModuleGroupFilters) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if filters != nil {
			if filters.NameContains != nil {
				tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
			}
			if filters.CloudPlatform != nil {
				tx = tx.Where("cloud_platform = ?", *filters.CloudPlatform)
			}
		}
		return tx
	}
}

func moduleGroupIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id")
}

func (c *APIClient) GetModuleGroupsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var moduleGroups []*models.ModuleGroup
	tx := c.db.Scopes()
	err := tx.Find(&moduleGroups, keys.Keys()).Error
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

	response := make([]*dataloader.Result, len(moduleGroups))
	for i := range moduleGroups {
		index := keyToIndex[idToString(moduleGroups[i].ID)]
		response[index] = &dataloader.Result{Data: moduleGroups[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("Module Group %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetModuleGroup(ctx context.Context, id uint) (*models.ModuleGroup, error) {
	var moduleGroup models.ModuleGroup
	tx := c.db.Scopes()
	err := tx.First(&moduleGroup, id).Error
	if err != nil {
		return nil, err
	}
	return &moduleGroup, nil
}

func (c *APIClient) GetModuleGroupBatched(ctx context.Context, id uint) (*models.ModuleGroup, error) {
	thunk := c.moduleGroupsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModuleGroup), nil
}

func (c *APIClient) GetModuleGroups(ctx context.Context, filters *models.ModuleGroupFilters, limit *int, offset *int) ([]*models.ModuleGroup, error) {
	var moduleGroups []*models.ModuleGroup
	tx := c.db.Scopes(applyPagination(limit, offset), moduleGroupFilters(filters), moduleGroupIDOrdering)
	err := tx.Find(&moduleGroups).Error
	if err != nil {
		return nil, err
	}
	return moduleGroups, nil
}

func (c *APIClient) CreateModuleGroup(ctx context.Context, input *models.NewModuleGroup) (*models.ModuleGroup, error) {
	moduleGroup := models.ModuleGroup{
		Name:          input.Name,
		CloudPlatform: input.CloudPlatform,
	}
	err := c.db.Create(&moduleGroup).Error
	if err != nil {
		return nil, err
	}

	return &moduleGroup, nil
}

func (c *APIClient) DeleteModuleGroup(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ModuleGroup{}, id).Error
}
