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

func awsIamPolicyFilters(filters *models.AwsIamPolicyFilters) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if filters != nil {
			if filters.NameContains != nil {
				tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
			}
		}
		return tx
	}
}

func awsIamPolicyIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id")
}

func (c *APIClient) GetAwsIamPoliciesByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var awsIamPolicies []*models.AwsIamPolicy
	tx := c.db.Scopes()
	err := tx.Find(&awsIamPolicies, keys.Keys()).Error
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

	response := make([]*dataloader.Result, len(awsIamPolicies))
	for i := range awsIamPolicies {
		index := keyToIndex[idToString(awsIamPolicies[i].ID)]
		response[index] = &dataloader.Result{Data: awsIamPolicies[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("AWS IAM Policy %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetAwsIamPolicy(ctx context.Context, id uint) (*models.AwsIamPolicy, error) {
	var awsIamPolicy models.AwsIamPolicy
	tx := c.db.Scopes()
	err := tx.First(&awsIamPolicy, id).Error
	if err != nil {
		return nil, err
	}
	return &awsIamPolicy, nil
}

func (c *APIClient) GetAwsIamPolicyBatched(ctx context.Context, id uint) (*models.AwsIamPolicy, error) {
	thunk := c.awsIamPolicyLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.AwsIamPolicy), nil
}

func (c *APIClient) GetAwsIamPolicies(ctx context.Context, filters *models.AwsIamPolicyFilters, limit *int, offset *int) ([]*models.AwsIamPolicy, error) {
	var awsIamPolicies []*models.AwsIamPolicy
	tx := c.db.Scopes(applyPagination(limit, offset), awsIamPolicyFilters(filters), awsIamPolicyIDOrdering)
	err := tx.Find(&awsIamPolicies).Error
	if err != nil {
		return nil, err
	}
	return awsIamPolicies, nil
}

func (c *APIClient) CreateAwsIamPolicy(ctx context.Context, input *models.NewAwsIamPolicy) (*models.AwsIamPolicy, error) {
	awsIamPolicy := models.AwsIamPolicy{
		Name:           input.Name,
		PolicyDocument: input.PolicyDocument,
	}
	err := c.db.Create(&awsIamPolicy).Error
	if err != nil {
		return nil, err
	}

	return &awsIamPolicy, nil
}

func (c *APIClient) DeleteAwsIamPolicy(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.AwsIamPolicy{}, id).Error
}

func (c *APIClient) UpdateAwsIamPolicy(ctx context.Context, id uint, update *models.AwsIamPolicyUpdate) (*models.AwsIamPolicy, error) {
	awsIamPolicy := models.AwsIamPolicy{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.AwsIamPolicy{}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		if update.Name != nil {
			updates.Name = *update.Name
		}
		if update.PolicyDocument != nil {
			updates.PolicyDocument = *update.PolicyDocument
		}
		err := tx.Model(&awsIamPolicy).Updates(updates).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &awsIamPolicy, nil
}

func (c *APIClient) GetAwsIamPoliciesForCloudAccessRole(ctx context.Context, cloudAccessRoleId uint, filters *models.AwsIamPolicyFilters, limit *int, offset *int) ([]*models.AwsIamPolicy, error) {
	var awsIamPolicies []*models.AwsIamPolicy
	tx := c.db.Scopes(applyPagination(limit, offset), awsIamPolicyFilters(filters), awsIamPolicyIDOrdering)
	err := tx.Model(&models.CloudAccessRole{Model: gorm.Model{ID: cloudAccessRoleId}}).Association("AwsIamPoliciesAssociation").Find(&awsIamPolicies)
	if err != nil {
		return nil, err
	}
	return awsIamPolicies, nil
}
