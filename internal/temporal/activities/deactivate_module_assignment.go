package activities

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

func (r *Activities) DeactivateModuleAssignment(ctx context.Context, moduleAssignmentID uint) error {
	newStatus := models.ModuleAssignmentStatusInactive
	_, err := r.apiClient.UpdateModuleAssignment(ctx, moduleAssignmentID, &models.ModuleAssignmentUpdate{
		Status: &newStatus,
	})
	return err
}
