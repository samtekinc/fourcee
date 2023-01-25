package api

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyPlanExecutionRequestFilters(tx *gorm.DB, filters *models.PlanExecutionRequestFilters) *gorm.DB {
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

func applyPlanExecutionRequestPreloads(tx *gorm.DB) *gorm.DB {
	return tx
}

func (c *APIClient) GetPlanExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var planExecutionRequests []*models.PlanExecutionRequest
	tx := applyPlanExecutionRequestPreloads(c.db)
	err := tx.Find(&planExecutionRequests, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: planExecutionRequests[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetPlanExecutionRequest(ctx context.Context, id uint) (*models.PlanExecutionRequest, error) {
	var planExecutionRequest models.PlanExecutionRequest
	tx := applyPlanExecutionRequestPreloads(c.db)
	err := tx.First(&planExecutionRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &planExecutionRequest, nil
}

func (c *APIClient) GetPlanExecutionRequestBatched(ctx context.Context, id uint) (*models.PlanExecutionRequest, error) {
	thunk := c.planExecutionRequestsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.PlanExecutionRequest), nil
}

func (c *APIClient) GetPlanExecutionRequests(ctx context.Context, filters *models.PlanExecutionRequestFilters, limit *int, offset *int) ([]*models.PlanExecutionRequest, error) {
	var planExecutionRequests []*models.PlanExecutionRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyPlanExecutionRequestFilters(tx, filters)
	tx = applyPlanExecutionRequestPreloads(tx)
	err := tx.Find(&planExecutionRequests).Error
	if err != nil {
		return nil, err
	}
	return planExecutionRequests, nil
}

func (c *APIClient) CreatePlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	planExecutionRequest := models.PlanExecutionRequest{
		ModuleAssignmentID:           input.ModuleAssignmentID,
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		TerraformDriftCheckRequestID: input.TerraformDriftCheckRequestID,
		TerraformExecutionRequestID:  input.TerraformExecutionRequestID,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.RequestStatusPending,
	}
	err := c.db.Create(&planExecutionRequest).Error
	if err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}

func (c *APIClient) DeletePlanExecutionRequest(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.PlanExecutionRequest{}, id).Error
}

func (c *APIClient) UpdatePlanExecutionRequest(ctx context.Context, id uint, update *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error) {
	planExecutionRequest := models.PlanExecutionRequest{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.PlanExecutionRequest{}

	if update.InitOutput != nil {
		updates.InitOutput = update.InitOutput
	}
	if update.PlanOutput != nil {
		updates.PlanOutput = update.PlanOutput
	}
	if update.PlanFile != nil {
		updates.PlanFile = update.PlanFile
	}
	if update.PlanJSON != nil {
		updates.PlanJSON = update.PlanJSON
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

	err := c.db.Model(&planExecutionRequest).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}
