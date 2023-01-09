package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *OrganizationsRouter) getModuleAssignmentsByModulePropagationId(c *gin.Context) {
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
	response, err := r.apiClient.GetModuleAssignmentsByModulePropagationId(c, modulePropagationId, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) getModuleAssignment(c *gin.Context) {
	modulePropagationId := c.Param("modulePropagationId")
	orgAccountId := c.Param("orgAccountId")
	response, err := r.apiClient.GetModuleAssignment(c, modulePropagationId, orgAccountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
