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

func applyTerraformDriftCheckRequestFilters(tx *gorm.DB, filters *models.TerraformDriftCheckRequestFilters) *gorm.DB {
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
		if filters.SyncStatus != nil {
			tx = tx.Where("sync_status = ?", *filters.SyncStatus)
		}
	}
	return tx
}

func applyTerraformDriftCheckRequestPreloads(tx *gorm.DB) *gorm.DB {
	return tx
}

func (c *APIClient) GetTerraformDriftCheckRequestsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var terraformDriftCheckRequests []*models.TerraformDriftCheckRequest
	tx := applyTerraformDriftCheckRequestPreloads(c.db)
	err := tx.Find(&terraformDriftCheckRequests, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: terraformDriftCheckRequests[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetTerraformDriftCheckRequest(ctx context.Context, id uint) (*models.TerraformDriftCheckRequest, error) {
	var terraformDriftCheckRequest models.TerraformDriftCheckRequest
	tx := applyTerraformDriftCheckRequestPreloads(c.db)
	err := tx.First(&terraformDriftCheckRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &terraformDriftCheckRequest, nil
}

func (c *APIClient) GetTerraformDriftCheckRequestBatched(ctx context.Context, id uint) (*models.TerraformDriftCheckRequest, error) {
	thunk := c.terraformDriftCheckRequestsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.TerraformDriftCheckRequest), nil
}

func (c *APIClient) GetTerraformDriftCheckRequests(ctx context.Context, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error) {
	var terraformDriftCheckRequests []*models.TerraformDriftCheckRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyTerraformDriftCheckRequestFilters(tx, filters)
	tx = applyTerraformDriftCheckRequestPreloads(tx)
	tx = tx.Order("created_at DESC")
	err := tx.Find(&terraformDriftCheckRequests).Error
	if err != nil {
		return nil, err
	}
	return terraformDriftCheckRequests, nil
}

func (c *APIClient) GetTerraformDriftCheckRequestsForModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationDriftCheckRequestID uint, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error) {
	var terraformDriftCheckRequests []*models.TerraformDriftCheckRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyTerraformDriftCheckRequestFilters(tx, filters)
	tx = applyTerraformDriftCheckRequestPreloads(tx)
	tx = tx.Order("created_at DESC")
	err := tx.Model(&models.ModulePropagationDriftCheckRequest{Model: gorm.Model{ID: modulePropagationDriftCheckRequestID}}).Association("TerraformDriftCheckRequestsAssociation").Find(&terraformDriftCheckRequests)
	if err != nil {
		return nil, err
	}
	return terraformDriftCheckRequests, nil
}

func (c *APIClient) GetTerraformDriftCheckRequestsForModuleAssignment(ctx context.Context, moduleAssignmentID uint, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error) {
	var terraformDriftCheckRequests []*models.TerraformDriftCheckRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyTerraformDriftCheckRequestFilters(tx, filters)
	tx = applyTerraformDriftCheckRequestPreloads(tx)
	tx = tx.Order("created_at DESC")
	err := tx.Model(&models.ModuleAssignment{Model: gorm.Model{ID: moduleAssignmentID}}).Association("TerraformDriftCheckRequestsAssociation").Find(&terraformDriftCheckRequests)
	if err != nil {
		return nil, err
	}
	return terraformDriftCheckRequests, nil
}

func (c *APIClient) CreateTerraformDriftCheckRequest(ctx context.Context, input *models.NewTerraformDriftCheckRequest) (*models.TerraformDriftCheckRequest, error) {
	terraformDriftCheckRequest := models.TerraformDriftCheckRequest{
		ModuleAssignmentID:                   input.ModuleAssignmentID,
		ModulePropagationID:                  input.ModulePropagationID,
		Destroy:                              input.Destroy,
		CallbackTaskToken:                    input.CallbackTaskToken,
		Status:                               models.RequestStatusPending,
		SyncStatus:                           models.TerraformDriftCheckStatusPending,
		ModulePropagationDriftCheckRequestID: input.ModulePropagationDriftCheckRequestID,
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&terraformDriftCheckRequest).Error
		if err != nil {
			return err
		}

		// start the temporal workflow
		_, err = c.temporalClient.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{TaskQueue: constants.TFOMTaskQueue}, workflows.TerraformDriftCheckWorkflow, &terraformDriftCheckRequest)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &terraformDriftCheckRequest, nil
}

func (c *APIClient) DeleteTerraformDriftCheckRequest(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.TerraformDriftCheckRequest{}, id).Error
}

func (c *APIClient) UpdateTerraformDriftCheckRequest(ctx context.Context, id uint, update *models.TerraformDriftCheckRequestUpdate) (*models.TerraformDriftCheckRequest, error) {
	terraformDriftCheckRequest := models.TerraformDriftCheckRequest{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.TerraformDriftCheckRequest{}

	if update.Status != nil {
		updates.Status = *update.Status
	}
	if update.StartedAt != nil {
		updates.StartedAt = update.StartedAt
	}
	if update.CompletedAt != nil {
		updates.CompletedAt = update.CompletedAt
	}
	if update.SyncStatus != nil {
		updates.SyncStatus = *update.SyncStatus
	}

	err := c.db.Model(&terraformDriftCheckRequest).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &terraformDriftCheckRequest, nil
}
