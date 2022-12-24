package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/services/organizations/api"
)

type OrganizationsRouter struct {
	apiClient api.OrganizationsAPIClientInterface
}

func NewOrganizationsRouter(apiClient api.OrganizationsAPIClientInterface) *OrganizationsRouter {
	return &OrganizationsRouter{
		apiClient: apiClient,
	}
}

func (r *OrganizationsRouter) RegisterRoutes(router *gin.RouterGroup) {
	orgs := router.Group("/organizations")
	modules := router.Group("/modules")

	orgs.GET("/dimensions", r.getOrganizationalDimensions)
	orgs.GET("/dimensions/:orgDimensionId", r.getOrganizationalDimension)
	orgs.POST("/dimensions", r.putOrganizationalDimension)
	orgs.DELETE("/dimensions/:orgDimensionId", r.deleteOrganizationalDimension)
	orgs.POST("/dimensions/:orgDimensionId/update-hierarchies", r.updateHierarchies)

	orgs.GET("/dimensions/:orgDimensionId/ou-memberships", r.getOrganizationalUnitMembershipsByDimension)
	orgs.POST("/dimensions/:orgDimensionId/ou-memberships", r.putOrganizationalUnitMembership)
	orgs.DELETE("/dimensions/:orgDimensionId/ou-memberships/:accountId", r.deleteOrganizationalUnitMembership)

	orgs.GET("/dimensions/:orgDimensionId/ous", r.getOrganizationalUnitsByDimension)
	orgs.GET("/dimensions/:orgDimensionId/ous/:orgUnitId", r.getOrganizationalUnit)
	orgs.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/children", r.getOrganizationalUnitChildren)
	orgs.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/downstream", r.getOrganizationalUnitDownstream)
	orgs.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/propagations", r.getModulePropagationsByOrgUnitId)
	orgs.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/ou-memberships", r.getOrganizationalUnitMembershipsByOrgUnit)
	orgs.POST("/dimensions/:orgDimensionId/ous", r.putOrganizationalUnit)
	orgs.PATCH("/dimensions/:orgDimensionId/ous/:orgUnitId", r.updateOrganizationalUnit)
	orgs.DELETE("/dimensions/:orgDimensionId/ous/:orgUnitId", r.deleteOrganizationalUnit)

	orgs.GET("/accounts", r.getOrganizationalAccounts)
	orgs.POST("/accounts", r.putOrganizationalAccount)
	orgs.GET("/accounts/:accountId", r.getOrganizationalAccount)
	orgs.DELETE("/accounts/:accountId", r.deleteOrganizationalAccount)
	orgs.GET("/accounts/:accountId/ou-memberships", r.getOrganizationalUnitMembershipsByAccount)

	modules.GET("/groups", r.getModuleGroups)
	modules.POST("/groups", r.putModuleGroup)
	modules.GET("/groups/:moduleGroupId", r.getModuleGroup)
	modules.DELETE("/groups/:moduleGroupId", r.deleteModuleGroup)
	modules.GET("/groups/:moduleGroupId/propagations", r.getModulePropagationsByModuleGroupId)

	modules.GET("/groups/:moduleGroupId/versions", r.getModuleVersions)
	modules.POST("/groups/:moduleGroupId/versions", r.putModuleVersion)
	modules.GET("/groups/:moduleGroupId/versions/:moduleVersionId", r.getModuleVersion)
	modules.DELETE("/groups/:moduleGroupId/versions/:moduleVersionId", r.deleteModuleVersion)
	modules.GET("/groups/:moduleGroupId/versions/:moduleVersionId/propagations", r.getModulePropagationsByModuleVersionId)

	modules.GET("/propagations", r.getModulePropagations)
	modules.POST("/propagations", r.putModulePropagation)
	modules.GET("/propagations/:modulePropagationId", r.getModulePropagation)
	modules.GET("/propagations/:modulePropagationId/downstream-ous", r.getModulePropagationDownstreamOUs)
	modules.DELETE("/propagations/:modulePropagationId", r.deleteModulePropagation)

	modules.GET("/propagations/:modulePropagationId/executions", r.getModulePropagationExecutionRequestsByModulePropagationId)
	modules.GET("/propagations/:modulePropagationId/executions/:modulePropagationExecutionRequestId", r.getModulePropagationExecutionRequest)
	modules.POST("/propagations/:modulePropagationId/executions", r.putModulePropagationExecutionRequest)
}
