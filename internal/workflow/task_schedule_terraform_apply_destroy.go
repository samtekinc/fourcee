package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformApplyDestroy Task = "ScheduleTerraformApplyDestroy"
)

type ScheduleTerraformApplyDestroyInput struct {
	ModuleAccountAssociation models.ModuleAccountAssociation
}

type ScheduleTerraformApplyDestroyOutput struct {
}

func (t *TaskHandler) ScheduleTerraformApplyDestroy(ctx context.Context, input ScheduleTerraformApplyDestroyInput) (*ScheduleTerraformApplyDestroyOutput, error) {
	// do stuff
	return &ScheduleTerraformApplyDestroyOutput{}, nil
}
