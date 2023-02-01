package config

import (
	"os"
)

type Config struct {
	WorkingDirectory   string
	Prefix             string
	StateBucket        string
	StateRegion        string
	ResultsBucket      string
	AccountId          string
	Region             string
	AlertsTopic        string
	DBConnectionString string
}

func ConfigFromEnv() Config {
	workingDirectory := os.Getenv("TFOM_WORKING_DIRECTORY")
	if workingDirectory == "" {
		panic("TFOM_WORKING_DIRECTORY must be set")
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
	dbConnectionString := os.Getenv("TFOM_DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		panic("TFOM_DB_CONNECTION_STRING must be set")
	}

	return Config{
		WorkingDirectory:   workingDirectory,
		Prefix:             prefix,
		StateBucket:        stateBucket,
		StateRegion:        stateRegion,
		ResultsBucket:      resultsBucket,
		AccountId:          accountId,
		Region:             region,
		AlertsTopic:        alertsTopic,
		DBConnectionString: dbConnectionString,
	}
}
