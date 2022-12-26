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

type DeactivateModuleAccountAssociationOutput struct {
}

func (t *TaskHandler) DeactivateModuleAccountAssociation(ctx context.Context, input DeactivateModuleAccountAssociationInput) (*DeactivateModuleAccountAssociationOutput, error) {
	// do stuff
	return &DeactivateModuleAccountAssociationOutput{}, nil
}
