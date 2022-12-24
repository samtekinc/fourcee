package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/pkg/execution/models"
)

func (r *ExecutionRouter) getApplyExecutionRequests(c *gin.Context) {
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.GetApplyExecutionRequests(c, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *ExecutionRouter) getApplyExecutionRequest(c *gin.Context) {
	id := c.Param("applyExecutionRequestId")
	response, err := r.apiClient.GetApplyExecutionRequest(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *ExecutionRouter) putApplyExecutionRequest(c *gin.Context) {
	var newApplyRequest models.NewApplyExecutionRequest
	err := c.BindJSON(&newApplyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.PutApplyExecutionRequest(c, &newApplyRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
