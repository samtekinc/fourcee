package activities

import (
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/config"
)

type Activities struct {
	apiClient api.APIClientInterface
	config    *config.Config
}

func NewActivities(apiClient api.APIClientInterface, config *config.Config) *Activities {
	return &Activities{
		apiClient: apiClient,
		config:    config,
	}
}
