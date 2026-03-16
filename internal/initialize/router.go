package initialize

import (
	"fmt"
	"go-rbac-example/internal/api"
	"go-rbac-example/internal/dao"
	"go-rbac-example/internal/global"
	"go-rbac-example/internal/service"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	serverConfig := global.Config.Server
	r := gin.Default()
	db := global.DB

	// USER
	userDao := dao.NewUserDao(db)
	userService := service.NewUserService(userDao)
	userAPI := api.NewUserAPI(userService)
	userGroup := r.Group("/users")
	{
		userGroup.POST("/create", userAPI.CreateUser)
		userGroup.DELETE("/delete", userAPI.DeleteUser)
		userGroup.PUT("/update", userAPI.UpdateUser)
		userGroup.POST("/list", userAPI.ListUsers)
	}

	addr := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	go r.Run(addr)
}
