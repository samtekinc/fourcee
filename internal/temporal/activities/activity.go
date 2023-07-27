package activities

import (
	"github.com/samtekinc/fourcee/internal/api"
	"github.com/samtekinc/fourcee/internal/config"
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
