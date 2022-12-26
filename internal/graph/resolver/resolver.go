package resolver

import "github.com/sheacloud/tfom/internal/api"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	apiClient api.OrganizationsAPIClientInterface
}

func NewResolver(apiClient api.OrganizationsAPIClientInterface) *Resolver {
	return &Resolver{apiClient: apiClient}
}
