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
	router.GET("/dimensions/:dimensionId", r.getOrganizationalDimension)
	router.DELETE("/dimensions/:dimensionId", r.deleteOrganizationalDimension)
	router.POST("/dimensions/:dimensionId/update-hierarchies", r.updateHierarchies)

	router.GET("/dimensions/:dimensionId/organizational-units", r.getOrganizationalUnitsByDimension)
	router.GET("/dimensions/:dimensionId/organizational-units/:orgUnitId", r.getOrganizationalUnit)
	router.GET("/dimensions/:dimensionId/organizational-units/:orgUnitId/children", r.getOrganizationalUnitChildren)
	router.GET("/dimensions/:dimensionId/organizational-units/:orgUnitId/downstream", r.getOrganizationalUnitDownstream)
	router.PATCH("/dimensions/:dimensionId/organizational-units/:orgUnitId", r.updateOrganizationalUnit)
	router.POST("/dimensions/:dimensionId/organizational-units", r.putOrganizationalUnit)
	router.DELETE("/dimensions/:dimensionId/organizational-units/:orgUnitId", r.deleteOrganizationalUnit)

	router.GET("/accounts", r.getOrganizationalAccounts)
	router.POST("/accounts", r.putOrganizationalAccount)
	router.GET("/accounts/:accountId", r.getOrganizationalAccount)
	router.DELETE("/accounts/:accountId", r.deleteOrganizationalAccount)

	router.GET("/accounts/:accountId/organizational-unit-memberships", r.getOrganizationalUnitMembershipsByAccount)
	router.GET("/dimensions/:dimensionId/organizational-units/:orgUnitId/organizational-unit-memberships", r.getOrganizationalUnitMembershipsByOrgUnit)
	router.GET("/dimensions/:dimensionId/organizational-unit-memberships", r.getOrganizationalUnitMembershipsByDimension)
	router.POST("/dimensions/:dimensionId/organizational-unit-memberships", r.putOrganizationalUnitMembership)
	router.DELETE("/dimensions/:dimensionId/organizational-unit-memberships/:accountId", r.deleteOrganizationalUnitMembership)
}
