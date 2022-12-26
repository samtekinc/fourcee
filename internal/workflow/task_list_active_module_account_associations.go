package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskListActiveModuleAccountAssociations Task = "ListActiveModuleAccountAssociations"
)

type ListActiveModuleAccountAssociationsInput struct {
	ModulePropagationId                 string
	ModulePropagationExecutionRequestId string
}

type ListActiveModuleAccountAssociationsOutput struct {
	ModuleAccountAssociations []models.ModuleAccountAssociation
}

func (t *TaskHandler) ListActiveModuleAccountAssociations(ctx context.Context, input ListActiveModuleAccountAssociationsInput) (*ListActiveModuleAccountAssociationsOutput, error) {
	moduleAccountAssociations := []models.ModuleAccountAssociation{}
	nextCursor := ""

	for {
		moduleAccountAssociationsPage, err := t.apiClient.GetModuleAccountAssociationsByModulePropagationId(ctx, input.ModulePropagationId, 100, nextCursor)
		if err != nil {
			return nil, err
		}
		for _, moduleAccountAssociation := range moduleAccountAssociationsPage.Items {
			if moduleAccountAssociation.Status == models.ModuleAccountAssociationStatusActive {
				moduleAccountAssociations = append(moduleAccountAssociations, moduleAccountAssociation)
			}
		}
		if moduleAccountAssociationsPage.NextCursor == "" {
			break
		}
		nextCursor = moduleAccountAssociationsPage.NextCursor
	}

	return &ListActiveModuleAccountAssociationsOutput{
		ModuleAccountAssociations: moduleAccountAssociations,
	}, nil
}
