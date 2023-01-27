package client

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyModulePropagationFilters(tx *gorm.DB, filters *models.ModulePropagationFilters) *gorm.DB {
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

func applyModulePropagationPreloads(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Arguments").Preload("AwsProviderConfigurations").Preload("GcpProviderConfigurations")
}

func (c *APIClient) GetModulePropagationsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var modulePropagations []*models.ModulePropagation
	tx := applyModulePropagationPreloads(c.db)
	err := tx.Find(&modulePropagations, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: modulePropagations[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetModulePropagation(ctx context.Context, id uint) (*models.ModulePropagation, error) {
	var modulePropagation models.ModulePropagation
	tx := applyModulePropagationPreloads(c.db)
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
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationFilters(tx, filters)
	tx = applyModulePropagationPreloads(tx)
	err := tx.Find(&modulePropagations).Error
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForModuleGroup(ctx context.Context, moduleGroupId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationFilters(tx, filters)
	tx = applyModulePropagationPreloads(tx)
	err := tx.Model(&models.ModuleGroup{Model: gorm.Model{ID: moduleGroupId}}).Association("ModulePropagationsAssociation").Find(&modulePropagations)
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForModuleVersion(ctx context.Context, moduleVersionId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationFilters(tx, filters)
	tx = applyModulePropagationPreloads(tx)
	err := tx.Model(&models.ModuleVersion{Model: gorm.Model{ID: moduleVersionId}}).Association("ModulePropagationsAssociation").Find(&modulePropagations)
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationFilters(tx, filters)
	tx = applyModulePropagationPreloads(tx)
	err := tx.Model(&models.OrgUnit{Model: gorm.Model{ID: orgUnitId}}).Association("ModulePropagationsAssociation").Find(&modulePropagations)
	if err != nil {
		return nil, err
	}
	return modulePropagations, nil
}

func (c *APIClient) GetModulePropagationsForOrgDimension(ctx context.Context, orgDimensionId uint, filters *models.ModulePropagationFilters, limit *int, offset *int) ([]*models.ModulePropagation, error) {
	var modulePropagations []*models.ModulePropagation
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationFilters(tx, filters)
	tx = applyModulePropagationPreloads(tx)
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

		err := tx.Model(&modulePropagation).Updates(updates).Error
		if err != nil {
			return err
		}

		if update.Arguments != nil {
			err = tx.Model(&modulePropagation).Association("Arguments").Replace(ArgumentInputsToArguments(update.Arguments))
			if err != nil {
				return err
			}
		}

		if update.AwsProviderConfigurations != nil {
			err = tx.Model(&modulePropagation).Association("AwsProviderConfigurations").Replace(AwsProviderConfigurationInputsToAwsProviderConfigurations(update.AwsProviderConfigurations))
			if err != nil {
				return err
			}
		}

		if update.GcpProviderConfigurations != nil {
			err = tx.Model(&modulePropagation).Association("GcpProviderConfigurations").Replace(GcpProviderConfigurationInputsToGcpProviderConfigurations(update.GcpProviderConfigurations))
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &modulePropagation, nil
}
