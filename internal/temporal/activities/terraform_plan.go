package activities

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
)

func (r *Activities) TerraformPlan(ctx context.Context, planExecutionRequest *models.PlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	newStatus := models.RequestStatusRunning
	startedTime := time.Now().UTC()
	_, err := r.apiClient.UpdatePlanExecutionRequest(ctx, planExecutionRequest.ID, &models.PlanExecutionRequestUpdate{
		Status:    &newStatus,
		StartedAt: &startedTime,
	})

	if err != nil {
		return nil, err
	}

	workingDirectory, err := terraform.NewWorkingDirectory(filepath.Join(r.config.TfWorkingDirectory, strconv.FormatUint(uint64(planExecutionRequest.ID), 10)))
	if err != nil {
		return nil, err
	}
	defer workingDirectory.DeleteDirectory()

	installationDirectory, err := terraform.NewTerraformInstallationDirectory(r.config.TfInstallationDirectory)
	if err != nil {
		return nil, fmt.Errorf("unable to create terraform installation directory, %w", err)
	}

	// install terraform
	executable, err := installationDirectory.InstallTerraform(planExecutionRequest.TerraformVersion, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return nil, err
	}

	// add the TF config to the directory
	err = workingDirectory.AddFile(planExecutionRequest.TerraformConfiguration, "main.tf")
	if err != nil {
		return nil, err
	}

	initOutput := bytes.NewBuffer([]byte{})
	err = executable.TerraformInit(workingDirectory, initOutput)
	if err != nil {
		return nil, err
	}

	// plan terraform
	planOutput := bytes.NewBuffer([]byte{})
	if err != nil {
		return nil, err
	}
	planFileOutput := bytes.NewBuffer([]byte{})
	if err != nil {
		return nil, err
	}
	planJSONOutput := bytes.NewBuffer([]byte{})
	if err != nil {
		return nil, err
	}

	additionalArguments := []string{}
	if planExecutionRequest.AdditionalArguments != nil {
		additionalArguments = strings.Split(*planExecutionRequest.AdditionalArguments, " ")
	}
	err = executable.TerraformPlan(workingDirectory, additionalArguments, planOutput, planFileOutput, planJSONOutput)

	if err != nil {
		newStatus = models.RequestStatusFailed
	} else {
		newStatus = models.RequestStatusSucceeded
	}
	completedTime := time.Now().UTC()
	_, updateErr := r.apiClient.UpdatePlanExecutionRequest(ctx, planExecutionRequest.ID, &models.PlanExecutionRequestUpdate{
		Status:      &newStatus,
		CompletedAt: &completedTime,
		InitOutput:  initOutput.Bytes(),
		PlanOutput:  planOutput.Bytes(),
		PlanFile:    planFileOutput.Bytes(),
		PlanJSON:    planJSONOutput.Bytes(),
	})
	if err != nil {
		return nil, err
	}

	if updateErr != nil {
		return nil, updateErr
	}

	return planExecutionRequest, nil
}
