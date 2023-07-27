package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) UpdateModulePropagationAssignments(ctx context.Context, modulePropagationID uint, moduleAssignments []*models.ModuleAssignment) ([]*models.ModuleAssignment, error) {
	// get the module propagation
	modulePropagation, err := r.apiClient.GetModulePropagation(ctx, modulePropagationID)
	if err != nil {
		return nil, err
	}

	for i, moduleAssignment := range moduleAssignments {
		update := models.ModuleAssignmentUpdate{}
		if moduleAssignment.Name != modulePropagation.Name {
			update.Name = &modulePropagation.Name
		}
		if moduleAssignment.Description != modulePropagation.Description {
			update.Description = &modulePropagation.Description
		}
		if moduleAssignment.ModuleVersionID != modulePropagation.ModuleVersionID {
			update.ModuleVersionID = &modulePropagation.ModuleVersionID
		}
		updatedModuleAssignment, err := r.apiClient.UpdateModuleAssignment(ctx, moduleAssignment.ID, &update)
		if err != nil {
			return nil, err
		}
		moduleAssignments[i] = updatedModuleAssignment
	}

	return moduleAssignments, nil
}
