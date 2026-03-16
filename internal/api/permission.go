package api

import (
	"go-rbac-example/internal/model"
	"go-rbac-example/internal/service"

	"github.com/gin-gonic/gin"
)

type PermissionAPI struct {
	service *service.PermissionService
}

func NewPermissionAPI(service *service.PermissionService) *PermissionAPI {
	return &PermissionAPI{service: service}
}

func (api *PermissionAPI) CreatePermission(c *gin.Context) {
	var req model.PermissionCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.CreatePermission(req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *PermissionAPI) DeletePermission(c *gin.Context) {
	var req model.DeleteIDs
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.DeletePermission(req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *PermissionAPI) UpdatePermission(c *gin.Context) {
	var permission model.PermissionUpdateReq
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.UpdatePermission(permission); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *PermissionAPI) ListPermissions(c *gin.Context) {
	var page model.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	permissions, total, err := api.service.ListPermissions(page)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": permissions, "total": total})
}
