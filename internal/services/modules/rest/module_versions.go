package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sheacloud/tfom/pkg/modules/models"
)

func (r *ModulesRouter) getModuleVersions(c *gin.Context) {
	moduleGroupId := c.Param("moduleGroupId")
	limitString := c.DefaultQuery("limit", "10")
	cursor := c.Query("cursor")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ods, err := r.apiClient.GetModuleVersions(c, moduleGroupId, int32(limit), cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ods)
}

func (r *ModulesRouter) getModuleVersion(c *gin.Context) {
	moduleGroupId := c.Param("moduleGroupId")
	moduleVersionId := c.Param("moduleVersionId")
	od, err := r.apiClient.GetModuleVersion(c, moduleGroupId, moduleVersionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, od)
}

func (r *ModulesRouter) putModuleVersion(c *gin.Context) {
	var input models.NewModuleVersion
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	od, err := r.apiClient.PutModuleVersion(c, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, od)
}

func (r *ModulesRouter) deleteModuleVersion(c *gin.Context) {
	moduleGroupId := c.Param("moduleGroupId")
	moduleVersionId := c.Param("moduleVersionId")
	err := r.apiClient.DeleteModuleVersion(c, moduleGroupId, moduleVersionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
