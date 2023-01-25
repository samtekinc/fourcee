package api

import (
	"strconv"

	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
)

func applyPagination(tx *gorm.DB, limit *int, offset *int) *gorm.DB {
	if limit != nil {
		tx = tx.Limit(*limit)
	}
	if offset != nil {
		tx = tx.Offset(*offset)
	}
	return tx
}

func ArgumentInputsToArguments(inputs []models.ArgumentInput) []models.Argument {
	arguments := make([]models.Argument, len(inputs))
	for i, input := range inputs {
		arguments[i] = models.Argument{
			Name:  input.Name,
			Value: input.Value,
		}
	}
	return arguments
}

func AwsProviderConfigurationInputsToAwsProviderConfigurations(inputs []models.AwsProviderConfigurationInput) []models.AwsProviderConfiguration {
	providers := make([]models.AwsProviderConfiguration, len(inputs))
	for i, input := range inputs {
		providers[i] = models.AwsProviderConfiguration{
			Region: input.Region,
			Alias:  input.Alias,
		}
	}
	return providers
}

func GcpProviderConfigurationInputsToGcpProviderConfigurations(inputs []models.GcpProviderConfigurationInput) []models.GcpProviderConfiguration {
	providers := make([]models.GcpProviderConfiguration, len(inputs))
	for i, input := range inputs {
		providers[i] = models.GcpProviderConfiguration{
			Region: input.Region,
			Alias:  input.Alias,
		}
	}
	return providers
}

func MetadataInputsToMetadata(inputs []models.MetadataInput) []models.Metadata {
	metadata := make([]models.Metadata, len(inputs))
	for i, input := range inputs {
		metadata[i] = models.Metadata{
			Name:  input.Name,
			Value: input.Value,
		}
	}
	return metadata
}

func idToString(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}
