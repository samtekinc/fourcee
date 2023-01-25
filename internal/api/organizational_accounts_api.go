package api

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyOrgAccountFilters(tx *gorm.DB, filters *models.OrgAccountFilters) *gorm.DB {
	if filters != nil {
		if filters.NameContains != nil {
			tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
		}
		if filters.CloudPlatform != nil {
			tx = tx.Where("cloud_platform = ?", *filters.CloudPlatform)
		}
		if filters.CloudIdentifier != nil {
			tx = tx.Where("cloud_identifier = ?", *filters.CloudIdentifier)
		}
	}
	return tx
}

func applyOrgAccountPreloads(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Metadata")
}

func (c *APIClient) GetOrgAccountsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var orgAccounts []*models.OrgAccount
	tx := applyOrgAccountPreloads(c.db)
	err := tx.Find(&orgAccounts, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: orgAccounts[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetOrgAccount(ctx context.Context, id uint) (*models.OrgAccount, error) {
	var orgAccount models.OrgAccount
	tx := applyOrgAccountPreloads(c.db)
	err := tx.First(&orgAccount, id).Error
	if err != nil {
		return nil, err
	}
	return &orgAccount, nil
}

func (c *APIClient) GetOrgAccountBatched(ctx context.Context, id uint) (*models.OrgAccount, error) {
	thunk := c.orgAccountsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.OrgAccount), nil
}

func (c *APIClient) GetOrgAccounts(ctx context.Context, filters *models.OrgAccountFilters, limit *int, offset *int) ([]*models.OrgAccount, error) {
	var orgAccounts []*models.OrgAccount
	tx := applyPagination(c.db, limit, offset)
	tx = applyOrgAccountFilters(tx, filters)
	tx = applyOrgAccountPreloads(tx)
	err := tx.Find(&orgAccounts).Error
	if err != nil {
		return nil, err
	}
	return orgAccounts, nil
}

func (c *APIClient) GetOrgAccountsForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.OrgAccountFilters, limit *int, offset *int) ([]*models.OrgAccount, error) {
	var orgAccounts []*models.OrgAccount
	tx := applyPagination(c.db, limit, offset)
	tx = applyOrgAccountFilters(tx, filters)
	err := tx.Model(&models.OrgUnit{Model: gorm.Model{ID: orgUnitId}}).Association("OrgAccountsAssociation").Find(&orgAccounts)
	if err != nil {
		return nil, err
	}
	return orgAccounts, nil
}

func (c *APIClient) CreateOrgAccount(ctx context.Context, input *models.NewOrgAccount) (*models.OrgAccount, error) {
	orgAccount := models.OrgAccount{
		Name:            input.Name,
		CloudPlatform:   input.CloudPlatform,
		CloudIdentifier: input.CloudIdentifier,
		AssumeRoleName:  input.AssumeRoleName,
		Metadata:        MetadataInputsToMetadata(input.Metadata),
	}
	err := c.db.Create(&orgAccount).Error
	if err != nil {
		return nil, err
	}
	return &orgAccount, nil
}

func (c *APIClient) DeleteOrgAccount(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.OrgAccount{}, id).Error
}

func (c *APIClient) UpdateOrgAccount(ctx context.Context, id uint, update *models.OrgAccountUpdate) (*models.OrgAccount, error) {
	orgAccount := models.OrgAccount{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.OrgAccount{}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		if update.Name != nil {
			updates.Name = *update.Name
		}
		if update.CloudPlatform != nil {
			updates.CloudPlatform = *update.CloudPlatform
		}
		if update.CloudIdentifier != nil {
			updates.CloudIdentifier = *update.CloudIdentifier
		}
		if update.AssumeRoleName != nil {
			updates.AssumeRoleName = *update.AssumeRoleName
		}

		err := tx.Model(&orgAccount).Updates(updates).Error
		if err != nil {
			return err
		}

		if update.Metadata != nil {
			err = tx.Model(&orgAccount).Association("Metadata").Replace(MetadataInputsToMetadata(update.Metadata))
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &orgAccount, nil
}
