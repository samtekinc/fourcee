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
		Status:     &newStatus,
		StartedAt:  &startedTime,
		InitOutput: []byte{}, // clear the init output (in case this is a retry)
		PlanOutput: []byte{}, // clear the plan output (in case this is a retry)
	})

	if err != nil {
		return nil, err
	}

	workingDirectory, err := terraform.NewWorkingDirectory(filepath.Join(r.config.WorkingDirectory, "plans", strconv.FormatUint(uint64(planExecutionRequest.ID), 10)))
	if err != nil {
		return nil, err
	}
	defer workingDirectory.DeleteDirectory()

	installationDirectory, err := terraform.NewTerraformInstallationDirectory(filepath.Join(r.config.WorkingDirectory, "executables"))
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
	// update the plan execution request with the init output, even if it failed
	// if it did fail, also update the status to failed and set the completed time
	update := &models.PlanExecutionRequestUpdate{
		InitOutput: initOutput.Bytes(),
	}
	if err != nil {
		newStatus = models.RequestStatusFailed
		completedTime := time.Now().UTC()
		update.Status = &newStatus
		update.CompletedAt = &completedTime
	}
	_, updateErr := r.apiClient.UpdatePlanExecutionRequest(ctx, planExecutionRequest.ID, update)
	if err != nil {
		return nil, err
	}
	if updateErr != nil {
		return nil, updateErr
	}

	// plan terraform
	planOutput := bytes.NewBuffer([]byte{})
	planFileOutput := bytes.NewBuffer([]byte{})
	planJSONOutput := bytes.NewBuffer([]byte{})

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
	_, updateErr = r.apiClient.UpdatePlanExecutionRequest(ctx, planExecutionRequest.ID, &models.PlanExecutionRequestUpdate{
		Status:      &newStatus,
		CompletedAt: &completedTime,
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
