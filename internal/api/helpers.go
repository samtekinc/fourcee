package api

import (
	"github.com/sheacloud/tfom/pkg/models"
)

func ArgumentInputsToArguments(inputs []models.ArgumentInput) []models.Argument {
	arguments := make([]models.Argument, len(inputs))
	for i, input := range inputs {
		arguments[i] = models.Argument(input)
	}
	return arguments
}

func AwsProviderConfigurationInputsToAwsProviderConfigurations(inputs []models.AwsProviderConfigurationInput) []models.AwsProviderConfiguration {
	providers := make([]models.AwsProviderConfiguration, len(inputs))
	for i, input := range inputs {
		providers[i] = models.AwsProviderConfiguration(input)
	}
	return providers
}

func GcpProviderConfigurationInputsToGcpProviderConfigurations(inputs []models.GcpProviderConfigurationInput) []models.GcpProviderConfiguration {
	providers := make([]models.GcpProviderConfiguration, len(inputs))
	for i, input := range inputs {
		providers[i] = models.GcpProviderConfiguration(input)
	}
	return providers
}

func MetadataInputsToMetadata(inputs []models.MetadataInput) []models.Metadata {
	metadata := make([]models.Metadata, len(inputs))
	for i, input := range inputs {
		metadata[i] = models.Metadata(input)
	}
	return metadata
}
