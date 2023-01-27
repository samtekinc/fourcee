package config

import (
	"os"
)

type Config struct {
	VersionInstallationDirectory string
	TfInstallationDirectory      string
	TfWorkingDirectory           string
	Prefix                       string
	StateBucket                  string
	StateRegion                  string
	ResultsBucket                string
	AccountId                    string
	Region                       string
	AlertsTopic                  string
}

func ConfigFromEnv() Config {
	versionInstallationDirectory := os.Getenv("TFOM_VERSION_INSTALLATION_DIRECTORY")
	if versionInstallationDirectory == "" {
		panic("TFOM_VERSION_INSTALLATION_DIRECTORY must be set")
	}
	tfInstallationDirectory := os.Getenv("TFOM_TF_INSTALLATION_DIRECTORY")
	if tfInstallationDirectory == "" {
		panic("TFOM_TF_INSTALLATION_DIRECTORY must be set")
	}
	tfWorkingDirectory := os.Getenv("TFOM_TF_WORKING_DIRECTORY")
	if tfWorkingDirectory == "" {
		panic("TFOM_TF_WORKING_DIRECTORY must be set")
	}
	prefix := os.Getenv("TFOM_PREFIX")
	if prefix == "" {
		panic("TFOM_PREFIX must be set")
	}
	stateBucket := os.Getenv("TFOM_STATE_BUCKET")
	if stateBucket == "" {
		panic("TFOM_STATE_BUCKET must be set")
	}
	stateRegion := os.Getenv("TFOM_STATE_REGION")
	if stateRegion == "" {
		panic("TFOM_STATE_REGION must be set")
	}
	resultsBucket := os.Getenv("TFOM_RESULTS_BUCKET")
	if resultsBucket == "" {
		panic("TFOM_RESULTS_BUCKET must be set")
	}
	accountId := os.Getenv("TFOM_ACCOUNT_ID")
	if accountId == "" {
		panic("TFOM_ACCOUNT_ID must be set")
	}
	region := os.Getenv("TFOM_REGION")
	if region == "" {
		panic("TFOM_REGION must be set")
	}
	alertsTopic := os.Getenv("TFOM_ALERTS_TOPIC")
	if alertsTopic == "" {
		panic("TFOM_ALERTS_TOPIC must be set")
	}

	return Config{
		VersionInstallationDirectory: versionInstallationDirectory,
		TfInstallationDirectory:      tfInstallationDirectory,
		TfWorkingDirectory:           tfWorkingDirectory,
		Prefix:                       prefix,
		StateBucket:                  stateBucket,
		StateRegion:                  stateRegion,
		ResultsBucket:                resultsBucket,
		AccountId:                    accountId,
		Region:                       region,
		AlertsTopic:                  alertsTopic,
	}
}
