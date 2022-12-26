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

func ProvidersInputToProviders(inputs []models.ProviderInput) []models.Provider {
	providers := make([]models.Provider, len(inputs))
	for i, input := range inputs {
		providers[i] = models.Provider{
			Name:      input.Name,
			Arguments: ArgumentInputsToArguments(input.Arguments),
		}
	}
	return providers
}
