package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) ListActiveModuleAssignmentsForPropagation(ctx context.Context, modulePropagationID uint) ([]*models.ModuleAssignment, error) {
	desiredStatus := models.ModuleAssignmentStatusActive
	return r.apiClient.GetModuleAssignmentsForModulePropagation(ctx, modulePropagationID, &models.ModuleAssignmentFilters{
		Status: &desiredStatus,
	}, nil, nil)
}
