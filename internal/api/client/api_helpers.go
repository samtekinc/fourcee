package client

import (
	"strconv"

	"github.com/samtekinc/fourcee/pkg/models"
	"gorm.io/gorm"
)

func applyPagination(limit *int, offset *int) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if limit != nil {
			tx = tx.Limit(*limit)
		}
		if offset != nil {
			tx = tx.Offset(*offset)
		}
		return tx
	}
}

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

func idToString(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}
