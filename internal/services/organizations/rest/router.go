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
	router.GET("/dimensions", r.getOrganizationalDimensions)
	router.POST("/dimensions", r.putOrganizationalDimension)
	router.GET("/dimensions/:orgDimensionId", r.getOrganizationalDimension)
	router.DELETE("/dimensions/:orgDimensionId", r.deleteOrganizationalDimension)
	router.POST("/dimensions/:orgDimensionId/update-hierarchies", r.updateHierarchies)

	router.GET("/dimensions/:orgDimensionId/organizational-units", r.getOrganizationalUnitsByDimension)
	router.GET("/dimensions/:orgDimensionId/organizational-units/:orgUnitId", r.getOrganizationalUnit)
	router.GET("/dimensions/:orgDimensionId/organizational-units/:orgUnitId/children", r.getOrganizationalUnitChildren)
	router.GET("/dimensions/:orgDimensionId/organizational-units/:orgUnitId/downstream", r.getOrganizationalUnitDownstream)
	router.PATCH("/dimensions/:orgDimensionId/organizational-units/:orgUnitId", r.updateOrganizationalUnit)
	router.POST("/dimensions/:orgDimensionId/organizational-units", r.putOrganizationalUnit)
	router.DELETE("/dimensions/:orgDimensionId/organizational-units/:orgUnitId", r.deleteOrganizationalUnit)

	router.GET("/accounts", r.getOrganizationalAccounts)
	router.POST("/accounts", r.putOrganizationalAccount)
	router.GET("/accounts/:accountId", r.getOrganizationalAccount)
	router.DELETE("/accounts/:accountId", r.deleteOrganizationalAccount)

	router.GET("/accounts/:accountId/organizational-unit-memberships", r.getOrganizationalUnitMembershipsByAccount)
	router.GET("/dimensions/:orgDimensionId/organizational-units/:orgUnitId/organizational-unit-memberships", r.getOrganizationalUnitMembershipsByOrgUnit)
	router.GET("/dimensions/:orgDimensionId/organizational-unit-memberships", r.getOrganizationalUnitMembershipsByDimension)
	router.POST("/dimensions/:orgDimensionId/organizational-unit-memberships", r.putOrganizationalUnitMembership)
	router.DELETE("/dimensions/:orgDimensionId/organizational-unit-memberships/:accountId", r.deleteOrganizationalUnitMembership)
}
