package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/api"
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
	orgsRouter := router.Group("/organizations")
	modulesRouter := router.Group("/modules")
	planRouter := router.Group("/plan")
	applyRouter := router.Group("/apply")

	planRouter.GET("/execution-requests", r.getPlanExecutionRequests)
	planRouter.GET("/execution-requests/:planExecutionRequestId", r.getPlanExecutionRequest)
	planRouter.POST("/execution-requests", r.putPlanExecutionRequest)

	applyRouter.GET("/execution-requests", r.getApplyExecutionRequests)
	applyRouter.GET("/execution-requests/:applyExecutionRequestId", r.getApplyExecutionRequest)
	applyRouter.POST("/execution-requests", r.putApplyExecutionRequest)

	orgsRouter.GET("/dimensions", r.getOrganizationalDimensions)
	orgsRouter.GET("/dimensions/:orgDimensionId", r.getOrganizationalDimension)
	orgsRouter.POST("/dimensions", r.putOrganizationalDimension)
	orgsRouter.DELETE("/dimensions/:orgDimensionId", r.deleteOrganizationalDimension)
	orgsRouter.POST("/dimensions/:orgDimensionId/update-hierarchies", r.updateHierarchies)

	orgsRouter.GET("/dimensions/:orgDimensionId/ou-memberships", r.getOrganizationalUnitMembershipsByDimension)
	orgsRouter.POST("/dimensions/:orgDimensionId/ou-memberships", r.putOrganizationalUnitMembership)
	orgsRouter.DELETE("/dimensions/:orgDimensionId/ou-memberships/:accountId", r.deleteOrganizationalUnitMembership)

	orgsRouter.GET("/dimensions/:orgDimensionId/ous", r.getOrganizationalUnitsByDimension)
	orgsRouter.GET("/dimensions/:orgDimensionId/ous/:orgUnitId", r.getOrganizationalUnit)
	orgsRouter.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/children", r.getOrganizationalUnitChildren)
	orgsRouter.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/downstream", r.getOrganizationalUnitDownstream)
	orgsRouter.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/propagations", r.getModulePropagationsByOrgUnitId)
	orgsRouter.GET("/dimensions/:orgDimensionId/ous/:orgUnitId/ou-memberships", r.getOrganizationalUnitMembershipsByOrgUnit)
	orgsRouter.POST("/dimensions/:orgDimensionId/ous", r.putOrganizationalUnit)
	orgsRouter.PATCH("/dimensions/:orgDimensionId/ous/:orgUnitId", r.updateOrganizationalUnit)
	orgsRouter.DELETE("/dimensions/:orgDimensionId/ous/:orgUnitId", r.deleteOrganizationalUnit)

	orgsRouter.GET("/accounts", r.getOrganizationalAccounts)
	orgsRouter.POST("/accounts", r.putOrganizationalAccount)
	orgsRouter.GET("/accounts/:accountId", r.getOrganizationalAccount)
	orgsRouter.DELETE("/accounts/:accountId", r.deleteOrganizationalAccount)
	orgsRouter.GET("/accounts/:accountId/ou-memberships", r.getOrganizationalUnitMembershipsByAccount)

	modulesRouter.GET("/groups", r.getModuleGroups)
	modulesRouter.POST("/groups", r.putModuleGroup)
	modulesRouter.GET("/groups/:moduleGroupId", r.getModuleGroup)
	modulesRouter.DELETE("/groups/:moduleGroupId", r.deleteModuleGroup)
	modulesRouter.GET("/groups/:moduleGroupId/propagations", r.getModulePropagationsByModuleGroupId)

	modulesRouter.GET("/groups/:moduleGroupId/versions", r.getModuleVersions)
	modulesRouter.POST("/groups/:moduleGroupId/versions", r.putModuleVersion)
	modulesRouter.GET("/groups/:moduleGroupId/versions/:moduleVersionId", r.getModuleVersion)
	modulesRouter.DELETE("/groups/:moduleGroupId/versions/:moduleVersionId", r.deleteModuleVersion)
	modulesRouter.GET("/groups/:moduleGroupId/versions/:moduleVersionId/propagations", r.getModulePropagationsByModuleVersionId)

	modulesRouter.GET("/propagations", r.getModulePropagations)
	modulesRouter.POST("/propagations", r.putModulePropagation)
	modulesRouter.GET("/propagations/:modulePropagationId", r.getModulePropagation)
	modulesRouter.GET("/propagations/:modulePropagationId/downstream-ous", r.getModulePropagationDownstreamOUs)
	modulesRouter.DELETE("/propagations/:modulePropagationId", r.deleteModulePropagation)

	modulesRouter.GET("/propagations/:modulePropagationId/executions", r.getModulePropagationExecutionRequestsByModulePropagationId)
	modulesRouter.GET("/propagations/:modulePropagationId/executions/:modulePropagationExecutionRequestId", r.getModulePropagationExecutionRequest)
	modulesRouter.POST("/propagations/:modulePropagationId/executions", r.putModulePropagationExecutionRequest)

	modulesRouter.GET("/propagations/:modulePropagationId/associations", r.getModuleAccountAssociationsByModulePropagationId)
	modulesRouter.GET("/propagations/:modulePropagationId/associations/:orgAccountId", r.getModuleAccountAssociation)
}
