package terraform

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func (t *TerraformExecutable) TerraformInit(workingDirectory *TerraformWorkingDirectory, output io.Writer) error {

	output.Write([]byte(fmt.Sprintf("%s init\n\n", t.Filepath)))

	err := executeCommand(t, []string{"init"}, workingDirectory.Directory, output, output)
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to run terraform init: %s", err.Error())))
		return err
	}

	return nil
}

func (t *TerraformExecutable) TerraformPlan(workingDirectory *TerraformWorkingDirectory, additionalArguments []string, output io.Writer, planWriter io.Writer, planJSONWriter io.Writer) error {
	output.Write([]byte(fmt.Sprintf("%s plan -out plan.tfplan %s\n\n", t.Filepath, strings.Join(additionalArguments, " "))))

	err := executeCommand(t, append([]string{"plan", "-out", "plan.tfplan"}, additionalArguments...), workingDirectory.Directory, output, output)
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to run terraform plan: %s", err.Error())))
		return err
	}

	err = executeCommand(t, []string{"show", "-json", "plan.tfplan"}, workingDirectory.Directory, planJSONWriter, planJSONWriter)
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to run terraform show: %s", err.Error())))
		return err
	}

	// read the plan file
	planFilePath := filepath.Join(workingDirectory.Directory, "plan.tfplan")
	planFile, err := os.ReadFile(planFilePath)
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to read plan file: %s", err.Error())))
		return err
	}
	planWriter.Write(planFile)

	// delete the plan file
	err = os.Remove(planFilePath)
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to delete plan file: %s", err.Error())))
		return err
	}

	return nil
}

func (t *TerraformExecutable) TerraformApply(workingDirectory *TerraformWorkingDirectory, terraformPlanFile []byte, additionalArguments []string, output io.Writer) error {
	// write the plan file
	err := workingDirectory.AddFile(terraformPlanFile, "plan.tfplan")
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to write plan file: %s", err.Error())))
		return err
	}

	args := append([]string{"apply"}, additionalArguments...)
	args = append(args, "-auto-approve", "plan.tfplan")

	output.Write([]byte(fmt.Sprintf("%s %s\n\n", t.Filepath, strings.Join(args, " "))))

	err = executeCommand(t, args, workingDirectory.Directory, output, output)
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to run terraform apply: %s", err.Error())))
		return err
	}

	// delete the plan file
	err = workingDirectory.DeleteFile("plan.tfplan")
	if err != nil {
		return err
	}

	return nil
}

func (t *TerraformExecutable) TerraformCommand(workingDirectory *TerraformWorkingDirectory, args []string, output io.Writer) error {
	err := executeCommand(t, args, workingDirectory.Directory, output, output)
	if err != nil {
		output.Write([]byte(fmt.Sprintf("failed to run terraform command: %s", err.Error())))
		return err
	}

	return nil
}
