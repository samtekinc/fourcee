package client

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyModuleGroupFilters(tx *gorm.DB, filters *models.ModuleGroupFilters) *gorm.DB {
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

func (c *APIClient) GetModuleGroupsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var moduleGroups []*models.ModuleGroup
	err := c.db.Find(&moduleGroups, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: moduleGroups[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetModuleGroup(ctx context.Context, id uint) (*models.ModuleGroup, error) {
	var moduleGroup models.ModuleGroup
	err := c.db.First(&moduleGroup, id).Error
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
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleGroupFilters(tx, filters)
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
