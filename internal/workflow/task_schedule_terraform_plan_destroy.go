package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformPlanDestroy Task = "ScheduleTerraformPlanDestroy"
)

type ScheduleTerraformPlanDestroyInput struct {
	ModuleAccountAssociation models.ModuleAccountAssociation
}

type ScheduleTerraformPlanDestroyOutput struct {
}

func (t *TaskHandler) ScheduleTerraformPlanDestroy(ctx context.Context, input ScheduleTerraformPlanDestroyInput) (*ScheduleTerraformPlanDestroyOutput, error) {
	// do stuff
	return &ScheduleTerraformPlanDestroyOutput{}, nil
}
