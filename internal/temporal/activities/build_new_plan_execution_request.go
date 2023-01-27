package activities

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
)

func (r *Activities) BuildNewPlanExecutionRequest(ctx context.Context, moduleAssignmentID uint, terraformDriftCheckRequestID *uint, terraformExecutionRequestID *uint, destroy bool) (*models.NewPlanExecutionRequest, error) {
	moduleAssignment, err := r.apiClient.GetModuleAssignment(ctx, moduleAssignmentID)
	if err != nil {
		return nil, err
	}

	var modulePropagation *models.ModulePropagation
	if moduleAssignment.ModulePropagationID != nil {
		modulePropagation, err = r.apiClient.GetModulePropagation(ctx, *moduleAssignment.ModulePropagationID)
		if err != nil {
			return nil, err
		}
	}

	moduleVersion, err := r.apiClient.GetModuleVersion(ctx, moduleAssignment.ModuleVersionID)
	if err != nil {
		return nil, err
	}

	orgAccount, err := r.apiClient.GetOrgAccount(ctx, moduleAssignment.OrgAccountID)
	if err != nil {
		return nil, err
	}

	terraformConfig, err := terraform.GetTerraformConfiguration(&terraform.TerraformConfigurationInput{
		ModuleAssignment:  moduleAssignment,
		ModulePropagation: modulePropagation,
		ModuleVersion:     moduleVersion,
		OrgAccount:        orgAccount,
		LockTableName:     r.config.Prefix + "-terraform-lock",
	})
	if err != nil {
		return nil, err
	}

	var additionalArguments *string
	if destroy {
		additionalArguments = aws.String("-destroy")
	}

	planExecutionRequest := &models.NewPlanExecutionRequest{
		ModuleAssignmentID:           moduleAssignmentID,
		CallbackTaskToken:            "",
		TerraformDriftCheckRequestID: terraformDriftCheckRequestID,
		TerraformExecutionRequestID:  terraformExecutionRequestID,
		TerraformConfiguration:       terraformConfig,
		TerraformVersion:             moduleVersion.TerraformVersion,
		AdditionalArguments:          additionalArguments,
	}

	return planExecutionRequest, nil
}
