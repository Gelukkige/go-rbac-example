package api

import (
	"go-rbac-example/internal/model"
	"go-rbac-example/internal/service"

	"github.com/gin-gonic/gin"
)

type RoleAPI struct {
	service *service.RoleService
}

func NewRoleAPI(service *service.RoleService) *RoleAPI {
	return &RoleAPI{service: service}
}

func (api *RoleAPI) CreateRole(c *gin.Context) {
	var req model.RoleCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.CreateRole(req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *RoleAPI) DeleteRole(c *gin.Context) {
	var req model.DeleteIDs
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.DeleteRole(req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *RoleAPI) UpdateRole(c *gin.Context) {
	var role model.RoleUpdateReq
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.UpdateRole(role); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *RoleAPI) ListRoles(c *gin.Context) {
	var page model.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	roles, total, err := api.service.ListRoles(page)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": roles, "total": total})
}
