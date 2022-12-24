package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/pkg/execution/models"
)

func (r *ExecutionRouter) getPlanExecutionRequests(c *gin.Context) {
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.GetPlanExecutionRequests(c, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *ExecutionRouter) getPlanExecutionRequest(c *gin.Context) {
	id := c.Param("planExecutionRequestId")
	response, err := r.apiClient.GetPlanExecutionRequest(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *ExecutionRouter) putPlanExecutionRequest(c *gin.Context) {
	var newPlanRequest models.NewPlanExecutionRequest
	err := c.BindJSON(&newPlanRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.PutPlanExecutionRequest(c, &newPlanRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
