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

func modulePropagationFilters(filters *models.ModulePropagationFilters) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if filters != nil {
			if filters.NameContains != nil {
				tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
			}
			if filters.DescriptionContains != nil {
				tx = tx.Where("description LIKE ?", "%"+*filters.DescriptionContains+"%")
			}
		}
		return tx
	}
}

func modulePropagationIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id")
}

func (c *APIClient) GetModulePropagationsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var modulePropagations []*models.ModulePropagation
	tx := c.db.Scopes()
	err := tx.Find(&modulePropagations, keys.Keys()).Error
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

	response := make([]*dataloader.Result, len(modulePropagations))
	for i := range modulePropagations {
		index := keyToIndex[idToString(modulePropagations[i].ID)]
		response[index] = &dataloader.Result{Data: modulePropagations[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetModulePropagation(ctx context.Context, id uint) (*models.ModulePropagation, error) {
	var modulePropagation models.ModulePropagation
	tx := c.db.Scopes()
	err := tx.First(&modulePropagation, id).Error
	if err != nil {
		return nil, err
	}
	return &modulePropagation, nil
}

func (c *APIClient) GetModulePropagationBatched(ctx context.Context, id uint) (*models.ModulePropagation, error) {
	thunk := c.modulePropagationsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagation), nil
}

func (c *APIClient) GetModulePropagations(ctx context.Context, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := c.db.Scopes(applyPagination(limit, offset), modulePropagationFilters(filters), modulePropagationIDOrdering)
	err := tx.Find(&modulePropagations).Error
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForModuleGroup(ctx context.Context, moduleGroupId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := c.db.Scopes(applyPagination(limit, offset), modulePropagationFilters(filters), modulePropagationIDOrdering)
	err := tx.Model(&models.ModuleGroup{Model: gorm.Model{ID: moduleGroupId}}).Association("ModulePropagationsAssociation").Find(&modulePropagations)
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForModuleVersion(ctx context.Context, moduleVersionId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := c.db.Scopes(applyPagination(limit, offset), modulePropagationFilters(filters), modulePropagationIDOrdering)
	err := tx.Model(&models.ModuleVersion{Model: gorm.Model{ID: moduleVersionId}}).Association("ModulePropagationsAssociation").Find(&modulePropagations)
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := c.db.Scopes(applyPagination(limit, offset), modulePropagationFilters(filters), modulePropagationIDOrdering)
	err := tx.Model(&models.OrgUnit{Model: gorm.Model{ID: orgUnitId}}).Association("ModulePropagationsAssociation").Find(&modulePropagations)
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForOrgDimension(ctx context.Context, orgDimensionId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := c.db.Scopes(applyPagination(limit, offset), modulePropagationFilters(filters), modulePropagationIDOrdering)
	err := tx.Model(&models.OrgDimension{Model: gorm.Model{ID: orgDimensionId}}).Association("ModulePropagationsAssociation").Find(&modulePropagations)
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) CreateModulePropagation(ctx context.Context, input *models.NewModulePropagation) (*models.ModulePropagation, error) {
	modulePropagation := models.ModulePropagation{
		ModuleVersionID:           input.ModuleVersionID,
		ModuleGroupID:             input.ModuleGroupID,
		OrgUnitID:                 input.OrgUnitID,
		OrgDimensionID:            input.OrgDimensionID,
		Name:                      input.Name,
		Description:               input.Description,
		Arguments:                 ArgumentInputsToArguments(input.Arguments),
		AwsProviderConfigurations: AwsProviderConfigurationInputsToAwsProviderConfigurations(input.AwsProviderConfigurations),
		GcpProviderConfigurations: GcpProviderConfigurationInputsToGcpProviderConfigurations(input.GcpProviderConfigurations),
	}
	err := c.db.Create(&modulePropagation).Error
	if err != nil {
		return nil, err
	}

	return &modulePropagation, nil
}

func (c *APIClient) DeleteModulePropagation(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ModulePropagation{}, id).Error
}

func (c *APIClient) UpdateModulePropagation(ctx context.Context, id uint, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error) {
	modulePropagation := models.ModulePropagation{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.ModulePropagation{}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		// TODO: validate new IDs exist and are valid w/r/t to related IDs
		if update.Name != nil {
			updates.Name = *update.Name
		}
		if update.Description != nil {
			updates.Description = *update.Description
		}
		if update.ModuleVersionID != nil {
			updates.ModuleVersionID = *update.ModuleVersionID
		}
		if update.OrgDimensionID != nil {
			updates.OrgDimensionID = *update.OrgDimensionID
		}
		if update.OrgUnitID != nil {
			updates.OrgUnitID = *update.OrgUnitID
		}
		if update.Arguments != nil {
			updates.Arguments = ArgumentInputsToArguments(update.Arguments)
		}
		if update.AwsProviderConfigurations != nil {
			updates.AwsProviderConfigurations = AwsProviderConfigurationInputsToAwsProviderConfigurations(update.AwsProviderConfigurations)
		}
		if update.GcpProviderConfigurations != nil {
			updates.GcpProviderConfigurations = GcpProviderConfigurationInputsToGcpProviderConfigurations(update.GcpProviderConfigurations)
		}

		err := tx.Model(&modulePropagation).Updates(updates).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &modulePropagation, nil
}
