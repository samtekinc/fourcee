package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/pkg/models"
)

func (r *OrganizationsRouter) getOrganizationalUnitsByDimension(c *gin.Context) {
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
	response, err := r.apiClient.GetOrganizationalUnitsByDimension(c, orgDimensionId, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) getOrganizationalUnitChildren(c *gin.Context) {
	orgDimensionId := c.Param("orgDimensionId")
	id := c.Param("orgUnitId")
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.GetOrganizationalUnitsByParent(c, orgDimensionId, id, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) getOrganizationalUnitDownstream(c *gin.Context) {
	orgDimensionId := c.Param("orgDimensionId")
	id := c.Param("orgUnitId")
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	parent, err := r.apiClient.GetOrganizationalUnit(c, orgDimensionId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	hierarchy := parent.Hierarchy + parent.OrgUnitId

	response, err := r.apiClient.GetOrganizationalUnitsByHierarchy(c, orgDimensionId, hierarchy, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) getOrganizationalUnit(c *gin.Context) {
	id := c.Param("orgUnitId")
	orgDimensionId := c.Param("orgDimensionId")
	response, err := r.apiClient.GetOrganizationalUnit(c, orgDimensionId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) putOrganizationalUnit(c *gin.Context) {
	var newOu models.NewOrganizationalUnit
	err := c.BindJSON(&newOu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.PutOrganizationalUnit(c, &newOu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) deleteOrganizationalUnit(c *gin.Context) {
	orgDimensionId := c.Param("orgDimensionId")
	id := c.Param("orgUnitId")
	err := r.apiClient.DeleteOrganizationalUnit(c, orgDimensionId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (r *OrganizationsRouter) updateOrganizationalUnit(c *gin.Context) {
	id := c.Param("orgUnitId")
	orgDimensionId := c.Param("orgDimensionId")
	var patch models.OrganizationalUnitUpdate
	err := c.BindJSON(&patch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := r.apiClient.UpdateOrganizationalUnit(c, orgDimensionId, id, &patch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *OrganizationsRouter) updateHierarchies(c *gin.Context) {
	orgDimensionId := c.Param("orgDimensionId")
	err := r.apiClient.UpdateOrganizationalUnitHierarchies(c, orgDimensionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
