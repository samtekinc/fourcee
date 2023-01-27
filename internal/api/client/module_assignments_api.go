package client

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyModuleAssignmentFilters(tx *gorm.DB, filters *models.ModuleAssignmentFilters) *gorm.DB {
	if filters != nil {
		if filters.NameContains != nil {
			tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
		}
		if filters.DescriptionContains != nil {
			tx = tx.Where("description LIKE ?", "%"+*filters.DescriptionContains+"%")
		}
		if filters.Status != nil {
			tx = tx.Where("status = ?", *filters.Status)
		}
		if filters.IsPropagated != nil {
			if *filters.IsPropagated {
				tx = tx.Where("module_propagation_id IS NOT NULL")
			} else {
				tx = tx.Where("module_propagation_id IS NULL")
			}
		}
	}
	return tx
}

func applyModuleAssignmentPreloads(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Arguments").Preload("AwsProviderConfigurations").Preload("GcpProviderConfigurations")
}

func (c *APIClient) GetModuleAssignmentsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var moduleAssignments []*models.ModuleAssignment
	tx := applyModuleAssignmentPreloads(c.db)
	err := tx.Find(&moduleAssignments, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: moduleAssignments[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetModuleAssignment(ctx context.Context, id uint) (*models.ModuleAssignment, error) {
	var moduleAssignment models.ModuleAssignment
	tx := applyModuleAssignmentPreloads(c.db)
	err := tx.First(&moduleAssignment, id).Error
	if err != nil {
		return nil, err
	}
	return &moduleAssignment, nil
}

func (c *APIClient) GetModuleAssignmentBatched(ctx context.Context, id uint) (*models.ModuleAssignment, error) {
	thunk := c.moduleAssignmentsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModuleAssignment), nil
}

func (c *APIClient) GetModuleAssignments(ctx context.Context, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error) {
	var moduleAssignments []*models.ModuleAssignment
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleAssignmentFilters(tx, filters)
	tx = applyModuleAssignmentPreloads(tx)
	err := tx.Find(&moduleAssignments).Error
	if err != nil {
		return nil, err
	}
	return moduleAssignments, nil
}

func (c *APIClient) GetModuleAssignmentsForModulePropagation(ctx context.Context, modulePropagationId uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error) {
	var moduleAssignments []*models.ModuleAssignment
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleAssignmentFilters(tx, filters)
	tx = applyModuleAssignmentPreloads(tx)
	err := tx.Model(&models.ModulePropagation{Model: gorm.Model{ID: modulePropagationId}}).Association("ModuleAssignmentsAssociation").Find(&moduleAssignments)
	if err != nil {
		return nil, err
	}
	return moduleAssignments, nil
}

func (c *APIClient) GetModuleAssignmentsForModuleGroup(ctx context.Context, moduleGroupId uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error) {
	var moduleAssignments []*models.ModuleAssignment
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleAssignmentFilters(tx, filters)
	tx = applyModuleAssignmentPreloads(tx)
	err := tx.Model(&models.ModuleGroup{Model: gorm.Model{ID: moduleGroupId}}).Association("ModuleAssignmentsAssociation").Find(&moduleAssignments)
	if err != nil {
		return nil, err
	}
	return moduleAssignments, nil
}

func (c *APIClient) GetModuleAssignmentsForModuleVersion(ctx context.Context, moduleVersionId uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error) {
	var moduleAssignments []*models.ModuleAssignment
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleAssignmentFilters(tx, filters)
	tx = applyModuleAssignmentPreloads(tx)
	err := tx.Model(&models.ModuleVersion{Model: gorm.Model{ID: moduleVersionId}}).Association("ModuleAssignmentsAssociation").Find(&moduleAssignments)
	if err != nil {
		return nil, err
	}
	return moduleAssignments, nil
}

func (c *APIClient) GetModuleAssignmentsForOrgAccount(ctx context.Context, orgAccountId uint, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error) {
	var moduleAssignments []*models.ModuleAssignment
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleAssignmentFilters(tx, filters)
	tx = applyModuleAssignmentPreloads(tx)
	err := tx.Model(&models.OrgAccount{Model: gorm.Model{ID: orgAccountId}}).Association("ModuleAssignmentsAssociation").Find(&moduleAssignments)
	if err != nil {
		return nil, err
	}
	return moduleAssignments, nil
}

func (c *APIClient) CreateModuleAssignment(ctx context.Context, input *models.NewModuleAssignment) (*models.ModuleAssignment, error) {
	moduleAssignment := models.ModuleAssignment{
		ModuleVersionID:           input.ModuleVersionID,
		ModuleGroupID:             input.ModuleGroupID,
		OrgAccountID:              input.OrgAccountID,
		Name:                      input.Name,
		Description:               input.Description,
		RemoteStateRegion:         c.remoteStateRegion,
		RemoteStateBucket:         c.remoteStateBucket,
		Status:                    models.ModuleAssignmentStatusActive,
		Arguments:                 ArgumentInputsToArguments(input.Arguments),
		AwsProviderConfigurations: AwsProviderConfigurationInputsToAwsProviderConfigurations(input.AwsProviderConfigurations),
		GcpProviderConfigurations: GcpProviderConfigurationInputsToGcpProviderConfigurations(input.GcpProviderConfigurations),
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&moduleAssignment).Error
		if err != nil {
			return err
		}
		err = tx.Model(&moduleAssignment).Update("remote_state_key", fmt.Sprintf("module-assignments/%v/terraform.tfstate", moduleAssignment.ID)).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &moduleAssignment, nil
}

func (c *APIClient) DeleteModuleAssignment(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ModuleAssignment{}, id).Error
}

func (c *APIClient) UpdateModuleAssignment(ctx context.Context, id uint, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error) {
	moduleAssignment := models.ModuleAssignment{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.ModuleAssignment{}

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
		if update.Status != nil {
			updates.Status = *update.Status
		}

		err := tx.Model(&moduleAssignment).Updates(updates).Error
		if err != nil {
			return err
		}

		if update.Arguments != nil {
			err = tx.Model(&moduleAssignment).Association("Arguments").Replace(ArgumentInputsToArguments(update.Arguments))
			if err != nil {
				return err
			}
		}

		if update.AwsProviderConfigurations != nil {
			err = tx.Model(&moduleAssignment).Association("AwsProviderConfigurations").Replace(AwsProviderConfigurationInputsToAwsProviderConfigurations(update.AwsProviderConfigurations))
			if err != nil {
				return err
			}
		}

		if update.GcpProviderConfigurations != nil {
			err = tx.Model(&moduleAssignment).Association("GcpProviderConfigurations").Replace(GcpProviderConfigurationInputsToGcpProviderConfigurations(update.GcpProviderConfigurations))
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &moduleAssignment, nil
}
