package client

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/temporal/constants"
	"github.com/sheacloud/tfom/internal/temporal/workflows"
	"github.com/sheacloud/tfom/pkg/models"
	"go.temporal.io/sdk/client"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyModulePropagationDriftCheckRequestFilters(tx *gorm.DB, filters *models.ModulePropagationDriftCheckRequestFilters) *gorm.DB {
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
		if filters.SyncStatus != nil {
			tx = tx.Where("sync_status = ?", *filters.SyncStatus)
		}
	}
	return tx
}

func applyModulePropagationDriftCheckRequestPreloads(tx *gorm.DB) *gorm.DB {
	return tx
}

func (c *APIClient) GetModulePropagationDriftCheckRequestsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var modulePropagationDriftCheckRequests []*models.ModulePropagationDriftCheckRequest
	tx := applyModulePropagationDriftCheckRequestPreloads(c.db)
	err := tx.Find(&modulePropagationDriftCheckRequests, keys.Keys()).Error
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

	response := make([]*dataloader.Result, len(modulePropagationDriftCheckRequests))
	for i := range modulePropagationDriftCheckRequests {
		index := keyToIndex[idToString(modulePropagationDriftCheckRequests[i].ID)]
		response[index] = &dataloader.Result{Data: modulePropagationDriftCheckRequests[i], Error: nil}
	}
	return response
}

func (c *APIClient) GetModulePropagationDriftCheckRequest(ctx context.Context, id uint) (*models.ModulePropagationDriftCheckRequest, error) {
	var modulePropagationDriftCheckRequest models.ModulePropagationDriftCheckRequest
	tx := applyModulePropagationDriftCheckRequestPreloads(c.db)
	err := tx.First(&modulePropagationDriftCheckRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &modulePropagationDriftCheckRequest, nil
}

func (c *APIClient) GetModulePropagationDriftCheckRequestBatched(ctx context.Context, id uint) (*models.ModulePropagationDriftCheckRequest, error) {
	thunk := c.modulePropagationDriftCheckRequestsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagationDriftCheckRequest), nil
}

func (c *APIClient) GetModulePropagationDriftCheckRequests(ctx context.Context, filters *models.ModulePropagationDriftCheckRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationDriftCheckRequest, error) {
	var modulePropagationDriftCheckRequests []*models.ModulePropagationDriftCheckRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationDriftCheckRequestFilters(tx, filters)
	tx = applyModulePropagationDriftCheckRequestPreloads(tx)
	tx = tx.Order("created_at DESC")
	err := tx.Find(&modulePropagationDriftCheckRequests).Error
	if err != nil {
		return nil, err
	}
	return modulePropagationDriftCheckRequests, nil
}

func (c *APIClient) GetModulePropagationDriftCheckRequestsForModulePropagation(ctx context.Context, modulePropagationId uint, filters *models.ModulePropagationDriftCheckRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationDriftCheckRequest, error) {
	var modulePropagationDriftCheckRequests []*models.ModulePropagationDriftCheckRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationDriftCheckRequestFilters(tx, filters)
	tx = applyModulePropagationDriftCheckRequestPreloads(tx)
	tx = tx.Order("created_at DESC")
	err := tx.Model(&models.ModulePropagation{Model: gorm.Model{ID: modulePropagationId}}).Association("ModulePropagationDriftCheckRequestsAssociation").Find(&modulePropagationDriftCheckRequests)
	if err != nil {
		return nil, err
	}
	return modulePropagationDriftCheckRequests, nil
}

func (c *APIClient) CreateModulePropagationDriftCheckRequest(ctx context.Context, input *models.NewModulePropagationDriftCheckRequest) (*models.ModulePropagationDriftCheckRequest, error) {
	modulePropagationDriftCheckRequest := models.ModulePropagationDriftCheckRequest{
		ModulePropagationID: input.ModulePropagationID,
		Status:              models.RequestStatusPending,
		SyncStatus:          models.TerraformDriftCheckStatusPending,
	}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&modulePropagationDriftCheckRequest).Error
		if err != nil {
			return err
		}

		// start the temporal workflow
		_, err = c.temporalClient.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{TaskQueue: constants.TFOMTaskQueue}, workflows.ModulePropagationDriftCheckWorkflow, &modulePropagationDriftCheckRequest)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &modulePropagationDriftCheckRequest, nil
}

func (c *APIClient) DeleteModulePropagationDriftCheckRequest(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ModulePropagationDriftCheckRequest{}, id).Error
}

func (c *APIClient) UpdateModulePropagationDriftCheckRequest(ctx context.Context, id uint, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error) {
	modulePropagationDriftCheckRequest := models.ModulePropagationDriftCheckRequest{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.ModulePropagationDriftCheckRequest{}

	if update.Status != nil {
		updates.Status = *update.Status
	}
	if update.SyncStatus != nil {
		updates.SyncStatus = *update.SyncStatus
	}
	if update.StartedAt != nil {
		updates.StartedAt = update.StartedAt
	}
	if update.CompletedAt != nil {
		updates.CompletedAt = update.CompletedAt
	}

	err := c.db.Model(&modulePropagationDriftCheckRequest).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &modulePropagationDriftCheckRequest, nil
}
