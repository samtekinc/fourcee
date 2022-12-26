package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/pkg/models"
)

func (r *OrganizationsRouter) getOrganizationalUnitMembershipsByAccount(c *gin.Context) {
	accountId := c.Param("accountId")
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.GetOrganizationalUnitMembershipsByAccount(c, accountId, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) getOrganizationalUnitMembershipsByOrgUnit(c *gin.Context) {
	orgUnitId := c.Param("orgUnitId")
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.GetOrganizationalUnitMembershipsByOrgUnit(c, orgUnitId, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) getOrganizationalUnitMembershipsByDimension(c *gin.Context) {
	orgDimensionId := c.Param("orgDimensionId")
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.GetOrganizationalUnitMembershipsByAccount(c, orgDimensionId, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) putOrganizationalUnitMembership(c *gin.Context) {
	var newOum models.NewOrganizationalUnitMembership
	err := c.BindJSON(&newOum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.PutOrganizationalUnitMembership(c, &newOum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) deleteOrganizationalUnitMembership(c *gin.Context) {
	orgDimensionId := c.Param("orgDimensionId")
	accountId := c.Param("accountId")
	err := r.apiClient.DeleteOrganizationalUnitMembership(c, orgDimensionId, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
