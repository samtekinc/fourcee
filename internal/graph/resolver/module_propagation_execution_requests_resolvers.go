package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/tfom/internal/graph/generated"
	"github.com/sheacloud/tfom/pkg/models"
)

// ModulePropagation is the resolver for the modulePropagation field.
func (r *modulePropagationExecutionRequestResolver) ModulePropagation(ctx context.Context, obj *models.ModulePropagationExecutionRequest) (*models.ModulePropagation, error) {
	return r.apiClient.GetModulePropagationBatched(ctx, obj.ModulePropagationId)
}

// TerraformExecutionRequests is the resolver for the terraformExecutionRequests field.
func (r *modulePropagationExecutionRequestResolver) TerraformExecutionRequests(ctx context.Context, obj *models.ModulePropagationExecutionRequest, limit *int, nextCursor *string) (*models.TerraformExecutionRequests, error) {
	if limit == nil {
		limit = aws.Int(100)
	}

	return r.apiClient.GetTerraformExecutionRequestsByModulePropagationExecutionRequestId(ctx, obj.ModulePropagationExecutionRequestId, int32(*limit), aws.ToString(nextCursor))
}

// CreateModulePropagationExecutionRequest is the resolver for the createModulePropagationExecutionRequest field.
func (r *mutationResolver) CreateModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequest models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error) {
	return r.apiClient.PutModulePropagationExecutionRequest(ctx, &modulePropagationExecutionRequest)
}

// ModulePropagationExecutionRequest is the resolver for the modulePropagationExecutionRequest field.
func (r *queryResolver) ModulePropagationExecutionRequest(ctx context.Context, modulePropagationID string, modulePropagationExecutionRequestID string) (*models.ModulePropagationExecutionRequest, error) {
	return r.apiClient.GetModulePropagationExecutionRequest(ctx, modulePropagationID, modulePropagationExecutionRequestID)
}

// ModulePropagationExecutionRequests is the resolver for the modulePropagationExecutionRequests field.
func (r *queryResolver) ModulePropagationExecutionRequests(ctx context.Context, limit *int, nextCursor *string) (*models.ModulePropagationExecutionRequests, error) {
	if limit == nil {
		limit = aws.Int(100)
	}

	return r.apiClient.GetModulePropagationExecutionRequests(ctx, int32(*limit), aws.ToString(nextCursor))
}

// ModulePropagationExecutionRequest returns generated.ModulePropagationExecutionRequestResolver implementation.
func (r *Resolver) ModulePropagationExecutionRequest() generated.ModulePropagationExecutionRequestResolver {
	return &modulePropagationExecutionRequestResolver{r}
}

type modulePropagationExecutionRequestResolver struct{ *Resolver }
