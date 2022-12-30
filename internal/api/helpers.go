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
