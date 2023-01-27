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

	workingDirectory, err := terraform.NewWorkingDirectory(filepath.Join(r.config.TfWorkingDirectory, strconv.FormatUint(uint64(applyExecutionRequest.ID), 10)))
	if err != nil {
		return nil, err
	}
	defer workingDirectory.DeleteDirectory()

	installationDirectory, err := terraform.NewTerraformInstallationDirectory(r.config.TfInstallationDirectory)
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
	if err != nil {
		fmt.Println(initOutput.String())
		return nil, err
	}

	additionalArguments := []string{}
	if applyExecutionRequest.AdditionalArguments != nil {
		additionalArguments = strings.Split(*applyExecutionRequest.AdditionalArguments, " ")
	}
	applyOutput := bytes.NewBuffer([]byte{})
	err = executable.TerraformApply(workingDirectory, applyExecutionRequest.TerraformPlan, additionalArguments, applyOutput)
	if err != nil {
		fmt.Println(applyOutput.String())
		return nil, err
	}

	newStatus = models.RequestStatusSucceeded
	completedTime := time.Now().UTC()
	_, err = r.apiClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequest.ID, &models.ApplyExecutionRequestUpdate{
		Status:      &newStatus,
		CompletedAt: &completedTime,
		InitOutput:  initOutput.Bytes(),
		ApplyOutput: applyOutput.Bytes(),
	})
	if err != nil {
		return nil, err
	}

	return applyExecutionRequest, nil
}
