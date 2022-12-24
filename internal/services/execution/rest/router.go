package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/internal/services/execution/api"
)

type ExecutionRouter struct {
	apiClient api.ExecutionAPIClientInterface
}

func NewExecutionRouter(apiClient api.ExecutionAPIClientInterface) *ExecutionRouter {
	return &ExecutionRouter{
		apiClient: apiClient,
	}
}

func (r *ExecutionRouter) RegisterRoutes(router *gin.RouterGroup) {
	planRouter := router.Group("/plan")
	applyRouter := router.Group("/apply")

	planRouter.GET("/execution-requests", r.getPlanExecutionRequests)
	planRouter.GET("/execution-requests/:planExecutionRequestId", r.getPlanExecutionRequest)
	planRouter.POST("/execution-requests", r.putPlanExecutionRequest)

	applyRouter.GET("/execution-requests", r.getApplyExecutionRequests)
	applyRouter.GET("/execution-requests/:applyExecutionRequestId", r.getApplyExecutionRequest)
	applyRouter.POST("/execution-requests", r.putApplyExecutionRequest)
}
