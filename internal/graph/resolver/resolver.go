package resolver

import (
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/config"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	apiClient api.APIClientInterface
	config    *config.Config
}

func NewResolver(apiClient api.APIClientInterface, config *config.Config) *Resolver {
	return &Resolver{apiClient: apiClient, config: config}
}
