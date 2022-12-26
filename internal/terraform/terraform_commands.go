package terraform

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sheacloud/tfom/pkg/models"
)

type TerraformExecutable struct {
	Filepath     string
	Version      string
	OsName       string
	Architecture string
}

func executeCommand(executable *TerraformExecutable, args []string, directory string, stdout io.Writer, stderr io.Writer) error {
	cmd := &exec.Cmd{
		Path:   executable.Filepath,
		Args:   append([]string{executable.Filepath}, args...),
		Dir:    directory,
		Stdin:  os.Stdin,
		Stdout: stdout,
		Stderr: stderr,
	}

	return cmd.Run()
}

func (t *TerraformExecutable) TerraformInit(workingDirectory *TerraformWorkingDirectory) *models.TerraformInitOutput {
	output := models.TerraformInitOutput{}

	stderrBuff := bytes.NewBuffer([]byte{})
	stdoutBuff := bytes.NewBuffer([]byte{})

	err := executeCommand(t, []string{"init"}, workingDirectory.Directory, stdoutBuff, stderrBuff)
	if err != nil {
		output.Error = fmt.Errorf("failed to run terraform init: %w", err)
		output.Stderr = stderrBuff.Bytes()
		return &output
	}

	output.Stdout = stdoutBuff.Bytes()
	output.Stderr = stderrBuff.Bytes()

	return &output
}

func (t *TerraformExecutable) TerraformPlan(workingDirectory *TerraformWorkingDirectory) *models.TerraformPlanOutput {
	output := models.TerraformPlanOutput{}

	stderrBuff := bytes.NewBuffer([]byte{})
	stdoutBuff := bytes.NewBuffer([]byte{})

	err := executeCommand(t, []string{"plan", "-out", "plan.tfplan"}, workingDirectory.Directory, stdoutBuff, stderrBuff)
	if err != nil {
		output.Error = fmt.Errorf("failed to run terraform plan: %w", err)
		output.Stderr = stderrBuff.Bytes()
		return &output
	}

	output.Stdout = stdoutBuff.Bytes()
	output.Stderr = stderrBuff.Bytes()

	// generate the JSON plan
	stderrBuff = bytes.NewBuffer([]byte{})
	stdoutBuff = bytes.NewBuffer([]byte{})

	err = executeCommand(t, []string{"show", "-json", "plan.tfplan"}, workingDirectory.Directory, stdoutBuff, stderrBuff)
	if err != nil {
		output.Error = fmt.Errorf("failed to run terraform show: %w", err)
		output.Stderr = stderrBuff.Bytes()
		return &output
	}

	output.PlanJSON = stdoutBuff.Bytes()

	// read the plan file
	planFilePath := filepath.Join(workingDirectory.Directory, "plan.tfplan")
	planFile, err := os.ReadFile(planFilePath)
	if err != nil {
		output.Error = fmt.Errorf("failed to read plan file: %w", err)
		return &output
	}
	output.PlanFile = planFile

	// delete the plan file
	err = os.Remove(planFilePath)
	if err != nil {
		output.Error = fmt.Errorf("failed to delete plan file: %w", err)
		return &output
	}

	return &output
}

func (t *TerraformExecutable) TerraformApply(workingDirectory *TerraformWorkingDirectory, terraformPlanFileBase64 string) *models.TerraformApplyOutput {
	output := models.TerraformApplyOutput{}

	// write the plan file

	err := workingDirectory.AddFile(terraformPlanFileBase64, "plan.tfplan")
	if err != nil {
		output.Error = fmt.Errorf("failed to write plan file: %w", err)
		return &output
	}

	stderrBuff := bytes.NewBuffer([]byte{})
	stdoutBuff := bytes.NewBuffer([]byte{})

	err = executeCommand(t, []string{"apply", "-auto-approve", "plan.tfplan"}, workingDirectory.Directory, stdoutBuff, stderrBuff)
	if err != nil {
		output.Error = fmt.Errorf("failed to run terraform apply: %w", err)
		output.Stderr = stderrBuff.Bytes()
		return &output
	}

	output.Stdout = stdoutBuff.Bytes()
	output.Stderr = stderrBuff.Bytes()

	// delete the plan file
	err = workingDirectory.DeleteFile("plan.tfplan")
	if err != nil {
		output.Error = fmt.Errorf("failed to delete plan file: %w", err)
		return &output
	}

	return &output
}

func (t *TerraformExecutable) TerraformCommand(workingDirectory *TerraformWorkingDirectory, args []string) *models.TerraformCommandOutput {
	output := models.TerraformCommandOutput{}

	stderrBuff := bytes.NewBuffer([]byte{})
	stdoutBuff := bytes.NewBuffer([]byte{})

	err := executeCommand(t, args, workingDirectory.Directory, stdoutBuff, stderrBuff)
	if err != nil {
		output.Error = fmt.Errorf("failed to run terraform command: %w", err)
		output.Stderr = stderrBuff.Bytes()
		return &output
	}

	output.Stdout = stdoutBuff.Bytes()
	output.Stderr = stderrBuff.Bytes()

	return &output
}
