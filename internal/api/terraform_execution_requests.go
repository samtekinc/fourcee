package api

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyTerraformExecutionRequestFilters(tx *gorm.DB, filters *models.TerraformExecutionRequestFilters) *gorm.DB {
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

func applyTerraformExecutionRequestPreloads(tx *gorm.DB) *gorm.DB {
	return tx
}

func (c *APIClient) GetTerraformExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var terraformExecutionRequests []*models.TerraformExecutionRequest
	tx := applyTerraformExecutionRequestPreloads(c.db)
	err := tx.Find(&terraformExecutionRequests, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: terraformExecutionRequests[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetTerraformExecutionRequest(ctx context.Context, id uint) (*models.TerraformExecutionRequest, error) {
	var terraformExecutionRequest models.TerraformExecutionRequest
	tx := applyTerraformExecutionRequestPreloads(c.db)
	err := tx.First(&terraformExecutionRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &terraformExecutionRequest, nil
}

func (c *APIClient) GetTerraformExecutionRequestBatched(ctx context.Context, id uint) (*models.TerraformExecutionRequest, error) {
	thunk := c.terraformExecutionRequestsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.TerraformExecutionRequest), nil
}

func (c *APIClient) GetTerraformExecutionRequests(ctx context.Context, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error) {
	var terraformExecutionRequests []*models.TerraformExecutionRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyTerraformExecutionRequestFilters(tx, filters)
	tx = applyTerraformExecutionRequestPreloads(tx)
	err := tx.Find(&terraformExecutionRequests).Error
	if err != nil {
		return nil, err
	}
	return terraformExecutionRequests, nil
}

func (c *APIClient) GetTerraformExecutionRequestsForModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestID uint, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error) {
	var terraformExecutionRequests []*models.TerraformExecutionRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyTerraformExecutionRequestFilters(tx, filters)
	tx = applyTerraformExecutionRequestPreloads(tx)
	err := tx.Model(&models.ModulePropagationExecutionRequest{Model: gorm.Model{ID: modulePropagationExecutionRequestID}}).Association("TerraformExecutionRequestsAssociation").Find(&terraformExecutionRequests)
	if err != nil {
		return nil, err
	}
	return terraformExecutionRequests, nil
}

func (c *APIClient) GetTerraformExecutionRequestsForModuleAssignment(ctx context.Context, moduleAssignmentID uint, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error) {
	var terraformExecutionRequests []*models.TerraformExecutionRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyTerraformExecutionRequestFilters(tx, filters)
	tx = applyTerraformExecutionRequestPreloads(tx)
	err := tx.Model(&models.ModuleAssignment{Model: gorm.Model{ID: moduleAssignmentID}}).Association("TerraformExecutionRequestsAssociation").Find(&terraformExecutionRequests)
	if err != nil {
		return nil, err
	}
	return terraformExecutionRequests, nil
}

func (c *APIClient) CreateTerraformExecutionRequest(ctx context.Context, input *models.NewTerraformExecutionRequest) (*models.TerraformExecutionRequest, error) {
	terraformExecutionRequest := models.TerraformExecutionRequest{
		ModuleAssignmentID:                  input.ModuleAssignmentID,
		ModulePropagationID:                 input.ModulePropagationID,
		Destroy:                             input.Destroy,
		CallbackTaskToken:                   input.CallbackTaskToken,
		Status:                              models.RequestStatusPending,
		ModulePropagationExecutionRequestID: input.ModulePropagationExecutionRequestID,
	}
	err := c.db.Create(&terraformExecutionRequest).Error
	if err != nil {
		return nil, err
	}

	return &terraformExecutionRequest, nil
}

func (c *APIClient) DeleteTerraformExecutionRequest(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.TerraformExecutionRequest{}, id).Error
}

func (c *APIClient) UpdateTerraformExecutionRequest(ctx context.Context, id uint, update *models.TerraformExecutionRequestUpdate) (*models.TerraformExecutionRequest, error) {
	terraformExecutionRequest := models.TerraformExecutionRequest{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.TerraformExecutionRequest{}

	if update.Status != nil {
		updates.Status = *update.Status
	}
	if update.StartedAt != nil {
		updates.StartedAt = update.StartedAt
	}
	if update.CompletedAt != nil {
		updates.CompletedAt = update.CompletedAt
	}
	if update.PlanExecutionRequestID != nil {
		updates.PlanExecutionRequestID = update.PlanExecutionRequestID
	}
	if update.ApplyExecutionRequestID != nil {
		updates.ApplyExecutionRequestID = update.ApplyExecutionRequestID
	}

	err := c.db.Model(&terraformExecutionRequest).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &terraformExecutionRequest, nil
}
