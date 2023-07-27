package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) CreateModuleAssignments(ctx context.Context, modulePropagationID uint, orgAccounts []*models.OrgAccount) ([]*models.ModuleAssignment, error) {
	// get the module propagation
	modulePropagation, err := r.apiClient.GetModulePropagation(ctx, modulePropagationID)
	if err != nil {
		return nil, err
	}

	moduleAssignments := make([]*models.ModuleAssignment, len(orgAccounts))
	for i, orgAccount := range orgAccounts {
		// check if there is an existing, but inactive module assignment
		existingModuleAssignments, err := r.apiClient.GetModuleAssignmentsForModulePropagation(ctx, modulePropagationID, &models.ModuleAssignmentFilters{
			OrgAccountID: &orgAccount.ID,
		}, nil, nil)
		if err != nil {
			return nil, err
		}
		if len(existingModuleAssignments) > 0 {
			// there is an existing module assignment, activate it and update the name, description and module version
			existingModuleAssignment := existingModuleAssignments[0]
			update := models.ModuleAssignmentUpdate{}
			if existingModuleAssignment.Name != modulePropagation.Name {
				update.Name = &modulePropagation.Name
			}
			if existingModuleAssignment.Description != modulePropagation.Description {
				update.Description = &modulePropagation.Description
			}
			if existingModuleAssignment.ModuleVersionID != modulePropagation.ModuleVersionID {
				update.ModuleVersionID = &modulePropagation.ModuleVersionID
			}
			updatedModuleAssignment, err := r.apiClient.UpdateModuleAssignment(ctx, existingModuleAssignment.ID, &update)
			if err != nil {
				return nil, err
			}
			moduleAssignments[i] = updatedModuleAssignment
		} else {
			// no existing module assignment, create a new one
			newModuleAssignment := &models.NewModuleAssignment{
				ModuleVersionID:     modulePropagation.ModuleVersionID,
				ModuleGroupID:       modulePropagation.ModuleGroupID,
				OrgAccountID:        orgAccount.ID,
				Name:                modulePropagation.Name,
				Description:         modulePropagation.Description,
				ModulePropagationID: &modulePropagation.ID,
			}
			moduleAssignment, err := r.apiClient.CreateModuleAssignment(ctx, newModuleAssignment)
			if err != nil {
				return nil, err
			}
			moduleAssignments[i] = moduleAssignment
		}
	}

	return moduleAssignments, nil
}
