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

// CreateOrganizationalDimension is the resolver for the createOrganizationalDimension field.
func (r *mutationResolver) CreateOrganizationalDimension(ctx context.Context, orgDimension models.NewOrganizationalDimension) (*models.OrganizationalDimension, error) {
	return r.apiClient.PutOrganizationalDimension(ctx, &orgDimension)
}

// DeleteOrganizationalDimension is the resolver for the deleteOrganizationalDimension field.
func (r *mutationResolver) DeleteOrganizationalDimension(ctx context.Context, orgDimensionID string) (bool, error) {
	err := r.apiClient.DeleteOrganizationalDimension(ctx, orgDimensionID)
	return err == nil, err
}

// RootOrgUnit is the resolver for the rootOrgUnit field.
func (r *organizationalDimensionResolver) RootOrgUnit(ctx context.Context, obj *models.OrganizationalDimension) (*models.OrganizationalUnit, error) {
	return r.apiClient.GetOrganizationalUnitBatched(ctx, obj.OrgDimensionId, obj.RootOrgUnitId)
}

// OrgUnits is the resolver for the orgUnits field.
func (r *organizationalDimensionResolver) OrgUnits(ctx context.Context, obj *models.OrganizationalDimension, limit *int, nextCursor *string) (*models.OrganizationalUnits, error) {
	if limit == nil {
		limit = aws.Int(100)
	}
	return r.apiClient.GetOrganizationalUnitsByDimension(ctx, obj.OrgDimensionId, int32(*limit), aws.ToString(nextCursor))
}

// OrgUnitMemberships is the resolver for the orgUnitMemberships field.
func (r *organizationalDimensionResolver) OrgUnitMemberships(ctx context.Context, obj *models.OrganizationalDimension, limit *int, nextCursor *string) (*models.OrganizationalUnitMemberships, error) {
	if limit == nil {
		limit = aws.Int(100)
	}
	return r.apiClient.GetOrganizationalUnitMembershipsByDimension(ctx, obj.OrgDimensionId, int32(*limit), aws.ToString(nextCursor))
}

// ModulePropagations is the resolver for the modulePropagations field.
func (r *organizationalDimensionResolver) ModulePropagations(ctx context.Context, obj *models.OrganizationalDimension, limit *int, nextCursor *string) (*models.ModulePropagations, error) {
	if limit == nil {
		limit = aws.Int(100)
	}
	return r.apiClient.GetModulePropagationsByOrgDimensionId(ctx, obj.OrgDimensionId, int32(*limit), aws.ToString(nextCursor))
}

// OrganizationalDimension is the resolver for the organizationalDimension field.
func (r *queryResolver) OrganizationalDimension(ctx context.Context, orgDimensionID string) (*models.OrganizationalDimension, error) {
	return r.apiClient.GetOrganizationalDimension(ctx, orgDimensionID)
}

// OrganizationalDimensions is the resolver for the organizationalDimensions field.
func (r *queryResolver) OrganizationalDimensions(ctx context.Context, limit *int, nextCursor *string) (*models.OrganizationalDimensions, error) {
	if limit == nil {
		limit = aws.Int(100)
	}
	return r.apiClient.GetOrganizationalDimensions(ctx, int32(*limit), aws.ToString(nextCursor))
}

// OrganizationalDimension returns generated.OrganizationalDimensionResolver implementation.
func (r *Resolver) OrganizationalDimension() generated.OrganizationalDimensionResolver {
	return &organizationalDimensionResolver{r}
}

type organizationalDimensionResolver struct{ *Resolver }
