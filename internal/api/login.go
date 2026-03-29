package api

import (
	"go-rbac-example/internal/permission"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LogInAPI struct {
}

func NewLogInAPI() *LogInAPI {
	return &LogInAPI{}
}

func (api *LogInAPI) LogIn(c *gin.Context) {
	uid := c.Request.Header.Get("uid")
	uidInt, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	pages, err := permission.GetUserAllPages(c.Request.Context(), uidInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pages": pages})
}
