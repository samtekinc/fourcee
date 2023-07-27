package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) ActivateModuleAssignment(ctx context.Context, moduleAssignmentID uint) error {
	newStatus := models.ModuleAssignmentStatusActive
	_, err := r.apiClient.UpdateModuleAssignment(ctx, moduleAssignmentID, &models.ModuleAssignmentUpdate{
		Status: &newStatus,
	})
	return err
}
