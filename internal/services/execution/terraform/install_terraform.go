package terraform

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type TerraformInstallationDirectory struct {
	Directory string
}

func NewTerraformInstallationDirectory(directory string) (*TerraformInstallationDirectory, error) {
	directory, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}
	// create the directory
	err = os.MkdirAll(directory, 0755)
	if err != nil {
		return nil, err
	}

	return &TerraformInstallationDirectory{
		Directory: directory,
	}, nil
}

func (t *TerraformInstallationDirectory) InstallTerraform(version string, osName string, architecture string) (*TerraformExecutable, error) {
	if osName != "linux" && osName != "darwin" && osName != "windows" {
		return nil, fmt.Errorf("unsupported os: %s", osName)
	}
	if architecture != "amd64" && architecture != "arm64" {
		return nil, fmt.Errorf("unsupported architecture: %s", architecture)
	}

	var extension string
	switch osName {
	case "windows":
		extension = ".exe"
	default:
		extension = ""
	}
	executableFilepath := filepath.Join(t.Directory, fmt.Sprintf("terraform_%s_%s_%s%s", version, osName, architecture, extension))
	// check if terraform is already installed
	if _, err := os.Stat(executableFilepath); err == nil {
		return &TerraformExecutable{
			Filepath:     executableFilepath,
			Version:      version,
			OsName:       osName,
			Architecture: architecture,
		}, nil
	}

	downloadUrl := fmt.Sprintf("https://releases.hashicorp.com/terraform/%s/terraform_%s_%s_%s.zip", version, version, osName, architecture)

	response, err := http.Get(downloadUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to download terraform: %w", err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNotFound:
		return nil, fmt.Errorf("terraform version %s not found", version)
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("failed to download terraform: %s", response.Status)
	}

	// read zip file into byte buffer
	buff := bytes.NewBuffer([]byte{})
	zipFileLength, err := io.Copy(buff, response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read terraform zip file: %w", err)
	}

	reader, err := zip.NewReader(bytes.NewReader(buff.Bytes()), zipFileLength)
	if err != nil {
		return nil, fmt.Errorf("failed to read terraform zip file: %w", err)
	}
	if len(reader.File) != 1 {
		return nil, errors.New("terraform zip file should contain only one file")
	}
	terraformBinaryZipd, err := reader.File[0].Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open terraform binary within ZIP file: %w", err)
	}
	defer terraformBinaryZipd.Close()

	terraformBinaryBuff := bytes.NewBuffer([]byte{})
	_, err = io.Copy(terraformBinaryBuff, terraformBinaryZipd)
	if err != nil {
		return nil, fmt.Errorf("failed to read terraform binary within ZIP file: %w", err)
	}

	err = os.WriteFile(executableFilepath, terraformBinaryBuff.Bytes(), 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to write terraform binary: %w", err)
	}

	return &TerraformExecutable{
		Filepath:     executableFilepath,
		Version:      version,
		OsName:       osName,
		Architecture: architecture,
	}, nil
}
