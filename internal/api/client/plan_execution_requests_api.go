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

func planExecutionRequestFilters(filters *models.PlanExecutionRequestFilters) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
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
}

func planExecutionRequestIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id")
}

func (c *APIClient) GetPlanExecutionRequestsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var planExecutionRequests []*models.PlanExecutionRequest
	tx := c.db.Scopes()
	err := tx.Find(&planExecutionRequests, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	var keyToIndex = map[string]int{}
	for i := range keys {
		keyToIndex[keys[i].String()] = i
	}

	response := make([]*dataloader.Result, len(planExecutionRequests))
	for i := range planExecutionRequests {
		index := keyToIndex[idToString(planExecutionRequests[i].ID)]
		response[index] = &dataloader.Result{Data: planExecutionRequests[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("Plan Execution Request %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetPlanExecutionRequest(ctx context.Context, id uint) (*models.PlanExecutionRequest, error) {
	var planExecutionRequest models.PlanExecutionRequest
	tx := c.db.Scopes()
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

func (c *APIClient) GetPlanExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId uint) (*models.PlanExecutionRequest, error) {
	var planExecutionRequest *models.PlanExecutionRequest
	tx := c.db.Scopes()
	err := tx.Model(&models.TerraformExecutionRequest{Model: gorm.Model{ID: terraformExecutionRequestId}}).Association("PlanExecutionRequestAssociation").Find(&planExecutionRequest)
	if err != nil {
		return nil, err
	}
	return planExecutionRequest, nil
}

func (c *APIClient) GetPlanExecutionRequestForTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestId uint) (*models.PlanExecutionRequest, error) {
	var planExecutionRequest *models.PlanExecutionRequest
	tx := c.db.Scopes()
	err := tx.Model(&models.TerraformDriftCheckRequest{Model: gorm.Model{ID: terraformDriftCheckRequestId}}).Association("PlanExecutionRequestAssociation").Find(&planExecutionRequest)
	if err != nil {
		return nil, err
	}
	return planExecutionRequest, nil
}

func (c *APIClient) GetPlanExecutionRequests(ctx context.Context, filters *models.PlanExecutionRequestFilters, limit *int, offset *int) ([]*models.PlanExecutionRequest, error) {
	var planExecutionRequests []*models.PlanExecutionRequest
	tx := c.db.Scopes(applyPagination(limit, offset), planExecutionRequestFilters(filters), planExecutionRequestIDOrdering)
	err := tx.Find(&planExecutionRequests).Error
	if err != nil {
		return nil, err
	}
	return planExecutionRequests, nil
}

func (c *APIClient) CreatePlanExecutionRequestForTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestID uint, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	planExecutionRequest := models.PlanExecutionRequest{
		ModuleAssignmentID:           input.ModuleAssignmentID,
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		TerraformConfiguration:       input.TerraformConfiguration,
		TerraformDriftCheckRequestID: nil,
		TerraformExecutionRequestID:  &terraformExecutionRequestID,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.RequestStatusPending,
	}
	err := c.db.Model(&models.TerraformExecutionRequest{Model: gorm.Model{ID: terraformExecutionRequestID}}).Association("PlanExecutionRequestAssociation").Append(&planExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}

func (c *APIClient) CreatePlanExecutionRequestForTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestID uint, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	planExecutionRequest := models.PlanExecutionRequest{
		ModuleAssignmentID:           input.ModuleAssignmentID,
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		TerraformConfiguration:       input.TerraformConfiguration,
		TerraformDriftCheckRequestID: &terraformDriftCheckRequestID,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.RequestStatusPending,
	}
	err := c.db.Model(&models.TerraformDriftCheckRequest{Model: gorm.Model{ID: terraformDriftCheckRequestID}}).Association("PlanExecutionRequestAssociation").Append(&planExecutionRequest)
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
