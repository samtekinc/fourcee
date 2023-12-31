package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/samtekinc/fourcee/internal/graph/generated"
	"github.com/samtekinc/fourcee/pkg/models"
)

// ModulePropagation is the resolver for the modulePropagation field.
func (r *modulePropagationExecutionRequestResolver) ModulePropagation(ctx context.Context, obj *models.ModulePropagationExecutionRequest) (*models.ModulePropagation, error) {
	return r.apiClient.GetModulePropagationBatched(ctx, obj.ModulePropagationID)
}

// TerraformExecutionRequests is the resolver for the terraformExecutionRequests field.
func (r *modulePropagationExecutionRequestResolver) TerraformExecutionRequests(ctx context.Context, obj *models.ModulePropagationExecutionRequest, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error) {
	return r.apiClient.GetTerraformExecutionRequestsForModulePropagationExecutionRequest(ctx, obj.ID, filters, limit, offset)
}

// CreateModulePropagationExecutionRequest is the resolver for the createModulePropagationExecutionRequest field.
func (r *mutationResolver) CreateModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequest models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error) {
	return r.apiClient.CreateModulePropagationExecutionRequest(ctx, &modulePropagationExecutionRequest)
}

// ModulePropagationExecutionRequest is the resolver for the modulePropagationExecutionRequest field.
func (r *queryResolver) ModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequestID uint) (*models.ModulePropagationExecutionRequest, error) {
	return r.apiClient.GetModulePropagationExecutionRequest(ctx, modulePropagationExecutionRequestID)
}

// ModulePropagationExecutionRequests is the resolver for the modulePropagationExecutionRequests field.
func (r *queryResolver) ModulePropagationExecutionRequests(ctx context.Context, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error) {
	return r.apiClient.GetModulePropagationExecutionRequests(ctx, filters, limit, offset)
}

// ModulePropagationExecutionRequest returns generated.ModulePropagationExecutionRequestResolver implementation.
func (r *Resolver) ModulePropagationExecutionRequest() generated.ModulePropagationExecutionRequestResolver {
	return &modulePropagationExecutionRequestResolver{r}
}

type modulePropagationExecutionRequestResolver struct{ *Resolver }
