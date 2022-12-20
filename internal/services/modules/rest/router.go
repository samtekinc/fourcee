package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/services/modules/api"
)

type ModulesRouter struct {
	apiClient api.ModulesAPIClientInterface
}

func NewModulesRouter(apiClient api.ModulesAPIClientInterface) *ModulesRouter {
	return &ModulesRouter{
		apiClient: apiClient,
	}
}

func (r *ModulesRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/groups", r.getModuleGroups)
	router.POST("/groups", r.putModuleGroup)
	router.GET("/groups/:moduleGroupId", r.getModuleGroup)
	router.DELETE("/groups/:moduleGroupId", r.deleteModuleGroup)

	router.GET("/groups/:moduleGroupId/versions", r.getModuleVersions)
	router.POST("/groups/:moduleGroupId/versions", r.putModuleVersion)
	router.GET("/groups/:moduleGroupId/versions/:moduleVersionId", r.getModuleVersion)
	router.DELETE("/groups/:moduleGroupId/versions/:moduleVersionId", r.deleteModuleVersion)

}
