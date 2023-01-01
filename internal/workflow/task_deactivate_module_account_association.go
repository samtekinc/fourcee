package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskDeactivateModuleAccountAssociation Task = "DeactivateModuleAccountAssociation"
)

type DeactivateModuleAccountAssociationInput struct {
	ModuleAccountAssociation models.ModuleAccountAssociation
}

type DeactivateModuleAccountAssociationOutput struct{}

func (t *TaskHandler) DeactivateModuleAccountAssociation(ctx context.Context, input DeactivateModuleAccountAssociationInput) (*DeactivateModuleAccountAssociationOutput, error) {

	newStatus := models.ModuleAccountAssociationStatusInactive
	_, err := t.apiClient.UpdateModuleAccountAssociation(ctx, input.ModuleAccountAssociation.ModulePropagationId, input.ModuleAccountAssociation.OrgAccountId, &models.ModuleAccountAssociationUpdate{
		Status: &newStatus,
	})
	if err != nil {
		return nil, err
	}

	return &DeactivateModuleAccountAssociationOutput{}, nil
}
