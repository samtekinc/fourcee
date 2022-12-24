package terraform

import (
	"encoding/base64"
	"os"
	"path/filepath"
)

type TerraformWorkingDirectory struct {
	Directory string
}

func NewWorkingDirectory(directory string) (*TerraformWorkingDirectory, error) {
	directory, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}
	// create the directory
	err = os.MkdirAll(directory, 0755)
	if err != nil {
		return nil, err
	}

	return &TerraformWorkingDirectory{
		Directory: directory,
	}, nil
}

func (t *TerraformWorkingDirectory) AddFile(fileBase64 string, fileName string) error {
	// create the file
	file, err := os.Create(filepath.Join(t.Directory, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	// write the file
	fileBytes, err := base64.StdEncoding.DecodeString(fileBase64)
	if err != nil {
		return err
	}
	_, err = file.Write(fileBytes)
	if err != nil {
		return err
	}

	return nil
}

func (t *TerraformWorkingDirectory) DeleteFile(fileName string) error {
	err := os.Remove(filepath.Join(t.Directory, fileName))
	if err != nil {
		return err
	}

	return nil
}

func (t *TerraformWorkingDirectory) DeleteDirectory() error {
	err := os.RemoveAll(t.Directory)
	if err != nil {
		return err
	}

	return nil
}
