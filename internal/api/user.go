package api

import (
	"go-rbac-example/internal/model"
	"go-rbac-example/internal/service"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	service *service.UserService
}

func NewUserAPI(service *service.UserService) *UserAPI {
	return &UserAPI{service: service}
}

func (api *UserAPI) CreateUser(c *gin.Context) {
	var req model.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.CreateUser(req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *UserAPI) DeleteUser(c *gin.Context) {
	var req model.DeleteIDs
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.DeleteUser(req.IDs); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *UserAPI) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.UpdateUser(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (api *UserAPI) ListUsers(c *gin.Context) {
	var page model.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	users, total, err := api.service.ListUsers(page)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": users, "total": total})
}
