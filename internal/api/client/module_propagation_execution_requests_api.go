package client

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/samtekinc/fourcee/internal/helpers"
	"github.com/samtekinc/fourcee/internal/temporal/constants"
	"github.com/samtekinc/fourcee/internal/temporal/workflows"
	"github.com/samtekinc/fourcee/pkg/models"
	"go.temporal.io/sdk/client"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func modulePropagationExecutionRequestFilters(filters *models.ModulePropagationExecutionRequestFilters) func(tx *gorm.DB) *gorm.DB {
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
		}
		return tx
	}
}

func modulePropagationExecutionRequestIDOrdering(tx *gorm.DB) *gorm.DB {
	return tx.Order("id DESC")
}

func (c *APIClient) GetModulePropagationExecutionRequestsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var modulePropagationExecutionRequests []*models.ModulePropagationExecutionRequest
	tx := c.db.Scopes()
	err := tx.Find(&modulePropagationExecutionRequests, keys.Keys()).Error
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

	response := make([]*dataloader.Result, len(modulePropagationExecutionRequests))
	for i := range modulePropagationExecutionRequests {
		index := keyToIndex[idToString(modulePropagationExecutionRequests[i].ID)]
		response[index] = &dataloader.Result{Data: modulePropagationExecutionRequests[i], Error: nil}
	}

	for i, key := range keys {
		if response[i] == nil {
			response[i] = &dataloader.Result{Error: helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation Execution Request %s not found", key.String())}}
		}
	}

	return response
}

func (c *APIClient) GetModulePropagationExecutionRequest(ctx context.Context, id uint) (*models.ModulePropagationExecutionRequest, error) {
	var modulePropagationExecutionRequest models.ModulePropagationExecutionRequest
	tx := c.db.Scopes()
	err := tx.First(&modulePropagationExecutionRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &modulePropagationExecutionRequest, nil
}

func (c *APIClient) GetModulePropagationExecutionRequestBatched(ctx context.Context, id uint) (*models.ModulePropagationExecutionRequest, error) {
	thunk := c.modulePropagationExecutionRequestsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagationExecutionRequest), nil
}

func (c *APIClient) GetModulePropagationExecutionRequests(ctx context.Context, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error) {
	var modulePropagationExecutionRequests []*models.ModulePropagationExecutionRequest
	tx := c.db.Scopes(applyPagination(limit, offset), modulePropagationExecutionRequestFilters(filters), modulePropagationExecutionRequestIDOrdering)
	err := tx.Find(&modulePropagationExecutionRequests).Error
	if err != nil {
		return nil, err
	}
	return modulePropagationExecutionRequests, nil
}

func (c *APIClient) GetModulePropagationExecutionRequestsForModulePropagation(ctx context.Context, modulePropagationId uint, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error) {
	var modulePropagationExecutionRequests []*models.ModulePropagationExecutionRequest
	tx := c.db.Scopes(applyPagination(limit, offset), modulePropagationExecutionRequestFilters(filters), modulePropagationExecutionRequestIDOrdering)
	err := tx.Model(&models.ModulePropagation{Model: gorm.Model{ID: modulePropagationId}}).Association("ModulePropagationExecutionRequestsAssociation").Find(&modulePropagationExecutionRequests)
	if err != nil {
		return nil, err
	}
	return modulePropagationExecutionRequests, nil
}

func (c *APIClient) CreateModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error) {
	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{
		ModulePropagationID: input.ModulePropagationID,
		Status:              models.RequestStatusPending,
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&modulePropagationExecutionRequest).Error
		if err != nil {
			return err
		}

		// start the temporal workflow
		_, err = c.temporalClient.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{TaskQueue: constants.TFOMTaskQueue}, workflows.ModulePropagationExecutionWorkflow, &modulePropagationExecutionRequest)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}

func (c *APIClient) DeleteModulePropagationExecutionRequest(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ModulePropagationExecutionRequest{}, id).Error
}

func (c *APIClient) UpdateModulePropagationExecutionRequest(ctx context.Context, id uint, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error) {
	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.ModulePropagationExecutionRequest{}

	if update.Status != nil {
		updates.Status = *update.Status
	}
	if update.StartedAt != nil {
		updates.StartedAt = update.StartedAt
	}
	if update.CompletedAt != nil {
		updates.CompletedAt = update.CompletedAt
	}

	err := c.db.Model(&modulePropagationExecutionRequest).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}
