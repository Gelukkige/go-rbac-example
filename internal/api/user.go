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
