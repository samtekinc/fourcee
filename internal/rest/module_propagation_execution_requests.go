package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/pkg/models"
)

func (r *OrganizationsRouter) getModulePropagationExecutionRequestsByModulePropagationId(c *gin.Context) {
	modulePropagationId := c.Param("modulePropagationId")
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.GetModulePropagationExecutionRequestsByModulePropagationId(c, modulePropagationId, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) getModulePropagationExecutionRequest(c *gin.Context) {
	modulePropagationId := c.Param("modulePropagationId")
	modulePropagationExecutionRequestId := c.Param("modulePropagationExecutionRequestId")
	response, err := r.apiClient.GetModulePropagationExecutionRequest(c, modulePropagationId, modulePropagationExecutionRequestId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) putModulePropagationExecutionRequest(c *gin.Context) {
	var newRequest models.NewModulePropagationExecutionRequest
	err := c.BindJSON(&newRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.PutModulePropagationExecutionRequest(c, &newRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
