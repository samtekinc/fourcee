package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskListActiveModulePropagationAssignments Task = "ListActiveModulePropagationAssignments"
)

type ListActiveModulePropagationAssignmentsInput struct {
	ModulePropagationId                 string
	ModulePropagationExecutionRequestId string
}

type ListActiveModulePropagationAssignmentsOutput struct {
	ModuleAssignments []models.ModuleAssignment
}

func (t *TaskHandler) ListActiveModulePropagationAssignments(ctx context.Context, input ListActiveModulePropagationAssignmentsInput) (*ListActiveModulePropagationAssignmentsOutput, error) {
	moduleAssignments := []models.ModuleAssignment{}
	nextCursor := ""

	for {
		moduleAssignmentsPage, err := t.apiClient.GetModuleAssignmentsByModulePropagationId(ctx, input.ModulePropagationId, 100, nextCursor)
		if err != nil {
			return nil, err
		}
		for _, moduleAssignment := range moduleAssignmentsPage.Items {
			if moduleAssignment.Status == models.ModuleAssignmentStatusActive {
				moduleAssignments = append(moduleAssignments, moduleAssignment)
			}
		}
		if moduleAssignmentsPage.NextCursor == "" {
			break
		}
		nextCursor = moduleAssignmentsPage.NextCursor
	}

	return &ListActiveModulePropagationAssignmentsOutput{
		ModuleAssignments: moduleAssignments,
	}, nil
}
