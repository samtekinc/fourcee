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

func (r *Activities) TerraformApply(ctx context.Context, applyExecutionRequest *models.ApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	newStatus := models.RequestStatusRunning
	startedTime := time.Now().UTC()
	_, err := r.apiClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequest.ID, &models.ApplyExecutionRequestUpdate{
		Status:    &newStatus,
		StartedAt: &startedTime,
	})

	if err != nil {
		return nil, err
	}

	workingDirectory, err := terraform.NewWorkingDirectory(filepath.Join(r.config.WorkingDirectory, "applies", strconv.FormatUint(uint64(applyExecutionRequest.ID), 10)))
	if err != nil {
		return nil, err
	}
	defer workingDirectory.DeleteDirectory()

	installationDirectory, err := terraform.NewTerraformInstallationDirectory(filepath.Join(r.config.WorkingDirectory, "executables"))
	if err != nil {
		return nil, fmt.Errorf("unable to create terraform installation directory, %w", err)
	}

	// install terraform
	executable, err := installationDirectory.InstallTerraform(applyExecutionRequest.TerraformVersion, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return nil, err
	}

	// add the TF config to the directory
	err = workingDirectory.AddFile(applyExecutionRequest.TerraformConfiguration, "main.tf")
	if err != nil {
		return nil, err
	}

	initOutput := bytes.NewBuffer([]byte{})
	err = executable.TerraformInit(workingDirectory, initOutput)
	// update the apply execution request with the init output, even if it failed
	// if it did fail, also update the status to failed and set the completed time
	update := &models.ApplyExecutionRequestUpdate{
		InitOutput: initOutput.Bytes(),
	}
	if err != nil {
		newStatus = models.RequestStatusFailed
		completedTime := time.Now().UTC()
		update.Status = &newStatus
		update.CompletedAt = &completedTime
	}
	_, updateErr := r.apiClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequest.ID, update)
	if err != nil {
		return nil, err
	}
	if updateErr != nil {
		return nil, updateErr
	}

	applyOutput := bytes.NewBuffer([]byte{})

	additionalArguments := []string{}
	if applyExecutionRequest.AdditionalArguments != nil {
		additionalArguments = strings.Split(*applyExecutionRequest.AdditionalArguments, " ")
	}

	err = executable.TerraformApply(workingDirectory, applyExecutionRequest.TerraformPlan, additionalArguments, applyOutput)
	if err != nil {
		newStatus = models.RequestStatusFailed
	} else {
		newStatus = models.RequestStatusSucceeded
	}
	completedTime := time.Now().UTC()
	_, updateErr = r.apiClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequest.ID, &models.ApplyExecutionRequestUpdate{
		Status:      &newStatus,
		CompletedAt: &completedTime,
		ApplyOutput: applyOutput.Bytes(),
	})
	if err != nil {
		return nil, err
	}

	if updateErr != nil {
		return nil, updateErr
	}

	return applyExecutionRequest, nil
}
