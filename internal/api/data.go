package api

import (
	"go-rbac-example/internal/permission"
	"go-rbac-example/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DataAPI struct {
	service *service.DataService
}

func NewDataAPI(service *service.DataService) *DataAPI {
	return &DataAPI{service: service}
}

func (api *DataAPI) ListData(c *gin.Context) {
	uidStr := c.Request.Header.Get("uid")
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	fields, err := permission.GetUserPagePermissions(c, uid, "data", "select")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := api.service.ListData(fields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
