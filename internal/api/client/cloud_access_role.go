package client

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func cloudAccessRoleFilters(filters *models.CloudAccessRoleFilters) func(tx *gorm.DB) *gorm.DB {
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

func cloudAccessRoleIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id")
}

func (c *APIClient) GetCloudAccessRolesByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var cloudAccessRoles []*models.CloudAccessRole
	tx := c.db.Scopes()
	err := tx.Find(&cloudAccessRoles, keys.Keys()).Error
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

	response := make([]*dataloader.Result, len(cloudAccessRoles))
	for i := range cloudAccessRoles {
		index := keyToIndex[idToString(cloudAccessRoles[i].ID)]
		response[index] = &dataloader.Result{Data: cloudAccessRoles[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("AWS IAM Policy %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetCloudAccessRole(ctx context.Context, id uint) (*models.CloudAccessRole, error) {
	var cloudAccessRole models.CloudAccessRole
	tx := c.db.Scopes()
	err := tx.First(&cloudAccessRole, id).Error
	if err != nil {
		return nil, err
	}
	return &cloudAccessRole, nil
}

func (c *APIClient) GetCloudAccessRoleBatched(ctx context.Context, id uint) (*models.CloudAccessRole, error) {
	thunk := c.cloudAccessRoleLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.CloudAccessRole), nil
}

func (c *APIClient) GetCloudAccessRoles(ctx context.Context, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error) {
	var cloudAccessRoles []*models.CloudAccessRole
	tx := c.db.Scopes(applyPagination(limit, offset), cloudAccessRoleFilters(filters), cloudAccessRoleIDOrdering)
	err := tx.Find(&cloudAccessRoles).Error
	if err != nil {
		return nil, err
	}
	return cloudAccessRoles, nil
}

func (c *APIClient) CreateCloudAccessRole(ctx context.Context, input *models.NewCloudAccessRole) (*models.CloudAccessRole, error) {
	cloudAccessRole := models.CloudAccessRole{
		Name:                      input.Name,
		CloudPlatform:             input.CloudPlatform,
		OrgUnitID:                 input.OrgUnitID,
		AwsIamPoliciesAssociation: []models.AwsIamPolicy{},
	}
	for _, policyID := range input.AwsIamPolicies {
		cloudAccessRole.AwsIamPoliciesAssociation = append(cloudAccessRole.AwsIamPoliciesAssociation, models.AwsIamPolicy{Model: gorm.Model{ID: policyID}})
	}
	err := c.db.Create(&cloudAccessRole).Error
	if err != nil {
		return nil, err
	}

	return &cloudAccessRole, nil
}

func (c *APIClient) DeleteCloudAccessRole(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.CloudAccessRole{}, id).Error
}

func (c *APIClient) UpdateCloudAccessRole(ctx context.Context, id uint, update *models.CloudAccessRoleUpdate) (*models.CloudAccessRole, error) {
	cloudAccessRole := models.CloudAccessRole{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.CloudAccessRole{}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		if update.Name != nil {
			updates.Name = *update.Name
		}
		if update.OrgUnitID != nil {
			updates.OrgUnitID = *update.OrgUnitID
		}
		err := tx.Model(&cloudAccessRole).Updates(updates).Error
		if err != nil {
			return err
		}

		// update policies assocation
		if update.AwsIamPolicies != nil {
			updates.AwsIamPoliciesAssociation = []models.AwsIamPolicy{}
			for _, policyID := range *update.AwsIamPolicies {
				updates.AwsIamPoliciesAssociation = append(updates.AwsIamPoliciesAssociation, models.AwsIamPolicy{Model: gorm.Model{ID: policyID}})
			}
			err := tx.Model(&cloudAccessRole).Association("AwsIamPoliciesAssociation").Replace(updates.AwsIamPoliciesAssociation)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &cloudAccessRole, nil
}

func (c *APIClient) GetCloudAccessRolesForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error) {
	var cloudAccessRoles []*models.CloudAccessRole
	tx := c.db.Scopes(applyPagination(limit, offset), cloudAccessRoleFilters(filters), cloudAccessRoleIDOrdering)
	err := tx.Model(&models.OrgUnit{Model: gorm.Model{ID: orgUnitId}}).Association("CloudAccessRolesAssociation").Find(&cloudAccessRoles)
	if err != nil {
		return nil, err
	}
	return cloudAccessRoles, nil
}

func (c *APIClient) GetInheritedCloudAccessRolesForOrgUnit(ctx context.Context, orgUnitId uint, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error) {
	var cloudAccessRoles []*models.CloudAccessRole
	tx := c.db.Scopes(applyPagination(limit, offset), cloudAccessRoleFilters(filters), cloudAccessRoleIDOrdering)
	// subquery to get all upstream org units
	subQuery := tx.Debug().Select("unnest(string_to_array(trim(leading ':' from hierarchy), ':')::bigint[])").Table("org_units").Where("id = ?", orgUnitId)
	// find all cloud access roles associated with the org units found in the subquery
	err := tx.Debug().Model(&models.CloudAccessRole{}).Select("*").Where("org_unit_id IN (?)", subQuery).Find(&cloudAccessRoles).Error
	if err != nil {
		return nil, err
	}
	return cloudAccessRoles, nil
}

func (c *APIClient) GetCloudAccessRolesForOrgAccount(ctx context.Context, orgAccountId uint, filters *models.CloudAccessRoleFilters, limit *int, offset *int) ([]*models.CloudAccessRole, error) {
	var cloudAccessRoles []*models.CloudAccessRole
	tx := c.db.Scopes(applyPagination(limit, offset), cloudAccessRoleFilters(filters), cloudAccessRoleIDOrdering)
	// subquery to get all direct & upstream org units for the org account
	subQuery := tx.Debug().Select("unnest(array_append(string_to_array(trim(leading ':' from hierarchy), ':')::bigint[], org_unit_id))").Table("org_accounts_org_units").Joins("JOIN org_units direct on org_accounts_org_units.org_unit_id = direct.id").Where("org_account_id = ?", orgAccountId)
	// find all cloud access roles associated with the org units found in the subquery
	err := tx.Debug().Model(&models.CloudAccessRole{}).Select("*").Where("org_unit_id IN (?)", subQuery).Find(&cloudAccessRoles).Error
	if err != nil {
		return nil, err
	}
	return cloudAccessRoles, nil
}
