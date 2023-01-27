package client

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyApplyExecutionRequestFilters(tx *gorm.DB, filters *models.ApplyExecutionRequestFilters) *gorm.DB {
	if filters != nil {
		if filters.StartedBefore != nil {
			tx = tx.Where("started_at < ?", *filters.StartedBefore)
		}
		if filters.StartedAfter != nil {
			tx = tx.Where("started_at > ?", *filters.StartedAfter)
		}
		if filters.CompletedBefore != nil {
			tx = tx.Where("completed_at < ?", *filters.CompletedBefore)
		}
		if filters.CompletedAfter != nil {
			tx = tx.Where("completed_at > ?", *filters.CompletedAfter)
		}
		if filters.Status != nil {
			tx = tx.Where("status = ?", *filters.Status)
		}
		if filters.Destroy != nil {
			tx = tx.Where("destroy = ?", *filters.Destroy)
		}
	}
	return tx
}

func applyApplyExecutionRequestPreloads(tx *gorm.DB) *gorm.DB {
	return tx
}

func (c *APIClient) GetApplyExecutionRequestsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var applyExecutionRequests []*models.ApplyExecutionRequest
	tx := applyApplyExecutionRequestPreloads(c.db)
	err := tx.Find(&applyExecutionRequests, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: applyExecutionRequests[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetApplyExecutionRequest(ctx context.Context, id uint) (*models.ApplyExecutionRequest, error) {
	var applyExecutionRequest models.ApplyExecutionRequest
	tx := applyApplyExecutionRequestPreloads(c.db)
	err := tx.First(&applyExecutionRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &applyExecutionRequest, nil
}

func (c *APIClient) GetApplyExecutionRequestBatched(ctx context.Context, id uint) (*models.ApplyExecutionRequest, error) {
	thunk := c.applyExecutionRequestsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ApplyExecutionRequest), nil
}

func (c *APIClient) GetApplyExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId uint) (*models.ApplyExecutionRequest, error) {
	var applyExecutionRequest *models.ApplyExecutionRequest
	tx := applyApplyExecutionRequestPreloads(c.db)
	err := tx.Model(&models.TerraformExecutionRequest{Model: gorm.Model{ID: terraformExecutionRequestId}}).Association("ApplyExecutionRequestAssociation").Find(&applyExecutionRequest)
	if err != nil {
		return nil, err
	}
	return applyExecutionRequest, nil
}

func (c *APIClient) GetApplyExecutionRequests(ctx context.Context, filters *models.ApplyExecutionRequestFilters, limit *int, offset *int) ([]*models.ApplyExecutionRequest, error) {
	var applyExecutionRequests []*models.ApplyExecutionRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyApplyExecutionRequestFilters(tx, filters)
	tx = applyApplyExecutionRequestPreloads(tx)
	err := tx.Find(&applyExecutionRequests).Error
	if err != nil {
		return nil, err
	}
	return applyExecutionRequests, nil
}

func (c *APIClient) CreateApplyExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestID uint, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequest := models.ApplyExecutionRequest{
		ModuleAssignmentID:          input.ModuleAssignmentID,
		TerraformVersion:            input.TerraformVersion,
		CallbackTaskToken:           input.CallbackTaskToken,
		TerraformConfiguration:      input.TerraformConfiguration,
		TerraformPlan:               input.TerraformPlan,
		TerraformExecutionRequestID: terraformExecutionRequestID,
		AdditionalArguments:         input.AdditionalArguments,
		Status:                      models.RequestStatusPending,
	}
	err := c.db.Model(&models.TerraformExecutionRequest{Model: gorm.Model{ID: terraformExecutionRequestID}}).Association("ApplyExecutionRequestAssociation").Append(&applyExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}

func (c *APIClient) DeleteApplyExecutionRequest(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ApplyExecutionRequest{}, id).Error
}

func (c *APIClient) UpdateApplyExecutionRequest(ctx context.Context, id uint, update *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequest := models.ApplyExecutionRequest{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.ApplyExecutionRequest{}

	if update.InitOutput != nil {
		updates.InitOutput = update.InitOutput
	}
	if update.ApplyOutput != nil {
		updates.ApplyOutput = update.ApplyOutput
	}
	if update.Status != nil {
		updates.Status = *update.Status
	}
	if update.StartedAt != nil {
		updates.StartedAt = update.StartedAt
	}
	if update.CompletedAt != nil {
		updates.CompletedAt = update.CompletedAt
	}

	err := c.db.Model(&applyExecutionRequest).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}
