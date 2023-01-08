package models

type CloudPlatform string

const (
	CloudPlatformAWS   CloudPlatform = "aws"
	CloudPlatformAzure CloudPlatform = "azure"
	CloudPlatformGCP   CloudPlatform = "gcp"
)
